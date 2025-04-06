package main

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font/opentype"
	"image"
	_ "image/png"
	"io"
	"io/fs"
)

//go:embed assets/*
var assets embed.FS

var (
	FruitSprites = mustLoadImages("assets/fruits/*.png")
	VT323Font    = mustLoadFont("assets/VT323/VT323-Regular.ttf")
)

func mustLoadImage(path string) *ebiten.Image {
	f, err := assets.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func mustLoadImages(globPat string) (images []*ebiten.Image) {
	matches, err := fs.Glob(assets, globPat)
	if err != nil {
		panic(err)
	}

	for _, m := range matches {
		images = append(images, mustLoadImage(m))
	}

	return images
}

func mustLoadFont(path string) *opentype.Font {
	f, err := assets.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	parsedFont, err := opentype.ParseReaderAt(f.(io.ReaderAt))
	if err != nil {
		panic(err)
	}

	return parsedFont
}
