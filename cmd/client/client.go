package client

import (
	"fmt"
	"log"
	"net"
	"strconv"

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

	var frameData = make([]byte, 614400)

	if _, err := conn.Write([]byte("Hello from client")); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Message sent")
	i := 1
	for {
		_, err = conn.Read(frameData)
		if err != nil {
			log.Fatal(err)
		}

		frameIndex := strconv.Itoa(i)
		internal.ByteToImage(frameData, frameIndex)
		i++
	}
}
