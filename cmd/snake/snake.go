package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
)

const (
	Speed     = 10
	BlockSize = 10
)

var (
	Green = color.RGBA{2, 245, 27, 255}
	Red   = color.RGBA{255, 0, 0, 255}

	CurrentColor = Green
)

type Snake struct {
	Body []*Rectangle
	Dir  Direction
}

func NewSnake() *Snake {
	body := make([]*Rectangle, 0, 20)
	body = append(body, &Rectangle{
		Width:  BlockSize,
		Height: BlockSize,
	})

	return &Snake{
		Body: body,
		Dir:  DirRight,
	}
}

func (s *Snake) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyUp) && s.Dir != DirDown {
		s.Dir = DirUp
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) && s.Dir != DirLeft {
		s.Dir = DirRight
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) && s.Dir != DirUp {
		s.Dir = DirDown
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) && s.Dir != DirRight {
		s.Dir = DirLeft
	} else if inpututil.IsKeyJustPressed(ebiten.KeySpace) { // test block add
		s.addBodyBlock()
	}

	for i := len(s.Body) - 1; i >= 1; i-- {
		s.Body[i].Pos.X = s.Body[i-1].Pos.X
		s.Body[i].Pos.Y = s.Body[i-1].Pos.Y
	}

	if s.Dir == DirUp {
		s.Body[0].Pos.Y -= Speed
	} else if s.Dir == DirRight {
		s.Body[0].Pos.X += Speed
	} else if s.Dir == DirDown {
		s.Body[0].Pos.Y += Speed
	} else if s.Dir == DirLeft {
		s.Body[0].Pos.X -= Speed
	}

	// TEMP: just testing collision detection
	for i := 1; i < len(s.Body); i++ {
		if s.Body[0].isColliding(s.Body[i]) {
			CurrentColor = Red
			break
		}
	}

	return nil
}

func (s *Snake) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Blocks: %d", len(s.Body)))

	for _, square := range s.Body {
		snakeBlock := ebiten.NewImage(10, 10)
		snakeBlock.Fill(CurrentColor)

		opt := &ebiten.DrawImageOptions{}

		opt.GeoM.Translate(float64(square.Pos.X), float64(square.Pos.Y))

		screen.DrawImage(snakeBlock, opt)
	}
}

func (s *Snake) addBodyBlock() {
	newBlock := &Rectangle{
		Width:  BlockSize,
		Height: BlockSize,
	}

	s.Body = append(s.Body, newBlock)
}
