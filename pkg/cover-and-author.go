package pkg

import (
	"fmt"
	"github.com/disintegration/imaging"
	log "github.com/sirupsen/logrus"
	"image"
	"podcast-animations/internal"
	"sync"
)

func CoverAndAuthor(videoConfiguration internal.VideoConfiguration, backgroundFileName string, coverFileName string, authorFileName string) {
	internal.PrepareTempPaths()

	background, _ := imaging.Open(backgroundFileName)
	book := internal.OpenSideImage(videoConfiguration, coverFileName)
	author := internal.OpenSideImage(videoConfiguration, authorFileName)

	numberOfFrames := int(videoConfiguration.Fps * videoConfiguration.DurationInSeconds)
	var wg sync.WaitGroup

	for i := 0; i < numberOfFrames; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()

			generateFrame(
				videoConfiguration,
				internal.CloneToRGBA(background),
				internal.CloneToRGBA(book),
				internal.CloneToRGBA(author),
				uint16(i))
		}()
	}
	wg.Wait()
	internal.GenerateVideo(videoConfiguration, "tmp/frame%04d.png", "output/video.mp4")

	internal.CleanupTempFiles()
}

func generateFrame(configuration internal.VideoConfiguration, canvas image.Image, book image.Image, author image.Image, frameNumber uint16) {
	log.Infof("Generating frame %d", frameNumber)

	coverAlpha := alphaForFrame(configuration, frameNumber, uint16(configuration.Fps/2))
	canvas = internal.PasteWithTransparency(canvas, book, image.Point{X: 192, Y: 100}, coverAlpha)

	authorAlpha := alphaForFrame(configuration, frameNumber, uint16(configuration.Fps*2))
	canvas = internal.PasteWithTransparency(canvas, author, image.Point{X: 1100, Y: 100}, authorAlpha)

	fileName := fmt.Sprintf("tmp/frame%04d.png", frameNumber)

	err := imaging.Save(canvas, fileName)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
	log.Infof("Frame %d completed", frameNumber)
}

func alphaForFrame(configuration internal.VideoConfiguration, frame uint16, offset uint16) float64 {
	startingPoint := uint16(configuration.Fps) + offset
	finishingPoint := uint16(configuration.Fps*2) + offset
	alpha := float64(1)
	if frame < startingPoint {
		alpha = 0
	} else if frame >= startingPoint && frame < finishingPoint {
		alpha = float64(frame-startingPoint) / float64(configuration.Fps)
	}
	return alpha
}
