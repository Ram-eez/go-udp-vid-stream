package internal

import (
	"fmt"
	"log"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func FFmpegFrameCapture() {
	err := ffmpeg_go.Input("/home/rameez/Downloads/New Folder/DespicableMe/DespicableMe.mp4").
		Output("/home/rameez/Downloads/frametest/frame_%02d.jpg", ffmpeg_go.KwArgs{"vf": "fps=1", "q:v": "1"}).
		Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Frames Extracted Successfully")
}
