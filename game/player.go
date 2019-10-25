package game

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	X, Y, W, H int32
	hitboxes   []*sdl.Rect
	velocity   float32
	gravity    float32
	lift       float32
	canJump    bool
	tex        *sdl.Texture
}

func NewPlayer(sw, sh int32, t *sdl.Texture) *Player {

	h := int32(float32(sh) * 0.2)
	w := int32(float32(h) * 1.05)

	p := &Player{
		W:        w,
		H:        h,
		gravity:  0.6,
		velocity: 0,
		lift:     -12,
		canJump:  true,
		tex:      t,
	}

	p.X = sw / 5
	p.Y = sh / 3 * 1

	return p
}

func (p *Player) Update(sw, sh int32) {

	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if p.canJump {
			p.canJump = false
			p.velocity = p.lift
		}
	} else {
		p.canJump = true
	}

	p.velocity += p.gravity
	p.Y += int32(p.velocity)

	if p.Y+p.H >= sh {
		p.Y = sh - p.H
		p.velocity = 0
	} else if p.Y < 0 {
		p.Y = 0
		p.velocity = 0
	}

	hb1 := &sdl.Rect{
		X: p.X + int32((float32(p.W) * 0.45)),
		Y: p.Y + int32((float32(p.H) * 0.07)),
		W: p.W - int32((float32(p.W) * 0.65)),
		H: p.H - int32((float32(p.H) * 0.15)),
	}

	hb2 := &sdl.Rect{
		X: p.X + int32((float32(p.W) * 0.35)),
		Y: p.Y + int32((float32(p.H) * 0.15)),
		W: p.W - int32((float32(p.W) * 0.47)),
		H: p.H - int32((float32(p.H) * 0.30)),
	}

	p.hitboxes = []*sdl.Rect{
		hb1, hb2,
	}
}

func (p *Player) Draw(r *sdl.Renderer) {
	src := sdl.Rect{X: 0, Y: 0, W: 647, H: 572}
	dst := sdl.Rect{X: p.X, Y: p.Y, W: p.W, H: p.H}
	r.Copy(p.tex, &src, &dst)
	// r.SetDrawColor(255, 0, 0, 255)
	// r.FillRect(p.hitboxes[0])
	// r.SetDrawColor(0, 255, 0, 255)
	// r.FillRect(p.hitboxes[1])
}

func (p *Player) Destroy() error {
	err := p.tex.Destroy()
	return err
}
