package asteroid

import (
	"math"
	"math/rand"

	"github.com/veandco/go-sdl2/gfx"

	"github.com/veandco/go-sdl2/sdl"
)

type asteroid struct {
	x, y, r  int16
	a        float64
	rank     int
	offset   []float64
	isBroken bool
}

func newAsteroid(x, y, r int16) *asteroid {

	rank := rand.Intn(6) + 6
	offset := make([]float64, rank)
	for i := range offset {
		offset[i] = rand.Float64()*0.6 + 0.5
	}

	a := asteroid{
		x:        x,
		y:        y,
		r:        r,
		a:        rand.Float64() * math.Pi * 2,
		rank:     rank,
		offset:   offset,
		isBroken: false,
	}
	return &a
}

func (a *asteroid) draw(r *sdl.Renderer) error {

	radius := float64(a.r)
	vx := make([]int16, a.rank)
	vy := make([]int16, a.rank)
	for i := 0; i < a.rank; i++ {
		t := float64(i) * 2.0 * math.Pi / float64(a.rank)
		vx[i] = int16(radius*a.offset[i]*math.Cos(t)) + a.x
		vy[i] = int16(radius*a.offset[i]*math.Sin(t)) + a.y
	}

	gfx.PolygonRGBA(r, vx, vy, 255, 255, 255, 255)
	return nil
}

func (a *asteroid) update() error {
	a.x += int16(math.Cos(a.a) * 3)
	a.y += int16(math.Sin(a.a) * 3)
	return nil
}

func (a *asteroid) edge(w, h int16) error {
	if w+a.r < a.x {
		a.x = -a.r
	} else if -a.r > a.x {
		a.x = w + a.r
	}

	if h+a.r < a.y {
		a.y = -a.r
	} else if -a.r > a.y {
		a.y = h + a.r
	}

	return nil
}

func (a *asteroid) breakup() []*asteroid {
	r := a.r / 2
	if r < 12 {
		return nil
	}
	s := make([]*asteroid, 2)
	s[0] = newAsteroid(a.x, a.y, r)
	s[1] = newAsteroid(a.x, a.y, r)
	return s
}
