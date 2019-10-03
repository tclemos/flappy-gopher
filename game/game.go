package game

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// Game represents a game
type Game struct {
	w, h, v   int32
	window    *sdl.Window
	renderer  *sdl.Renderer
	player    *Player
	pipePool  PipePool
	cloudPool CloudPool
	grassPool GrassPool
	running   bool
}

// New creates a new instance of Game
func New() Game {
	return Game{}
}

// Init initializes the game
func (g *Game) Init() error {

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("Initializing SDL: %s\n", err.Error())
	}
	defer sdl.Quit()

	g.window, err = sdl.CreateWindow("Flappy Gopher",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		1200, 720,
		sdl.WINDOW_OPENGL)
	if err != nil {
		return fmt.Errorf("Initializing window: %s\n", err.Error())
	}
	defer g.window.Destroy()
	//g.window.SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP)
	g.w, g.h = g.window.GetSize()

	g.renderer, err = sdl.CreateRenderer(g.window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		return fmt.Errorf("Initializing Renderer: %s\n", err.Error())
	}

	if err = g.restart(); err != nil {
		return err
	}

	go g.handleVelocity()
	go g.handlePipes()
	go g.handleClouds()
	go g.handleGrass()

	g.running = true
	for g.running {

		g.checkKeys()
		g.checkPollEvents()
		g.drawBackGround()
		g.drawClouds()
		g.drawGrasses()
		g.drawPipes()
		g.drawPlayer()

		g.renderer.Present()
		time.Sleep(1000 / 60 * time.Millisecond)
	}
	return nil
}

func (g *Game) createPlayer() error {
	image, err := img.Load("../game/sprites/player.png")
	if err != nil {
		return fmt.Errorf("Failed to load PNG: %s\n", err)
	}

	texture, err := g.renderer.CreateTextureFromSurface(image)
	if err != nil {
		return fmt.Errorf("Failed to create texture: %s\n", err)
	}

	g.player = NewPlayer(g.w, g.h, texture)
	return nil
}

func (g *Game) createPipes() error {
	g.pipePool = PipePool([]*Pipe{
		NewPipe(g.h, g.w),
		NewPipe(g.h, g.w),
		NewPipe(g.h, g.w),
		NewPipe(g.h, g.w),
	})
	return nil
}

func (g *Game) createClouds() error {
	image, _ := img.Load("../game/sprites/cloud.png")
	tx, _ := g.renderer.CreateTextureFromSurface(image)

	g.cloudPool = CloudPool([]*Cloud{
		NewCloud(g.h, g.w, tx),
		NewCloud(g.h, g.w, tx),
		NewCloud(g.h, g.w, tx),
		NewCloud(g.h, g.w, tx),
		NewCloud(g.h, g.w, tx),
		NewCloud(g.h, g.w, tx),
		NewCloud(g.h, g.w, tx),
		NewCloud(g.h, g.w, tx),
		NewCloud(g.h, g.w, tx),
		NewCloud(g.h, g.w, tx),
	})
	return nil
}

func (g *Game) createGrasses() error {
	image, _ := img.Load("../game/sprites/grass.png")
	tx, _ := g.renderer.CreateTextureFromSurface(image)

	g.grassPool = GrassPool([]*Grass{
		NewGrass(g.h, g.w, tx),
		NewGrass(g.h, g.w, tx),
		NewGrass(g.h, g.w, tx),
		NewGrass(g.h, g.w, tx),
		NewGrass(g.h, g.w, tx),
		NewGrass(g.h, g.w, tx),
		NewGrass(g.h, g.w, tx),
		NewGrass(g.h, g.w, tx),
		NewGrass(g.h, g.w, tx),
		NewGrass(g.h, g.w, tx),
		NewGrass(g.h, g.w, tx),
		NewGrass(g.h, g.w, tx),
		NewGrass(g.h, g.w, tx),
		NewGrass(g.h, g.w, tx),
		NewGrass(g.h, g.w, tx),
		NewGrass(g.h, g.w, tx),
	})
	return nil
}

func (g *Game) handleVelocity() {
	for {
		g.v++
		time.Sleep(6 * time.Second)
	}
}

func (g *Game) handlePipes() {
	for {
		if p, ok := g.pipePool.Next(); ok {
			p.Active = true
		}
		time.Sleep(3 * time.Second)
	}
}

func (g *Game) handleClouds() {
	for {
		if c, ok := g.cloudPool.Next(); ok {
			c.Active = true
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func (g *Game) handleGrass() {
	for {
		if g, ok := g.grassPool.Next(); ok {
			g.Active = true
		}
		time.Sleep(time.Duration((500 - (10 * g.v))) * time.Millisecond)
	}
}

func (g *Game) checkKeys() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_RETURN] == 1 {
		g.restart()
	} else if keys[sdl.SCANCODE_ESCAPE] == 1 {
		g.quit()
	}
}

func (g *Game) checkPollEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			g.quit()
			break
		}
	}
}

func (g *Game) drawBackGround() {
	sky := &sdl.Rect{X: 0, Y: 0, W: g.w, H: g.h / 3 * 2}
	floor := &sdl.Rect{X: 0, Y: g.h / 3 * 2, W: g.w, H: g.h / 3}
	g.renderer.SetDrawColor(0, 122, 255, 255)
	g.renderer.FillRect(sky)
	g.renderer.SetDrawColor(100, 185, 8, 255)
	g.renderer.FillRect(floor)
}

func (g *Game) drawClouds() {
	for _, c := range g.cloudPool {
		c.Update()
		c.Draw(g.renderer)
		if c.OffScreen() {
			c.Reset(g.h, g.w)
		}
	}
}

func (g *Game) drawGrasses() {
	for _, gr := range g.grassPool {
		gr.Update(g.v)
		gr.Draw(g.renderer)
		if gr.OffScreen() {
			gr.Reset(g.h, g.w)
		}
	}
}

func (g *Game) drawPipes() {
	for _, p := range g.pipePool {
		p.Update(g.w, g.h, g.v)
		p.Draw(g.renderer)
		if p.OffScreen() {
			p.Reset(g.h, g.w)
		}
	}
}

func (g *Game) drawPlayer() {
	g.player.Update(g.w, g.h)
	g.player.Draw(g.renderer)
}

func (g *Game) restart() error {
	if err := g.createPlayer(); err != nil {
		return err
	}

	if err := g.createPipes(); err != nil {
		return err
	}

	if err := g.createClouds(); err != nil {
		return err
	}

	if err := g.createGrasses(); err != nil {
		return err
	}

	g.v = int32(5)
	return nil
}

func (g *Game) quit() {
	g.running = false
}
