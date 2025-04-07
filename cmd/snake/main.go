package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

type Game struct {
	Snake *Snake
	Fruit *Fruit
	Over  bool
}

func (g *Game) Update() error {
	if g.Over {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.Snake = NewSnake()
			g.Fruit = NewFruit()
			g.Over = false
		}
		return nil
	}

	if g.Snake.Head().isColliding(g.Fruit.Rect) {
		MustPlayWavSound(ChewingSound)
		g.Snake.Grow()
		g.Fruit = NewFruit()
	}

	err := g.Snake.Update()
	if err != nil {
		return err
	}

	if g.Snake.SelfBite() {
		g.Over = true
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if !g.Over {
		g.Snake.Draw(screen)
		g.Fruit.Draw(screen)
		return
	}

	titleOpt := &text.DrawOptions{}
	titleOpt.GeoM.Translate(ScreenWidth/2-96, ScreenHeight/2-48)
	text.Draw(screen, "Game over!", MustBeGoXFaceWithSize(VT323Font, 48), titleOpt)

	msgOpt := &text.DrawOptions{}
	msgOpt.GeoM.Translate(ScreenWidth/2-92, ScreenHeight/2)
	text.Draw(screen, "Press ENTER to try again.", MustBeGoXFaceWithSize(VT323Font, 18), msgOpt)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	g := &Game{
		Snake: NewSnake(),
		Fruit: NewFruit(),
	}

	ebiten.SetTPS(15)

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
