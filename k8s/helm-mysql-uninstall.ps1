helm uninstall mysql
kubectl delete $(kubectl get pvc -o=name | Select-String "data-mysql-0")