package internal

import (
	log "github.com/sirupsen/logrus"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func GenerateVideo(configuraton VideoConfiguration, inputFile string, outputFile string) {
	outputArguments := ffmpeg.KwArgs{
		"c:v":       "libx264",
		"pix_fmt":   "yuv420p",
		"framerate": configuraton.Fps,
		"preset":    "slow",
	}

	err := ffmpeg.Input(inputFile).
		Output(outputFile, outputArguments).
		OverWriteOutput().ErrorToStdOut().Run()

	if err != nil {
		log.Fatalf("failed to transcode video: %v", err)
	}
}
