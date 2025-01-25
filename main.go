package main

import (
	"time"

	"main.go/cmd/client"
	"main.go/cmd/server"
)

func main() {
	go server.UDPListen()

	// time.Sleep(time.Second * 1)

	client.UDPDial()

	time.Sleep(time.Second * 10)

}
