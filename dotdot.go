package dotdot

import (
	"image/color"
	"math"
	"math/rand"
)

const (
	defaultDistance = 100
	defaultNumber   = 100
	defaultSpeed    = 150
	defaultSize     = 6
)

type DotDot struct {
	Distance float64
	Number   int
	MaxSpeed float64
	MinSpeed float64
	Speed    float64
	Size     float64

	width  float64
	height float64

	dots []*Dot
}

func New(width, height float64) (d *DotDot) {
	d = &DotDot{
		Distance: defaultDistance,
		Number:   defaultNumber,
		Speed:    defaultSpeed,
		Size:     defaultSize,
		width:    width,
		height:   height,
	}
	return
}

func (d *DotDot) InitDots() {
	d.normalizeArgs()
	d.dots = make([]*Dot, d.Number)
	randomSpeed := d.MinSpeed < d.MaxSpeed && d.MinSpeed > 0
	for i := 0; i < d.Number; i++ {
		dot := &Dot{
			color: color.RGBA{
				R: uint8(rand.Intn(255)),
				G: uint8(rand.Intn(255)),
				B: uint8(rand.Intn(255)),
				A: 255,
			},
			dir: [2]float64{
				rand.NormFloat64(),
				rand.NormFloat64(),
			},
			pos: [2]float64{
				rand.Float64() * d.width,
				rand.Float64() * d.height,
			},
			size: d.Size,
		}
		if randomSpeed {
			dot.speed = rand.Float64()*(d.MaxSpeed-d.MinSpeed) + d.MinSpeed
		} else {
			dot.speed = d.Speed
		}
		d.dots[i] = dot
	}
}

func (d *DotDot) Start() {
	d.InitDots()
}

func (d *DotDot) GetDots() []*Dot {
	return d.dots
}

func (d *DotDot) Update(cursorPos [2]float64) {
	for _, dot := range d.dots {
		// bounce
		if dot.pos[0] < 0 {
			dot.pos[0] = 0
			dot.dir[0] *= -1
		}
		if dot.pos[0] > d.width {
			dot.pos[0] = d.width
			dot.dir[0] *= -1
		}
		if dot.pos[1] < 0 {
			dot.pos[1] = 0
			dot.dir[1] *= -1
		}
		if dot.pos[1] > d.height {
			dot.pos[1] = d.height
			dot.dir[1] *= -1
		}

		// cursor
		dx := cursorPos[0] - dot.pos[0]
		dy := cursorPos[1] - dot.pos[1]
		dis := math.Sqrt(dx*dx + dy*dy)
		dot.distance = dis
		if dis < d.Distance {
			if !dot.captured {
				dot.captured = true
			}
			if dis > d.Distance/2 {
				// move towards mouse
				dot.dir[0] = dx / dis
				dot.dir[1] = dy / dis
			} else {
				// do not move
				dot.dir[0] = 0
				dot.dir[1] = 0
			}
		} else if dot.captured {
			dot.captured = false
		}

		// move
		dot.pos[0] += dot.dir[0] * dot.speed * 0.01
		dot.pos[1] += dot.dir[1] * dot.speed * 0.01
	}
}

func (d *DotDot) normalizeArgs() {
	if d.Number < 1 {
		d.Number = defaultNumber
	}
	if d.Size < 1 {
		d.Size = defaultSize
	}
	if d.Speed < 1 {
		d.Speed = defaultSpeed
	}
}
