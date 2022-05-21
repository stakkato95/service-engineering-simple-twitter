# service-engineering-simple-twitter
very simple version of twitter

besonderheit
- go-service-lib https://github.com/stakkato95/service-engineering-go-lib
- helm
- jwt https://jwt.io/

Dockerfile - soll im projekt bleiben

https://www.howtographql.com/graphql-go/1-getting-started/

go get google.golang.org/grpc
https://grpc.io/docs/languages/go/quickstart/
https://github.com/grpc/grpc-go/blob/master/examples/route_guide/routeguide/route_guide.proto
https://github.com/grpc/grpc-go/blob/9f4b31a11cc4deba7f5c542399d5ec71fab3a053/examples/route_guide/client/client.go#L48

docker run --rm -e POSTGRES_PASSWORD=root -e POSTGRES_USER=root -e POSTGRES_DB=tweets -p 5432:5432 -d postgres:latest

https://gorm.io/docs/index.html
