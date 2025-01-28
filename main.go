package main

import (
	"image"
	"image/png"
	"log"
	"os"

	"golang.org/x/image/draw"
)

func loadImage(name string) (image.Image, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func pngSave(img image.Image, name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	e := png.Encoder{CompressionLevel: png.BestCompression}
	return e.Encode(f, img)
}

func main() {
	src, err := loadImage("tmp/in001.png")
	if err != nil {
		log.Fatal("failed to load image: %s", err)
	}
	dst := image.NewRGBA(image.Rect(0, 0, 224, 224))
	draw.CatmullRom.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Src, nil)
	err = pngSave(dst, "tmp/out001go.png")
	if err != nil {
		log.Fatal("failed to save an image: %s", err)
	}
}
