package cmd

import (
	"fmt"
	"log"
	"net"

	"main.go/internal"
)

func UDPListen() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	fmt.Println("Running on port :3000")

	for i := 0; i <= internal.GetTotalImages("/home/rameez/Downloads/frametest/"); i++ {

	}
}
