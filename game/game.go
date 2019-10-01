package game

import (
	"errors"
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type Component interface {
	Update(sw, sh int)
	Draw(*sdl.Surface)
}

type Game struct {
	window     *sdl.Window
	surface    *sdl.Surface
	components []Component
}

func New() Game {
	return Game{}
}

func (g *Game) Init() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return errors.New(fmt.Sprintf("Initializing SDL: %s", err.Error()))
	}
	defer sdl.Quit()

	g.window, err = sdl.CreateWindow("Flappy Gopher",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		0, 0,
		sdl.WINDOW_OPENGL)
	if err != nil {
		return errors.New(fmt.Sprintf("Initializing window: %s", err.Error()))
	}
	defer g.window.Destroy()
	g.window.SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP)
	sw, sh := g.window.GetSize()

	g.surface, err = g.window.GetSurface()
	if err != nil {
		return errors.New(fmt.Sprintf("Getting surface: %s", err.Error()))
	}

	g.components = make([]Component, 0)
	g.components = append(g.components, NewPlayer(int(sw)/4, int(sh)/2))

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
		g.surface.FillRect(&sdl.Rect{X: 0, Y: 0, W: sw, H: sh / 3 * 2}, 0xff0000FF)
		g.surface.FillRect(&sdl.Rect{X: 0, Y: sh / 3 * 2, W: sw, H: sh / 3}, 0xff00FF00)

		for _, c := range g.components {
			c.Update(int(sw), int(sh))
			c.Draw(g.surface)
		}
		g.window.UpdateSurface()
		time.Sleep(1000 / 60 * time.Millisecond)
	}
	return nil
}
