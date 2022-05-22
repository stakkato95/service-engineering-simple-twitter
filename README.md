# Simple Tweeter (Kaliaha Artsiom, s2110455009)

Das Ziel dieses Projekts bestand darin, eine extrem einfache Version von Twitter nachzumachen.

Übersicht über die verwendeten Technologien:
- Programmiersprache: Golang
- Web Framework: Gin
- Bibliotheken:
    + chi (Routing)
    + gin (Web-Framework)
    + uber zap (strukturiertes Logging)
    + viper (Konfigurationsmanagement)
    + gorm (ORM für Postgres)
    + gqlgen (Graphql Generator)
    + grpc & protobuf (RPC Calls)
    + service-engineering-go-lib (eine eigene Sammlung von Hilfsfunktionen)
    + jwt-go (Authentifizierung mittels JWT Tokens)
- Persistenz: MySQL (Userdaten), Postgres (Tweets)
- Deployment: Helm

### Architektur der Lösung

Die Architektur der Lösung besteht aus drei Services:
- Users Service: Registrierung der User, Ausstellung der JWT Tokens, Validierung der JWT Tokens. Dieser Service verwendet MySQL Datenbank.
- Tweets Service: Speicherung und Bereitstellung von Tweets. Tweets werden in der Postgres Datenbank gespeichert.
- Graphql Service (a.k.a "Backend for Frontend" oder "API-Gateway"). Dieser Service leitet Daten an entsprechende Services weiter, liefert die Responses zurück, wandelt die Kommunikationsprotokolle um, um mit den entsprechenden Services kommunizieren zu können. 

![1_architecture](/images/1_architecture.jpg)

### Kommunikation

Kommunikation erfolgt über folgende Schnittstellen:
- Users Service - Graphql Service: HTTP ist die einzige Schnittstelle
- Tweets Service - Graphql Service: gRPC als primärer Kommunikationskanal, HTTP als Fallback / sekundäre Schnittstelle.
 
`Users Service - Graphql Service: gRPC`
```proto
syntax = "proto3";

option go_package = "github.com/stakkato95/twitter-service-users/protoservice";

package protoservice;

service UsersService {
    rpc CreateUser(User) returns (NewUser);

    rpc AuthUser(User) returns (Token);

    rpc AuthUserByToken(Token) returns (User);
}

message User {
    int64 id = 1;
    string username = 2;
    string password = 3;
}

message Token {
    string token = 1;
}

message NewUser {
    User user = 1;
	Token token = 2;
}
```

`Users Service - Graphql Service: HTTP`
| Method | Routes        | Description                                                     |
| ------ | ------------- | --------------------------------------------------------------- |
| POST   | /debug/create | Einen User im System registrieren                               |
| POST   | /debug/auth   | Einen JWT Token für einen registroerten User ausstellen         |

`Tweets Service - Graphql Service: HTTP`
| Method | Routes          | Description                                                     |
| ------ | --------------- | --------------------------------------------------------------- |
| POST   | /tweets         | Einen Tweet erstellen                                           |
| GET    | /tweets/:userId | Alle Tweets eines Users abfragen                                |

### Datenmodell

Daten werden von Tweets und Users Services in eigenen Datenbanken verwaltet.

![6_data_model](/images/6_data_model.jpg)

### Kurzer Überblick über Technologien

### service-engineering-go-lib

Das ist eine von mir geschriebene nano-Bibliothek, die die Funktionen / Komponenten umfasst, die in der letzten Übung einfach von einem zum anderen Service dupliziert wurden. Diese Komponenten sind Logging (Wrapper für uber zap Bibliothek), Konfigurationsmanagement (Wrapper für Viper Bibliothek) und eine Funktion, die in HTTP Handlers einer HTTP Response erleichtert. Link: https://github.com/stakkato95/service-engineering-go-lib

![2_go_lib](/images/2_go_lib.jpg)

### gin

![13_gin](/images/13_gin.jpg)

In Go Community ist Gin das beliebteste Framework. In Microservices, die mit Go entwickelt werden, werden oft einfach Routers eingesetzt, aber Gin bietet zusätzlich Parametervalidierung, Serving statischer Dateien und mehrere Arten von Middleware. In meinem Projekt wurde Gin nicht in allen Services eingesetzt, nur im Tweets Service (im Users Service wird "chi" Router verwendet).

```go
func Start() {
	repo := domain.NewTweetsRepo()
	service := service.NewTweetsService(repo)

	h := TweetsHandler{service}

	router := gin.Default()
	router.POST("/tweets", h.addTweet)
	router.GET("/tweets/:userId", h.getTweets)
	router.Run(config.AppConfig.ServerPort)
}

type TweetsHandler struct {
	service service.TweetsService
}

func (h *TweetsHandler) addTweet(ctx *gin.Context) {
	var tweetDto dto.TweetDto
	if err := ctx.ShouldBindJSON(&tweetDto); err != nil {
		errorResponse(ctx, err)
		return
	}

	createdTweet := h.service.AddTweet(tweetDto)
	ctx.JSON(http.StatusOK, dto.ResponseDto{Data: *createdTweet})
}
```

### gqlgen (Graphql Generator)

gqlgen ermöglicht Generierung eines Graphql Services auf Basis des Graphql Schemas. Graphql Schema meines Backend-for-Frontend Services schaut wie folgt aus:

```graphql
type Tweet {
  id: Int!
  userId: Int!
  text: String!
}

type Query {
  tweets: [Tweet!]!
}

input NewUser {
  username: String!
  password: String!
}

input Login {
  username: String!
  password: String!
}

input NewTweet {
  text: String!
}

type Mutation {
  createUser(input: NewUser!): String!

  login(input: Login!): String!

  createTweet(input: NewTweet!): Tweet!
}
```

Mit dem oben dargestellten Schema kann man folgende Queries ausführen:

```graphql
mutation {
  createUser(input: {username: "user1", password: "pass"})
}

mutation {
  login(input: {username: "user1", password: "pass"})
}

mutation {
  createTweet(input: {userId: 1, text: "new tweet"}) {
    id
    userId
    text
  }
}

{
  tweets {
    id
    userId
    text
  }
}
```

Go Handler-Funktionen, die aus dem Schema erzeugt wurden, schauen so aus:
```go
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	return r.UserService.Create(input)
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	return r.UserService.Authenticate(input)
}

func (r *mutationResolver) CreateTweet(ctx context.Context, input model.NewTweet) (*model.Tweet, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return nil, errors.New("invalid authorization")
	}

	return r.TweetService.CreateTweet(input, int(user.Id))
}

func (r *queryResolver) Tweets(ctx context.Context) ([]*model.Tweet, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return nil, errors.New("invalid authorization")
	}

	return r.TweetService.GetTweets(int(user.Id))
}
```

### gorm

Im letzten Projekt (zu Message Oriented Middleware) wurden keine ORMs ausprobiert, die das Mapping von Entities auf Tabellen einer Datenbank ermöglichen. Einer der Gründe dafür war geringe beliebtheit von ORMs in Go Community. Viele Projekte verwenden einfach SQL Driver für entsprechende Datenbanken.

Dieses Mal wurde von mir eine der bekanntesten (und warscheinlich auch wenigen) ORM Bibliotheken ausprobiert, und zwar gorm. Gorm bietet alle gewöhnliche Funktionen eines ORMs, inklusive Migrationen, Select mit Where-Bediengungen, Updates usw.

```go
func NewTweetsRepo() TweetsRepo {
	db, err := gorm.Open(postgres.Open(config.AppConfig.DbSource), &gorm.Config{})
	db.AutoMigrate(&Tweet{})
	return &postgresTweetsRepo{db}
}

func (r *postgresTweetsRepo) AddTweet(tweet Tweet) *Tweet {
	r.db.Create(&tweet)
	return &tweet
}

func (r *postgresTweetsRepo) GetAllTweets(userId int) []Tweet {
	tweets := []Tweet{}
	r.db.Where("user_id = ?", userId).Find(&tweets)
	return tweets
}
```

### jwt-go

Diese Bibliothek bietet Heilfsfunktionen zur Erstellung und Validierung der JWT Tokens. Die Struktur eines Tokens im Falle meiner Anwendung schut so aus:

![7_jwt](/images/7_jwt.jpg)

Zusätzliche Hilfsfunktionen auf Basis jwt-go in meinem Projekt:

```go
var (
	SecretKey = []byte(config.AppConfig.JwtSecret)
)

func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}
``` 

### Packaging für Deployment

Als Packaging für den Service wurde wie bei allen anderen Projekten Docker verwendet (Dockerfile im Rootverzeichnis jedes einzelnen Services). Build eines Images wird durch Makefile gestartet (auch wie im letzten Projekt). Jeder Service hat sein eigenes Repository auf Docker Hub.

![3_dockerhub](/images/3_dockerhub.jpg)

### Deployment der Infrastruktur

Die Infrastruktur dieses Projekts (MySQL und Prostgres) wird mittels Helm (Package Manager für Kubernetes) im k8s installiert (k8s Cluster im Docker Desktop). Zwecks Vereinfachung wiederkehrender Operationen wurden insgesamt vier PowerShell Skripts geschrieben (jeweils zwei pro Datenbamk). Skripts im Verzeichnis `k8s\mysql`:

`install.ps1`
```powershell
$name = "mysql"

# 1
kubectl delete $(kubectl get pvc -l app.kubernetes.io/name=$name -o=name)

# 2
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install $name bitnami/mysql --set auth.rootPassword=root --set auth.database=users

# 3
echo "`nwaiting for pod to be ready..."
kubectl wait --for=condition=Ready $(kubectl get pod -l app.kubernetes.io/name=$name -o=name)
kubectl port-forward svc/$name 3306:3306
```

Operationen im install Skript:
1) Persistence Volume Claim der zuletzt angelegten Datenbank löschen. Dadurch wird automatisch Persistence Volume und folglich alle Daten der Datenbank entfernt.
2) Helm Repo, falls noch nicht hinzugefügt, hinzufügen
3) Abwarten bis der Pod im Redy zustand ist und dann den Port auf localhost mappen (zwecks Visualisierung der Daten in einer Database Viewer App)

`uninstall.ps1`
```powershell
$name = "mysql"

#1
helm uninstall $name

#2
kubectl delete $(kubectl get pvc -l app.kubernetes.io/name=$name -o=name)
```

Operationen im uninstall Skript:
1) Das Deployment der Datenbank mit Helm löschen
2) Anschließend (einfach zur Sicherheit) ebenfalls Persistence Volume Claim löschen

### Deployment der Infrastruktur

Die Anwendung wird dieses Mal nicht nur mit PowerShell Skripts, sondern auch (zum Teil) mit Helm deployed. Im Verzeichnis jeden einzelnen Services wurde mithilfe Helm "helm" Verzeichnis angelegt. In dem Verzeichnis liegen so genannte "Templates". Zu Templates zählen Deployments, Services, Ingress und alle andere Ressourcen, die man mit `kubecetl apply -f XYZ.yaml` im k8s Cluster erstellen kann. 

![4_helm](/images/4_helm.jpg)

Deployment hat folgende Struktur:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .Chart.Name }}-deployment
  labels:
    {{- include "helm.labels" . | nindent 4 }}
spec:
  replicas: 1
  selector:
    matchLabels:
      {{- include "helm.selectorLabels" . | nindent 6 }}
  template:
    metadata:
      labels:
        {{- include "helm.selectorLabels" . | nindent 8 }}
    spec:
      containers:
        - name: {{ .Chart.Name }}-container
          image: "{{ .Values.image.repository }}:{{ .Values.image.tag }}"
          imagePullPolicy: {{ .Values.image.pullPolicy }}
          ports:
            - name: {{ .Values.service.http.name }}-container
              containerPort: 8080
              protocol: TCP
```

Die Werte, die in eckigen Klammern `{{ }}` stehen, werden aus der Chart-Definition und aus der `values.yaml` Datei übernommen.

`Chart.yaml`
```yaml
apiVersion: v2
name: twitter-service-tweets
description: A Helm chart for twitter tweets service

home: https://github.com/stakkato95/service-engineering-simple-twitter

maintainers:
  - name: Artsiom Kaliaha
    email: stakkato95@gmail.com

type: application

# Chart version. Versions are expected to follow Semantic Versioning (https://semver.org/)
version: 0.1.0

# Application version
appVersion: "0.1.0"
```

`values.yaml`
```yaml
image:
  repository: stakkato95/twitter-service-tweets
  pullPolicy: Always
  tag: "latest"

nameOverride: ""
fullnameOverride: ""

service:
  type: ClusterIP
  http:
    name: http
    port: 80

ingress:
  enabled: true
```

Gerade bevor ein Helm Chart deployed wird, wird er kompiliert. Das Ergebniss für das Deployment:

`helm install tweets helm --dry-run`
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: twitter-service-tweets-deployment
  labels:
    helm.sh/chart: twitter-service-tweets-0.1.0
    app.kubernetes.io/name: twitter-service-tweets
    app.kubernetes.io/instance: tweets
    app.kubernetes.io/version: "0.1.0"
    app.kubernetes.io/managed-by: Helm
spec:
  replicas: 1
  selector:
    matchLabels:
      app.kubernetes.io/name: twitter-service-tweets
      app.kubernetes.io/instance: tweets
  template:
    metadata:
      labels:
        app.kubernetes.io/name: twitter-service-tweets
        app.kubernetes.io/instance: tweets
    spec:
      containers:
        - name: twitter-service-tweets-container
          image: "stakkato95/twitter-service-tweets:latest"
          imagePullPolicy: Always
          ports:
            - name: http-container
              containerPort: 8080
              protocol: TCP
```

Auf solche Art werden alle Services lokal deployed. Der Skript, der alle Services auf einmal installiert:

`app-install.ps1`
```powershell
cd ../twitter-service-graphql
helm install graphql helm
echo ""

cd ../twitter-service-tweets
helm install tweets helm
echo ""

cd ../twitter-service-users
helm install users helm

cd ../k8s
```

Anzahl der Pods, wenn Infrastruktur und alle Services deployed sind:

![5_pods](/images/5_pods.jpg)

### Testdurchlauf

Als Frontend für mein Projekt wird GraphiQL verwendet. 

Im ersten Schritt muss sich ein User im System anmelden. Das erfolgt mittels einer GraphQL Mutation Operation. Als Ergebniss bekommt man JWT Token. JWT Token muss man bei der Erstellung neuer Tweets oder beim Abfragen aller Tweets im Header mitgeben. In der Auth-Middleware wird der Token extrahiert und an Users Service geschickt, wo die Validierung erfolgt. GraphQL Service kriegt nur das Resultat der Validierung zurück (User oder Fehler Objekt). 

![8_create_user](/images/8_create_user.jpg)

Beim Verlust des Tokens kann ein neuer durch die Login-Operation ausgestellt werden.  

![9_login](/images/9_login.jpg)

Nach der Anmeldung kann ein User anfangen, Tweets zu posten. Dafür wird nur der Text des Tweets und der JWT Token im Auth-Header benötigt.

![10_create_tweet](/images/10_create_tweet.jpg)

Wenn keiner / falscher / ungültiger Token mitgegeben wird, bekommt User eine Fehlermeldung.

![11_auth_err](/images/11_auth_err.jpg)

Nach dem einige Tweets gepostet sind, können sie mittels einer Query abgefragt werden (dafür wird auch ein Token benötigt).

![12_all_tweets](/images/12_all_tweets.jpg)

### Con­clu­sio

Im vorliegenden Projekt wurde Folgendes ausprobiert:
- Graphql als "Frontend for Backend" Pattern
- Authentifizierung mittels JWT beim Graphql Server
- Go ORM gorm
- Go Web-Framework gin
- Packaging und Deployment der Microservices mittels Helm Charts
- Erstellung eigener Go Packages am Beispiel service-engineering-go-lib