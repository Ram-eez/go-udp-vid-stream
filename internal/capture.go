package internal

import (
	"fmt"
	"log"
	"os"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func FFmpegFrameCapture() {
	err := ffmpeg_go.Input("/home/rameez/Downloads/New Folder/DespicableMe/DespicableMe.mp4").
		Output("/home/rameez/Downloads/frametest/frame_%02d.jpg", ffmpeg_go.KwArgs{"vf": "fps=1", "q:v": "1"}).
		// add these lines in ffmpeg_go.KwArgs{} if you want to encode the file into H.264 codex
		// "c:v": "libx264"
		// "preset": "fast",      // Encoding speed (fast, medium, slow)
		Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Frames Extracted Successfully")
}

func ImageToByte(frame string) ([]byte, error) {
	byteStream, err := os.ReadFile(frame)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return byteStream, nil
}

func GetTotalImages(path string) int {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	return len(files)
}
