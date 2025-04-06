package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

const (
	BlockSize = 10
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
	}

	for i := len(s.Body) - 1; i >= 1; i-- {
		s.Body[i].Pos.X = s.Body[i-1].Pos.X
		s.Body[i].Pos.Y = s.Body[i-1].Pos.Y
	}

	if s.Dir == DirUp {
		s.Body[0].Pos.Y -= BlockSize
	} else if s.Dir == DirRight {
		s.Body[0].Pos.X += BlockSize
	} else if s.Dir == DirDown {
		s.Body[0].Pos.Y += BlockSize
	} else if s.Dir == DirLeft {
		s.Body[0].Pos.X -= BlockSize
	}

	if s.Body[0].Pos.X > ScreenWidth {
		s.Body[0].Pos.X = 0
	} else if s.Body[0].Pos.X < 0 {
		s.Body[0].Pos.X = ScreenWidth
	}

	if s.Body[0].Pos.Y > ScreenHeight {
		s.Body[0].Pos.Y = 0
	} else if s.Body[0].Pos.Y < 0 {
		s.Body[0].Pos.Y = ScreenHeight
	}

	return nil
}

func (s *Snake) Draw(screen *ebiten.Image) {
	for _, square := range s.Body {
		snakeBlock := ebiten.NewImage(10, 10)
		snakeBlock.Fill(colornames.Limegreen)

		opt := &ebiten.DrawImageOptions{}

		opt.GeoM.Translate(square.Pos.X, square.Pos.Y)

		screen.DrawImage(snakeBlock, opt)
	}
}

func (s *Snake) SelfBite() bool {
	for i := len(s.Body) - 1; i >= 1; i-- {
		if s.Body[0].isColliding(s.Body[i]) {
			return true
		}
	}
	return false
}

func (s *Snake) Head() *Rectangle {
	return s.Body[0]
}

func (s *Snake) Grow() {
	newBlock := &Rectangle{
		Width:  BlockSize,
		Height: BlockSize,
	}

	s.Body = append(s.Body, newBlock)
}
