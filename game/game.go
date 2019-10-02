package game

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

// Game represents a game
type Game struct {
	window   *sdl.Window
	surface  *sdl.Surface
	player   *Player
	pipePool PipePool
}

// New creates a new instance of Game
func New() Game {
	return Game{}
}

// Init initializes the game
func (g *Game) Init() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("Initializing SDL: %s", err.Error())
	}
	defer sdl.Quit()

	g.window, err = sdl.CreateWindow("Flappy Gopher",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		0, 0,
		sdl.WINDOW_OPENGL)
	if err != nil {
		return fmt.Errorf("Initializing window: %s", err.Error())
	}
	defer g.window.Destroy()
	g.window.SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP)
	sw, sh := g.window.GetSize()

	g.surface, err = g.window.GetSurface()
	if err != nil {
		return fmt.Errorf("Getting surface: %s", err.Error())
	}

	g.player = NewPlayer(sw/4, sh/2)
	g.pipePool = PipePool([]*Pipe{
		NewPipe(sh, sw),
		NewPipe(sh, sw),
		NewPipe(sh, sw),
		NewPipe(sh, sw),
	})

	gameVelocity := int32(1)
	go func(pp *PipePool) {
		for {
			if p, ok := pp.Next(); ok {
				p.Active = true
			}
			gameVelocity += 1
			time.Sleep(1 * time.Second)
		}
	}(&g.pipePool)

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
		// Background
		g.surface.FillRect(&sdl.Rect{X: 0, Y: 0, W: sw, H: sh / 3 * 2}, 0xff0000FF)
		g.surface.FillRect(&sdl.Rect{X: 0, Y: sh / 3 * 2, W: sw, H: sh / 3}, 0xff00FF00)

		g.player.Update(sw, sh)
		g.player.Draw(g.surface)

		for _, p := range g.pipePool {
			p.Update(sw, sh, gameVelocity)
			p.Draw(g.surface)
			if p.OffScreen() {
				p.Reset(sh, sw)
			}
		}

		g.window.UpdateSurface()
		time.Sleep(1000 / 60 * time.Millisecond)
	}
	return nil
}
