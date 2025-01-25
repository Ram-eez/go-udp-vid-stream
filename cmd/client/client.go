package client

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

	log.Println("Connected to server at", addr)
	defer conn.Close()

	var frameIndex = make([]byte, 1024*5)
	var frameData = make([]byte, 65536)
	for {
		n, err := conn.Read(frameIndex)
		if err != nil {
			log.Printf("Error reading frame index: %v\n", err)
			break
		}
		log.Printf("Received frame index: %s\n", string(frameIndex[:n]))

		_, err = conn.Read(frameData)
		if err != nil {
			log.Fatal(err)
		}

		internal.ByteToImage(frameData, frameIndex)

	}
}
