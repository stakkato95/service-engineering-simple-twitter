kubectl delete $(kubectl get pvc -o=name | Select-String "data-mysql-0")

helm repo add bitnami https://charts.bitnami.com/bitnami
helm install mysql bitnami/mysql --set auth.rootPassword=root --set auth.database=users

echo "`nwaiting for pod to be ready..."
kubectl wait --for=condition=Ready pod/mysql-0
kubectl port-forward svc/mysql 3306:3306