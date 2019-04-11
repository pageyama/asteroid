package asteroid

import (
	"math"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type ship struct {
	x        int16
	y        int16
	r        int16
	angle    float64
	rotation float64

	vx  float64
	vy  float64
	acc float64
}

func newShip(x, y, r int16) *ship {
	return &ship{x: x, y: y, r: r}
}

func (s *ship) draw(r *sdl.Renderer) error {

	cos := math.Cos(s.angle)
	sin := math.Sin(s.angle)
	fr := float64(s.r)

	vx := []int16{
		s.x + int16(fr*sin),
		s.x + int16(fr*cos-fr*sin),
		s.x + int16(-fr*cos-fr*sin),
	}

	vy := []int16{
		s.y + int16(-fr*cos),
		s.y + int16(fr*sin+fr*cos),
		s.y + int16(-fr*sin+fr*cos),
	}

	// c := sdl.Color{R: 255, G: 255, B: 255, A: 255}

	// gfx.FilledPolygonColor(r, vx, vy, c)
	gfx.PolygonRGBA(r, vx, vy, 255, 255, 255, 255)
	return nil
}

func (s *ship) update() error {
	//rotation
	s.angle += s.rotation
	//accelaration
	s.vx += s.acc * math.Cos(s.angle-math.Pi/2.0)
	s.vy += s.acc * math.Sin(s.angle-math.Pi/2.0)
	//deceleration
	s.vx *= 0.97
	s.vy *= 0.97
	//moving
	s.x += int16(s.vx)
	s.y += int16(s.vy)

	return nil
}

func (s *ship) edge(w, h int16) error {
	if w+s.r < s.x {
		s.x = -s.r
	} else if -s.r > s.x {
		s.x = w + s.r
	}

	if h+s.r < s.y {
		s.y = -s.r
	} else if -s.r > s.y {
		s.y = h + s.r
	}

	return nil
}

func (s *ship) collision(a *asteroid) bool {
	d := distance(s.x, s.y, a.x, a.y)
	if d < s.r+a.r {
		return true
	}
	return false
}
