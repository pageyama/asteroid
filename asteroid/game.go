package asteroid

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/gfx"

	"github.com/veandco/go-sdl2/sdl"
)

//Run start asteroid game
func Run(winWidth, winHeight int32, fps uint32) error {

	rand.Seed(time.Now().UnixNano())

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return fmt.Errorf("could not initialize SDL : %v", err)
	}
	defer sdl.Quit()

	w, r, err := sdl.CreateWindowAndRenderer(winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("could not create window : %v", err)
	}
	defer w.Destroy()
	defer r.Destroy()

	fpsMgr := new(gfx.FPSmanager)
	gfx.SetFramerate(fpsMgr, fps)

	s := newScene(winWidth, winHeight)

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			case *sdl.KeyboardEvent:
				s.onKeyboardEvent(t)
			}
		}

		//refresh
		r.Clear()
		r.SetDrawColor(0, 0, 0, 255)
		r.FillRect(&sdl.Rect{X: 0, Y: 0, W: winWidth, H: winHeight})

		s.update()

		if s.isGameOver {
			s = newScene(winWidth, winHeight)
			continue
		}

		s.draw(r)

		r.Present()

		gfx.FramerateDelay(fpsMgr)
	}

	return nil
}
