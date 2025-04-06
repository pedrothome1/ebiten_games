package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Snake *Snake
}

func (g *Game) Update() error {
	return g.Snake.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Snake.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	g := &Game{
		Snake: NewSnake(),
	}

	ebiten.SetTPS(15)

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
