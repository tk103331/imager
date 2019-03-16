package imager

import (
	"image"
	"image/png"
	"os"
	"testing"
)

func doTest(name string, creater func(img image.Image) image.Image) {
	oldFile, _ := os.Open("test.png")
	defer oldFile.Close()
	oldImg, _, _ := image.Decode(oldFile)

	newImg := creater(oldImg)
	newFile, _ := os.Create(name + ".png")
	defer newFile.Close()
	png.Encode(newFile, newImg)
}

func TestImager(t *testing.T) {
	doTest("Imager", func(image image.Image) image.Image {
		return New(image)
	})
}

func TestCircleImager(t *testing.T) {
	doTest("CircleImager", func(image image.Image) image.Image {
		return New(image).Circle(CircleShorter)
	})
}
