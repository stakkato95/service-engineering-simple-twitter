# https://artifacthub.io/packages/helm/bitnami/kafka
# git repo https://github.com/bitnami/charts/tree/master/bitnami/kafka
# helm formatting arrays https://newbedev.com/helm-passing-array-values-through-set#:~:text=Helm%3A%20Passing%20array%20values%20through%20--set%20If%20you,require%20quotes%29%3A%20--set%20test%3D%20%7Bx%2Cy%2Cz%7D%20--set%20%22test%3D%20%7Bx%2Cy%2Cz%7D%22
kubectl delete $(kubectl get pvc -l app.kubernetes.io/name=kafka -o=name)
kubectl delete $(kubectl get pvc -l app.kubernetes.io/name=zookeeper -o=name)

helm repo add bitnami https://charts.bitnami.com/bitnami
helm install kafka bitnami/kafka --set auth.clientProtocol=plaintext --set auth.sasl.jaas.clientUsers="{user}" --set auth.sasl.jaas.clientPasswords="{user}" --set deleteTopicEnable=true

echo "`nwaiting for kafka pod to be ready..."
kubectl wait --for=condition=Ready $(kubectl get pod -l app.kubernetes.io/name=kafka -o=name)
echo "`nwaiting for zookeeper pod to be ready..."
kubectl wait --for=condition=Ready $(kubectl get pod -l app.kubernetes.io/name=zookeeper -o=name)