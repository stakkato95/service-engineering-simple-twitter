# https://artifacthub.io/packages/helm/bitnami/mysql
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install mysql2 bitnami/mysql --set auth.rootPassword=root --set auth.database=users
# wait for running state
# kubectl port-forward svc/mysql2 3306:3306