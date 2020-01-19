package game

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Game represents a game
type Game struct {
	w, h, v   int32
	window    *sdl.Window
	renderer  *sdl.Renderer
	player    *Player
	trunkPool TrunkPool
	cloudPool CloudPool
	grassPool GrassPool
	score     *Score
	running   bool
	over      bool
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
		1200, 720,
		sdl.WINDOW_OPENGL)
	if err != nil {
		return fmt.Errorf("Initializing window: %s", err.Error())
	}
	defer g.window.Destroy()
	// g.window.SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP)
	g.w, g.h = g.window.GetSize()

	g.renderer, err = sdl.CreateRenderer(g.window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		return fmt.Errorf("initializing renderer: %s", err.Error())
	}

	if img.Init(img.INIT_JPG|img.INIT_PNG) == 0 {
		return fmt.Errorf("Initializing image library")
	}

	if ttf.Init() != nil {
		return fmt.Errorf("Initializing font library")
	}

	// initialize random number generator
	rand.Seed(time.Now().UnixNano())

	if err = g.restart(); err != nil {
		return err
	}

	go g.handleVelocity()
	go g.handleTrunks()
	go g.handleClouds()
	go g.handleGrass()

	g.running = true
	for g.running {

		g.checkKeys()
		g.checkPollEvents()
		g.drawBackGround()
		g.drawClouds()
		g.drawGrasses()
		g.drawTrunks()
		g.drawPlayer()
		g.drawScore()
		g.checkCollision()

		g.renderer.Present()
		time.Sleep(1000 / 60 * time.Millisecond)
	}
	return nil
}

func (g *Game) createPlayer() error {
	if g.player != nil {
		g.player.Destroy()
	}

	image, _ := img.Load("./game/sprites/player.png")
	defer image.Free()
	texture, _ := g.renderer.CreateTextureFromSurface(image)
	g.player = NewPlayer(g.w, g.h, texture)
	return nil
}

func (g *Game) createPipes() error {
	if g.trunkPool != nil {
		g.trunkPool.Destroy()
	}

	image, _ := img.Load("./game/sprites/trunk.png")
	defer image.Free()
	texture, _ := g.renderer.CreateTextureFromSurface(image)

	g.trunkPool = TrunkPool{}
	for x := 0; x < 4; x++ {
		g.trunkPool = append(g.trunkPool, NewTrunk(g.h, g.w, texture))
	}

	return nil
}

func (g *Game) createClouds() error {
	if g.cloudPool != nil {
		g.cloudPool.Destroy()
	}

	image, _ := img.Load("./game/sprites/cloud.png")
	defer image.Free()
	tx, _ := g.renderer.CreateTextureFromSurface(image)

	g.cloudPool = CloudPool{}
	for x := 0; x < 7; x++ {
		g.cloudPool = append(g.cloudPool, NewCloud(g.h, g.w, tx))
	}

	return nil
}

func (g *Game) createGrasses() error {
	if g.grassPool != nil {
		g.grassPool.Destroy()
	}

	image, _ := img.Load("./game/sprites/grass.png")
	defer image.Free()
	tx, _ := g.renderer.CreateTextureFromSurface(image)

	g.grassPool = GrassPool{}
	for x := 0; x < 16; x++ {
		g.grassPool = append(g.grassPool, NewGrass(g.h, g.w, tx))
	}

	return nil
}

func (g *Game) createScore() error {
	if g.score != nil {
		g.score.Destroy()
	}

	font, _ := ttf.OpenFont("./game/fonts/Corben/Corben-Bold.ttf", 60)
	color := sdl.Color{R: 235, G: 213, B: 52, A: 255}

	g.score = NewScore(font, color)

	return nil
}

func (g *Game) handleVelocity() {
	for {
		g.v++
		time.Sleep(5 * time.Second)
	}
}

func (g *Game) handleTrunks() {
	for {
		if p, ok := g.trunkPool.Next(); ok {
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
		c.Update(g.over)
		c.Draw(g.renderer)
		if c.OffScreen() {
			c.Reset(g.h, g.w)
		}
	}
}

func (g *Game) drawGrasses() {
	for _, gr := range g.grassPool {
		if !g.over {
			gr.Update(g.v)
		}
		gr.Draw(g.renderer)
		if gr.OffScreen() {
			gr.Reset(g.h, g.w)
		}
	}
}

func (g *Game) drawTrunks() {
	for _, p := range g.trunkPool {
		if !g.over {
			p.Update(g.w, g.h, g.v)
		}
		p.Draw(g.renderer)
		if p.OffScreen() {
			p.Reset(g.h, g.w)
		}
	}
}

func (g *Game) drawPlayer() {
	if !g.over {
		g.player.Update(g.w, g.h)
	}
	g.player.Draw(g.renderer)
}

func (g *Game) drawScore() {
	t, found := g.trunkPool.NextToPlayer(g.player)
	if !g.over && found {
		g.score.Update(g.player, t)
	}
	g.score.Draw(g.renderer)
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

	if err := g.createScore(); err != nil {
		return err
	}

	g.v = int32(5)
	g.over = false

	runtime.GC()

	return nil
}

func (g *Game) quit() {
	g.running = false
}

func (g *Game) checkCollision() {
	for _, t := range g.trunkPool {
		if !t.Active {
			continue
		}
		if t.ColidesWith(g.player) {
			g.over = true
		}
	}
}
