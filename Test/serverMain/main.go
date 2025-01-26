package main

import (
	"time"

	"main.go/cmd/server"
)

func main() {

	server.UDPListen()
	time.Sleep(time.Second * 10)

}
