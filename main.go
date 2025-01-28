package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"path/filepath"

	"golang.org/x/image/draw"
)

func Preconvert(input, output string, size int) error {
	src, err := Load(input)
	if err != nil {
		return fmt.Errorf("failed to load an image: %w", err)
	}

	dst := image.NewRGBA(image.Rect(0, 0, size, size))
	draw.CatmullRom.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Src, nil)

	err = Save(output, dst)
	if err != nil {
		return err
	}
	return nil
}

func appendFilename(name, suffix string) string {
	ext := filepath.Ext(name)
	return name[0:len(name)-len(ext)] + suffix
}

func main() {
	var (
		input  string
		output string
		size   int
	)
	flag.IntVar(&size, "size", 224, "size of output: 224, 448, 896")
	flag.StringVar(&output, "output", "", "output name, default auto-generated")
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatalf("require one input image file")
	}

	input = flag.Arg(0)
	if output == "" {
		output = appendFilename(input, fmt.Sprintf("_%[1]ds.jpg", size))
	}

	if err := Preconvert(input, output, size); err != nil {
		log.Fatal(err)
	}
}
