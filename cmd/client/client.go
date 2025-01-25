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

	var frameIndex []byte

	var frameData []byte

	for {
		_, err := conn.Read(frameIndex)
		if err != nil {
			log.Fatal(err)
		}

		_, err = conn.Read(frameData)
		if err != nil {
			log.Fatal(err)
		}

		internal.ByteToImage(frameData, frameIndex)

	}
}
