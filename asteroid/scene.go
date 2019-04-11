package asteroid

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type scene struct {
	w, h       int16
	ship       *ship
	asteroids  []*asteroid
	leasers    []*leaser
	isGameOver bool
	score      int
}

func newScene(width, height int32) *scene {

	w := int16(width)
	h := int16(height)

	ship := newShip(int16(w/2), int16(h/2), 24)

	n := 12
	asteroids := make([]*asteroid, n)
	for i := 0; i < n; i++ {

		x := randInt(w)
		y := randInt(h)
		r := randInt(24) + 36
		d := distance(ship.x, ship.y, x, y)

		if d < 2*ship.r+r {
			i--
			continue
		}

		asteroids[i] = newAsteroid(x, y, r)
	}

	leasers := make([]*leaser, 0)

	s := scene{
		w:          w,
		h:          h,
		ship:       ship,
		asteroids:  asteroids,
		leasers:    leasers,
		isGameOver: false,
		score:      0,
	}

	return &s
}

func (s *scene) update() error {
	if err := s.ship.update(); err != nil {
		return err
	}

	if err := s.ship.edge(int16(s.w), int16(s.h)); err != nil {
		return err
	}

	for _, asteroid := range s.asteroids {
		if err := asteroid.update(); err != nil {
			return err
		}

		if err := asteroid.edge(s.w, s.h); err != nil {
			return err
		}

		if s.ship.collision(asteroid) {
			s.isGameOver = true
			return nil
		}
	}

	isHitten := false
	count := 0
	for _, leaser := range s.leasers {
		if err := leaser.update(); err != nil {
			return nil
		}
		if leaser.offscrean(s.w, s.h) {
			leaser.isActive = false
			continue
		}

		for _, asteroid := range s.asteroids {
			if leaser.collision(asteroid) {
				leaser.isActive = false
				asteroid.isBroken = true
				isHitten = true
				break
			}
		}

		if leaser.isActive {
			count++
		}
	}

	if count != len(s.leasers) {
		leasers := make([]*leaser, count)
		index := 0
		for _, leaser := range s.leasers {
			if leaser.isActive {
				leasers[index] = leaser
				index++
			}
		}
		s.leasers = leasers
	}

	if isHitten {
		asteroids := make([]*asteroid, 0)
		for _, asteroid := range s.asteroids {
			if !asteroid.isBroken {
				asteroids = append(asteroids, asteroid)
				continue
			}

			news := asteroid.breakup()
			if news != nil {
				asteroids = append(asteroids, news...)
				continue
			}

			s.score++
		}
		s.asteroids = asteroids
	}

	for len(s.asteroids) < 12 {
		r := randInt(24) + 36
		asteroid := newAsteroid(-r, -r, r)
		s.asteroids = append(s.asteroids, asteroid)
	}

	return nil
}

func (s *scene) draw(r *sdl.Renderer) error {
	if err := s.ship.draw(r); err != nil {
		return err
	}

	for _, asteroid := range s.asteroids {
		if err := asteroid.draw(r); err != nil {
			return err
		}
	}

	for _, leaser := range s.leasers {
		if err := leaser.draw(r); err != nil {
			return nil
		}
	}

	gfx.StringRGBA(r, 0, 0, fmt.Sprintf("Score : %v", s.score), 255, 255, 255, 255)
	gfx.StringRGBA(r, 0, 10, fmt.Sprintf("Leasers Length : %v", len(s.leasers)), 255, 255, 255, 255)
	gfx.StringRGBA(r, 0, 20, fmt.Sprintf("Astroids Length : %v", len(s.asteroids)), 255, 255, 255, 255)

	return nil
}

func (s *scene) onKeyboardEvent(e *sdl.KeyboardEvent) error {

	if e.State == sdl.PRESSED {
		s.onKeyPressed(e.Keysym.Sym)
	} else if e.State == sdl.RELEASED {
		s.onKeyReleased(e.Keysym.Sym)
	}

	return nil
}

func (s *scene) onKeyPressed(k sdl.Keycode) error {
	if k == sdl.K_RIGHT {
		s.ship.rotation = 0.1
	} else if k == sdl.K_LEFT {
		s.ship.rotation = -0.1
	} else if k == sdl.K_UP {
		s.ship.acc = 0.8
	} else if k == sdl.K_SPACE {
		l := newLeaser(s.ship.x, s.ship.y, s.ship.angle)
		s.leasers = append(s.leasers, l)
	}

	return nil
}

func (s *scene) onKeyReleased(k sdl.Keycode) error {
	if k == sdl.K_RIGHT || k == sdl.K_LEFT {
		s.ship.rotation = 0.0
	} else if k == sdl.K_UP {
		s.ship.acc = 0.0
	}

	return nil
}

func randInt(i int16) int16 {
	return int16(rand.Intn(int(i)))
}

func distance(x1, y1, x2, y2 int16) int16 {
	d := 0.0
	d += math.Pow(float64(x1)-float64(x2), 2.0)
	d += math.Pow(float64(y1)-float64(y2), 2.0)
	return int16(math.Sqrt(d))
}
