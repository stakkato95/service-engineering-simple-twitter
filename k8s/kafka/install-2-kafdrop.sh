# cd /mnt/c/Users/stakk/Documents/fh/master/2/service/service-engineering-simple-twitter/k8s/kafka

# https://artifacthub.io/packages/helm/main/kafdrop
# helm git repo https://github.com/bedag/helm-charts/tree/master/charts/kafdrop
# app git repo https://github.com/obsidiandynamics/kafdrop?msclkid=0befd7a0ceb111eca5cc40afafbcfbe9
helm repo add kafdrop https://bedag.github.io/helm-charts/
helm install kafdrop kafdrop/kafdrop --version 0.2.3 --set image.tag=3.31.0-SNAPSHOT --set config.kafka.connections=kafka.default.svc.cluster.local:9092 --set config.kafka.properties.content=`"{{ toString "sasl.jaas.config=org.apache.kafka.common.security.scram.PlainLoginModule required username='user' password='user';\nsecurity.protocol=PLAINTEXT" b64enc }}\n"`