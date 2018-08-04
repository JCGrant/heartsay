package renderer

type Rect struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func NewRect(x1, y1, x2, y2 int) Rect {
	return Rect{x1, y1, x2, y2}
}
