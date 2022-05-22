cd ../twitter-service-graphql
helm install graphql helm
echo ""

cd ../twitter-service-tweets
helm install tweets helm
echo ""

cd ../twitter-service-users
helm install users helm

cd ../k8s