package cmd

import (
	"fmt"
	"log"
	"net"
	"strconv"

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

	for i := 1; i <= internal.GetTotalImages("/home/rameez/Downloads/frametest/"); i++ {
		frame, err := internal.ImageToByte("/home/rameez/Downloads/frametest/frame_" + strconv.Itoa(i) + ".jpg")
		if err != nil {
			log.Fatal(err)
		}

		_, err = ln.WriteToUDP(frame, addr)
		if err != nil {
			log.Fatal(err)
		}
	}
}
