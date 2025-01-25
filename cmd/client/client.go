package cmd

import (
	"log"
	"net"

	"main.go/internal"
)

func UDPDial() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		frameIndex, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		internal.ByteToImage(frameIndex)

	}
}
