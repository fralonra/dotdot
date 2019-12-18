package dotdot

import "image/color"

type Dot struct {
	captured bool
	color    color.Color
	dir      [2]float64
	distance float64
	pos      [2]float64
	size     float64
	speed    float64
}

func (d *Dot) Captured() bool {
	return d.captured
}

func (d *Dot) Color() color.Color {
	return d.color
}

func (d *Dot) Distance() float64 {
	return d.distance
}

func (d *Dot) Dir() [2]float64 {
	return d.dir
}

func (d *Dot) Pos() [2]float64 {
	return d.pos
}

func (d *Dot) Size() float64 {
	return d.size
}

func (d *Dot) Speed() float64 {
	return d.speed
}
