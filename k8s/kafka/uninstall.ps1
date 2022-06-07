helm uninstall kafka
helm uninstall kafdrop

kubectl delete $(kubectl get pvc -l app.kubernetes.io/name=kafka -o=name)
kubectl delete $(kubectl get pvc -l app.kubernetes.io/name=zookeeper -o=name)