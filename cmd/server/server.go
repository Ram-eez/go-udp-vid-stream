package server

import (
	"fmt"
	"log"
	"net"
	"time"

	"main.go/internal"
)

const frameDelimiter = "END_OF_FRAME" // Define the frame delimiter

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

	fmt.Println("Server is running on port :3000")

	var buf [1024]byte
	_, clientAddr, err := ln.ReadFromUDP(buf[:])
	if err != nil {
		log.Fatal(err)
	}

	// Start FFmpeg frame capture in a goroutine
	go internal.FFmpegFrameCapture()

	// Wait for frames to be generated
	time.Sleep(time.Second * 5)

	const chunkSize = 1024
	frameIndex := 1

	for {
		// Load the frame file as bytes
		framePath := fmt.Sprintf("/home/rameez/Downloads/framecreation/frame_%d.jpg", frameIndex)
		frame, err := internal.ImageToByte(framePath)
		if err != nil {
			log.Printf("Error loading frame %d: %v. Stopping...", frameIndex, err)
			break
		}

		// Calculate frame size dynamically based on the actual frame
		frameSize := len(frame)
		totalChunks := (frameSize + chunkSize - 1) / chunkSize

		// Send the frame data in chunks
		for chunkIndex := 0; chunkIndex < totalChunks; chunkIndex++ {
			start := chunkIndex * chunkSize
			end := start + chunkSize
			if end > frameSize {
				end = frameSize
			}

			// Write the chunk to the client
			_, err := ln.WriteToUDP(frame[start:end], clientAddr)
			if err != nil {
				log.Printf("Error sending chunk %d/%d of frame %d: %v", chunkIndex+1, totalChunks, frameIndex, err)
				break
			}

			fmt.Printf("Sent chunk %d/%d of frame %d\n", chunkIndex+1, totalChunks, frameIndex)
		}

		// Send the delimiter to mark the end of the frame
		_, err = ln.WriteToUDP([]byte(frameDelimiter), clientAddr)
		if err != nil {
			log.Printf("Error sending delimiter for frame %d: %v", frameIndex, err)
			break
		}

		fmt.Printf("Frame %d sent successfully with delimiter\n", frameIndex)
		frameIndex++
	}
}
