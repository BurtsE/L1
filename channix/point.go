package channix

type coord float32
type Point struct {
	x, y coord
}

func (p *Point) Point(x, y coord) {
	p.x = x
	p.y = y
}
