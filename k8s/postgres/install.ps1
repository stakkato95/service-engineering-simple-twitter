$name = "postgresql"

kubectl delete $(kubectl get pvc -l app.kubernetes.io/name=$name -o=name)

helm repo add bitnami https://charts.bitnami.com/bitnami
helm install $name bitnami/postgresql --set global.postgresql.auth.username=root --set global.postgresql.auth.password=root --set global.postgresql.auth.database=tweetsdb

echo "`nwaiting for pod to be ready..."
kubectl wait --for=condition=Ready $(kubectl get pod -l app.kubernetes.io/name=$name -o=name)
kubectl port-forward svc/$name 5432:5432