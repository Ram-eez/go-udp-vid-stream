package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

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
	// buf := make([]byte, 1024)
	const chunkSize = 1024

	i := 1
	for {
		reader := bufio.NewReader(conn)
		frameSizeStr, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		frameSizeStr = strings.TrimSpace(frameSizeStr)
		fmt.Printf("Received frame size string: %s\n", frameSizeStr)

		frameSize, err := strconv.Atoi(frameSizeStr)
		if err != nil {
			log.Printf("Invalid frame size: %s", err)
			continue
		}

		var frameData []byte
		frameData = ReceiveChunks(chunkSize, frameData, frameSize, conn)

		frameIndex := strconv.Itoa(i)
		internal.ByteToImage(frameData, frameIndex)
		i++
	}
}

func ReceiveChunks(chunkSize int, frameData []byte, frameSize int, conn net.Conn) []byte {
	totalChunks := (frameSize + chunkSize - 1) / chunkSize
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
