package cmd

import (
	"log"
	"net"
)

func UDPDial() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

}
