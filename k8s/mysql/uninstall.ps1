$name = "mysql"
helm uninstall $name
kubectl delete $(kubectl get pvc -l app.kubernetes.io/name=$name -o=name)