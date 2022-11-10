package ytdownloader

import (
	"io"
	"os"
	"github.com/kkdai/youtube/v2"
	"github.com/google/uuid"
	"log"
)

func DownloadAndGetPath() string{
	videoID := "zGDzdps75ns"
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}

	log.Printf("[YT] video stream %s opened!", videoID)
	
	videoName := "static/" + uuid.New().String()[0:8] + ".mp3"
	file, err := os.Create(videoName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}

	log.Printf("[YT] video %s downloaded fully!", videoID)

	return videoName
}

