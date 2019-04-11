package asteroid

import (
	"math"

	"github.com/veandco/go-sdl2/gfx"

	"github.com/veandco/go-sdl2/sdl"
)

type leaser struct {
	x        int16
	y        int16
	angle    float64
	speed    float64
	isActive bool
}

func newLeaser(x, y int16, angle float64) *leaser {
	l := leaser{
		x:        x,
		y:        y,
		angle:    angle,
		speed:    16,
		isActive: true,
	}
	return &l
}

func (l *leaser) update() error {
	l.x += int16(l.speed * math.Cos(l.angle-math.Pi/2.0))
	l.y += int16(l.speed * math.Sin(l.angle-math.Pi/2.0))
	return nil
}

func (l *leaser) draw(r *sdl.Renderer) error {
	x := int32(l.x)
	y := int32(l.y)
	gfx.FilledCircleRGBA(r, x, y, 2, 255, 255, 255, 255)
	return nil
}

func (l *leaser) offscrean(w, h int16) bool {
	if l.x < 0 || w < int16(l.x) || l.y < 0 || h < int16(l.y) {
		return true
	}

	return false
}

func (l *leaser) collision(a *asteroid) bool {
	d := distance(l.x, l.y, a.x, a.y)
	if d < a.r {
		return true
	}
	return false
}
