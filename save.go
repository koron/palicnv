package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// Save saves an image as file.
// Save in the appropriate format depending on the file extension.
func Save(name string, img image.Image) error {
	ext := strings.ToLower(filepath.Ext(name))
	switch ext {
	case ".gif":
		return gifSave(name, img)
	case ".jpg", ".jpeg":
		return jpgSave(name, img)
	case ".png":
		return pngSave(name, img)
	default:
		return fmt.Errorf("unknown image format for ext: %q", ext)
	}
}

func pngSave(name string, img image.Image) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	e := png.Encoder{CompressionLevel: png.BestCompression}
	return e.Encode(f, img)
}

func gifSave(name string, img image.Image) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	return gif.Encode(f, img, nil)
}

func jpgSave(name string, img image.Image) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	opts := jpeg.Options{Quality: 100}
	return jpeg.Encode(f, img, &opts)
}
