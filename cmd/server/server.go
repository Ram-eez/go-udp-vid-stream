package server

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

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

	var buf [1024]byte
	_, clientAddr, err := ln.ReadFromUDP(buf[:])
	if err != nil {
		log.Fatal(err)
	}

	go internal.FFmpegFrameCapture()
	time.Sleep(time.Second * 5)

	const chunkSize = 1024
	i := 1
	for {

		frame, err := internal.ImageToByte("/home/rameez/Downloads/framecreation/frame_" + strconv.Itoa(i) + ".jpg")
		if err != nil {
			log.Fatal(err)
			break
		}

		// sending frameData
		frameSize := len(frame)

		_, err = ln.WriteToUDP([]byte(strconv.Itoa(frameSize)), clientAddr)
		if err != nil {
			log.Fatal(err)
		}
		totalChunks := (frameSize + chunkSize - 1) / chunkSize
		for chunkIndex := 0; chunkIndex < totalChunks; chunkIndex++ {
			start := chunkIndex * chunkSize
			end := start + chunkSize

			if end > frameSize {
				end = frameSize
			}

			_, err := ln.WriteToUDP(frame[start:end], clientAddr)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Sent chunk %d/%d\n", chunkIndex+1, totalChunks)
		}

		fmt.Println("frame sent with data")
		i++
	}
}
