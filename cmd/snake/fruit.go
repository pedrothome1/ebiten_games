package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

const (
	FruitScale = 0.08
)

type Fruit struct {
	Sprite *ebiten.Image
	Pos    Position
	Rect   *Rectangle
}

func NewFruit() *Fruit {
	sprite := FruitSprites[rand.Intn(len(FruitSprites))]

	fruitW := float64(sprite.Bounds().Dx()) * FruitScale
	fruitH := float64(sprite.Bounds().Dy()) * FruitScale
	padding := 20

	x := float64(rand.Intn(ScreenWidth-padding*2-int(fruitW)) + padding)
	y := float64(rand.Intn(ScreenHeight-padding*2-int(fruitH)) + padding)
	pos := Position{x, y}

	return &Fruit{
		Sprite: sprite,
		Pos:    pos,
		Rect: &Rectangle{
			Pos:    pos,
			Width:  fruitW,
			Height: fruitH,
		},
	}
}

func (f *Fruit) Update() error {
	return nil
}

func (f *Fruit) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(FruitScale, FruitScale)
	opt.GeoM.Translate(f.Pos.X, f.Pos.Y)
	screen.DrawImage(f.Sprite, opt)
}
