package main

import (
	_ "github.com/stakkato95/service-engineering-go-lib/logger"
	"github.com/stakkato95/twitter-service-users/app"
	"github.com/stakkato95/twitter-service-users/protoapp"
)

func main() {
	go protoapp.Start()
	app.Start()
}
