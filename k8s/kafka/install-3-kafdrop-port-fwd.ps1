$kafdrop = kubectl get po -o=name | Select-String "kafdrop"

echo "`nwaiting for kafdrop pod to be ready..."
kubectl wait --for=condition=Ready $kafdrop

kubectl port-forward $kafdrop 9000:9000