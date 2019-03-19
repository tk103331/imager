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
	doTest("Imager", func(img image.Image) image.Image {
		return New(img)
	})
}

func TestCircleImager(t *testing.T) {
	doTest("CircleImager", func(img image.Image) image.Image {
		return New(img).Circle(CircleLonger)
	})
}

func TestCropImager(t *testing.T) {
	doTest("CropImager", func(img image.Image) image.Image {
		return New(img).Crop(image.Rect(75, 50, 225, 150))
	})
}

func TestFlipImager(t *testing.T) {
	doTest("FlipImager", func(img image.Image) image.Image {
		return New(img).Flip(FlipHorizontal)
	})
}

func TestScaleImager(t *testing.T) {
	doTest("ScaleImager", func(img image.Image) image.Image {
		return New(img).Scale(0.55)
	})
}

func TestRoundImager(t *testing.T) {
	doTest("RoundImager", func(img image.Image) image.Image {
		return New(img).Round(50)
	})
}

func TestBlurImager(t *testing.T) {
	doTest("BlurImager", func(img image.Image) image.Image {
		return New(img).Blur(3)
	})
}

func TestCropFlipCircle(t *testing.T) {
	doTest("CropFlipCircle", func(img image.Image) image.Image {
		return New(img).Crop(image.Rect(75, 50, 225, 150)).Circle(CircleOuter).Flip(FlipHorizontal)
	})
}
