package game

import (
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

type Cloud struct {
	x, y, w, h int32
	velocity   int32
	tex        *sdl.Texture
	Active     bool
}

type CloudPool []*Cloud

func NewCloud(sh, sw int32, t *sdl.Texture) *Cloud {
	c := &Cloud{
		tex: t,
	}
	c.Reset(sh, sw)
	return c
}

func (c *Cloud) Reset(sh, sw int32) {
	c.y = rand.Int31n(sh/3*2 - 100)
	c.x = sw
	c.velocity = rand.Int31n(10) + 1
	c.w = sw/10*rand.Int31n(3) + 1
	c.h = c.w * 100 / 10 * 3 / 100
	c.Active = false
}

func (c *Cloud) Update(slow bool) {
	if !c.Active {
		return
	}

	v := c.velocity
	if slow {
		v = c.velocity / 3
	}
	c.x -= v
}

func (c *Cloud) Draw(r *sdl.Renderer) {
	if !c.Active {
		return
	}
	src := &sdl.Rect{0, 0, 557, 277}
	dst := &sdl.Rect{X: c.x, Y: c.y, W: c.w, H: c.h}
	r.SetDrawColor(255, 255, 255, 255)
	r.Copy(c.tex, src, dst)
}

func (c *Cloud) OffScreen() bool {
	return c.x+c.w < 0
}

func (c *Cloud) Destroy() error {
	return c.tex.Destroy()
}

func (cp CloudPool) Next() (*Cloud, bool) {
	for _, c := range cp {
		if !c.Active {
			return c, true
		}
	}

	return nil, false
}

func (cc CloudPool) Destroy() error {
	for _, c := range cc {
		if err := c.Destroy(); err != nil {
			return err
		}
	}
	return nil
}
