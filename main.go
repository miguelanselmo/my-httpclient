package main

import (
	"time"

	"github.com/miguelanselmo/my-httpclient/demo"
)

func main() {
	auth := demo.Authentication{UserName: "TestUser3", Password: "123456"}
	demo.Auth(auth)
	demo.Demo()
	time.Sleep(time.Second * 10)
}
