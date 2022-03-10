package internal

import (
	"github.com/disintegration/imaging"
	log "github.com/sirupsen/logrus"
	"image"
	"image/color"
	"image/draw"
)

func OpenSideImage(configuration VideoConfiguration, imageName string) image.Image {
	src, err := imaging.Open(imageName)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	targetWidth := int(float64(configuration.Width) / 3.33)
	resizedImage := imaging.Resize(src, targetWidth, 0, imaging.Lanczos)
	croppedImage := imaging.CropAnchor(resizedImage, int(configuration.Width), int(configuration.Height), imaging.Center)

	return croppedImage
}

func PasteWithTransparency(background image.Image, transparentImage image.Image, offset image.Point, alpha float64) image.Image {
	canvas := image.NewRGBA(background.Bounds())
	draw.Draw(canvas, canvas.Bounds(), background, image.Point{0, 0}, draw.Src)

	mask := image.NewUniform(color.Alpha{A: uint8(alpha * float64(255))})
	positionToPaste := image.Rectangle{offset, offset.Add(transparentImage.Bounds().Size())}

	draw.DrawMask(canvas, positionToPaste, transparentImage, image.Point{0, 0}, mask, image.Point{0, 0}, draw.Over)

	return canvas
}

func CloneToRGBA(src image.Image) draw.Image {
	bounds := src.Bounds()
	output := image.NewRGBA(bounds)
	draw.Draw(output, bounds, src, bounds.Min, draw.Src)
	return output
}
