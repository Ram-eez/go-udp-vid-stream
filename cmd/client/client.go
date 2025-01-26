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

	if _, err := conn.Write([]byte("Hello from client")); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Message sent")
	buf := make([]byte, 1024)
	const chunkSize = 1024
	var frameData []byte
	i := 1
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)

		}

		frameData = frameData[:0]
		frameData = ReceiveChunks(chunkSize, frameData, len(buf[:n]), conn)

		frameIndex := strconv.Itoa(i)
		internal.ByteToImage(frameData, frameIndex)
		i++
	}
}

func ReceiveChunks(chunkSize int, frameData []byte, Data int, conn net.Conn) []byte {
	totalChunks := (Data + chunkSize - 1) / chunkSize
	for chunkIndex := 0; chunkIndex < totalChunks; chunkIndex++ {
		chunk := make([]byte, chunkSize)
		n, err := conn.Read(chunk)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		frameData = append(frameData, chunk[:n]...)
		fmt.Printf("Received chunk %d/%d\n", chunkIndex+1, totalChunks)
	}
	return frameData
}
