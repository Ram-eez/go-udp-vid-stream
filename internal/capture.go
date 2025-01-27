package internal

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func FFmpegFrameCapture() {
	err := ffmpeg_go.Input("/home/rameez/Downloads/New Folder/DespicableMe/DespicableMe.mp4").
		Output("/home/rameez/Downloads/framecreation/frame_%d.jpg", ffmpeg_go.KwArgs{"vf": "fps=1", "q:v": "1"}).
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

// ByteToImage processes received frame data and saves it as a JPEG image.
func ByteToImage(frameData []byte, frameIndex string) {
	fmt.Printf("Received %d bytes of image data\n", len(frameData))

	// Verify frame data size
	if len(frameData) < 10 {
		log.Printf("Frame %s is too small to process\n", frameIndex)
		return
	}
	fmt.Printf("First 10 bytes of frame data: %x\n", frameData[:10])

	// Check for JPEG header (SOI marker 0xFFD8)
	if !hasJPEGHeader(frameData) {
		log.Printf("Skipping frame %s: Missing JPEG header\n", frameIndex)
		return
	}

	// Decode the JPEG image from frame data
	img, _, err := image.Decode(bytes.NewReader(frameData))
	if err != nil {
		log.Printf("Skipping invalid frame at index %s: %v\n", frameIndex, err)
		return
	}

	// Create output file for the frame
	outputPath := fmt.Sprintf("/home/rameez/Downloads/reception/frame_%s.jpg", frameIndex)
	frameFile, err := os.Create(outputPath)
	if err != nil {
		log.Printf("Failed to create file for frame %s: %v\n", frameIndex, err)
		return
	}
	defer frameFile.Close()

	// Encode and save the image as JPEG
	if err := jpeg.Encode(frameFile, img, nil); err != nil {
		log.Printf("Failed to save JPEG for frame %s: %v\n", frameIndex, err)
		return
	}
	fmt.Printf("Successfully saved frame to: %s\n", outputPath)
}

// hasJPEGHeader checks if the given data starts with a JPEG SOI marker (0xFFD8).
func hasJPEGHeader(data []byte) bool {
	return len(data) >= 2 && data[0] == 0xFF && data[1] == 0xD8
}
