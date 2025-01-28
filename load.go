package main

import (
	"image"
	"image/gif"
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/koron/palicnv/internal/gifutil"
)

// Load loads an image from file.
func Load(name string) (image.Image, error) {
	ext := strings.ToLower(filepath.Ext(name))
	switch ext {
	case ".gif":
		return loadGifRep(name)
	default:
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
}

func loadGif(name string) (*gif.GIF, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return gif.DecodeAll(f)
}

type frameInfo struct {
	i       int
	img     image.Image
	entropy float64
}

func paletteEntropy(img *image.Paletted) float64 {
	// calculate histogram of image.Paletted
	hist := map[uint8]int{}
	for _, p := range img.Pix {
		hist[p]++
	}

	sum := len(img.Pix)
	var entropy float64
	for _, n := range hist {
		p := float64(n) / float64(sum)
		entropy += -p * math.Log2(p)
	}
	return entropy
}

// loadGifRep loads an animation GIF image, and extract a representative frame.
func loadGifRep(name string) (image.Image, error) {
	g, err := gifutil.LoadFile(name)
	if err != nil {
		return nil, err
	}

	highest := frameInfo{i: -1, entropy: -1}
	for i, img := range gifutil.IterateComposed(g) {
		entropy := paletteEntropy(img)
		if entropy > highest.entropy {
			highest = frameInfo{i: i, img: img, entropy: entropy}
		}
	}

	return highest.img, nil
}
