package main

type Position struct {
	X int
	Y int
}

type Direction int

const (
	DirUp Direction = iota
	DirRight
	DirDown
	DirLeft
)

type Rectangle struct {
	Pos    Position
	Width  int
	Height int
}

func (r *Rectangle) isColliding(r2 *Rectangle) bool {
	return r.Pos.X < r2.Pos.X+r2.Width &&
		r.Pos.X+r.Width > r2.Pos.X &&
		r.Pos.Y < r2.Pos.Y+r2.Height &&
		r.Pos.Y+r.Height > r2.Pos.Y
}

//x1      < x2 + w2 &&
//x1 + w1 > x2      &&
//y1      < y2 + h2 &&
//y1 + h1 > y2

// (4, 4)  .. (14, 4)       (14, 4)  .. (24, 4)
//   .           .             .
//   .           .             .
// (4, 14) .. (14, 14)      (14, 14) .. (24, 14)
