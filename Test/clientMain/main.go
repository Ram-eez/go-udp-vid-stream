package main

import (
	"time"

	"main.go/cmd/client"
)

func main() {
	client.UDPDial()
	time.Sleep(time.Second * 100)
}
