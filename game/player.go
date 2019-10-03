package game

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	x, y, w, h int32
	velocity   float32
	gravity    float32
	lift       float32
	canJump    bool
	tex        *sdl.Texture
}

func NewPlayer(sw, sh int32, t *sdl.Texture) *Player {

	h := sh / 10 * 2
	w := h * 100 / 100 * 105 / 100

	p := &Player{
		w:        w,
		h:        h,
		gravity:  0.6,
		velocity: 0,
		lift:     -12,
		canJump:  true,
		tex:      t,
	}

	p.x = sw / 5
	p.y = sw / 5 * 2

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
	p.y += int32(p.velocity)

	if p.y+p.h >= sh {
		p.y = sh - p.h
		p.velocity = 0
	} else if p.y < 0 {
		p.y = 0
		p.velocity = 0
	}
}

func (p *Player) Draw(r *sdl.Renderer) {
	src := sdl.Rect{0, 0, 647, 572}
	dst := sdl.Rect{X: p.x, Y: p.y, W: p.w, H: p.h}
	r.SetDrawColor(0, 255, 255, 255)
	r.Copy(p.tex, &src, &dst)
}
