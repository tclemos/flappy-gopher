package game

import (
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type Grass struct {
	x, y, w, h int32
	tex        *sdl.Texture
	Active     bool
}

type GrassPool []*Grass

func NewGrass(sh, sw int32, t *sdl.Texture) *Grass {
	g := &Grass{
		tex: t,
	}
	g.Reset(sh, sw)
	return g
}

func (g *Grass) Reset(sh, sw int32) {
	rand.Seed(time.Now().UnixNano())
	g.x = sw
	g.w = sw * 1000 / 100 * (rand.Int31n(5) + 2) / 1000
	g.h = g.w
	g.y = rand.Int31n((sh / 3)) + (sh / 3 * 2) - g.h
	g.Active = false
}

func (g *Grass) Update(velocity int32) {
	if !g.Active {
		return
	}
	g.x -= velocity
}

func (g *Grass) Draw(r *sdl.Renderer) {
	if !g.Active {
		return
	}
	src := &sdl.Rect{0, 0, 500, 495}
	dst := &sdl.Rect{X: g.x, Y: g.y, W: g.w, H: g.h}
	r.SetDrawColor(255, 255, 255, 255)
	r.Copy(g.tex, src, dst)
}

func (g *Grass) Destroy() error {
	return g.tex.Destroy()
}

func (g *Grass) OffScreen() bool {
	return g.x+g.w < 0
}

func (gp GrassPool) Next() (*Grass, bool) {
	for _, g := range gp {
		if !g.Active {
			return g, true
		}
	}

	return nil, false
}

func (gg GrassPool) Destroy() error {
	for _, g := range gg {
		if err := g.Destroy(); err != nil {
			return err
		}
	}
	return nil
}
