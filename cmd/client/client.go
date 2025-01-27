package client

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"main.go/internal"
)

const frameDelimiter = "END_OF_FRAME"

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

	// Initial handshake message
	if _, err := conn.Write([]byte("Hello from client")); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Handshake message sent to server")

	const chunkSize = 1024
	frameIndex := 1

	for {
		// Receive chunks for a single frame
		frameData, err := ReceiveChunks(chunkSize, conn)
		if err != nil {
			log.Printf("Error receiving frame %d: %v. Exiting...", frameIndex, err)
			break
		}

		// Save the received frame as an image
		internal.ByteToImage(frameData, strconv.Itoa(frameIndex))
		fmt.Printf("Frame %d saved successfully\n", frameIndex)

		frameIndex++
	}
}

func ReceiveChunks(chunkSize int, conn net.Conn) ([]byte, error) {
	frameData := make([]byte, 0, chunkSize*100) // Preallocate a reasonable buffer for the frame

	for {
		chunk := make([]byte, chunkSize)
		n, err := conn.Read(chunk)
		if err != nil {
			return nil, fmt.Errorf("error reading chunk: %v", err)
		}

		frameData = append(frameData, chunk[:n]...)
		fmt.Printf("Received chunk of size %d bytes\n", n)

		// Stop reading when the server signals the end of a frame (delimiter received)
		if n < chunkSize {
			// Check if delimiter is received
			if string(chunk[:n]) == frameDelimiter {
				fmt.Println("End of frame detected (delimiter received)")
				break
			}
		}
	}

	// Check if the frame data is valid before proceeding
	if len(frameData) == 0 {
		return nil, fmt.Errorf("received empty frame data")
	}

	fmt.Printf("Total received frame size: %d bytes\n", len(frameData))
	return frameData, nil
}
