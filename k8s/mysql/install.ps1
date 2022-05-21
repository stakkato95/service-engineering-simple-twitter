$name = "mysql"

kubectl delete $(kubectl get pvc -l app.kubernetes.io/name=$name -o=name)

helm repo add bitnami https://charts.bitnami.com/bitnami
helm install $name bitnami/mysql --set auth.rootPassword=root --set auth.database=users

echo "`nwaiting for pod to be ready..."
kubectl wait --for=condition=Ready $(kubectl get pod -l app.kubernetes.io/name=$name -o=name)
kubectl port-forward svc/$name 3306:3306