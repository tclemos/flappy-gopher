package game

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	x, y, w, h int
	velocity   float32
	gravity    float32
	lift       float32
	jumpCD     bool
}

func NewPlayer(x, y int) *Player {
	p := &Player{
		w:        140,
		h:        180,
		gravity:  0.7,
		velocity: 0,
		lift:     -13,
		jumpCD:   false,
	}

	p.x = x - (p.w / 2)
	p.y = y - (p.h / 2)

	return p
}

func (p *Player) Update(sw, sh int) {
	// fmt.Println("Player Update")
	// fmt.Println(p.y, sh)

	keys := sdl.GetKeyboardState()

	fmt.Println(p.velocity)

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if !p.jumpCD {
			p.jumpCD = true
			p.velocity = p.lift
		}
	} else {
		p.jumpCD = false
	}

	p.velocity += p.gravity
	p.y += int(p.velocity)

	if p.y+p.h >= sh {
		p.y = sh - p.h
		p.velocity = 0
	} else if p.y < 0 {
		p.y = 0
		p.velocity = 0
	}
}

func (p *Player) Draw(s *sdl.Surface) {
	rect := sdl.Rect{
		X: int32(p.x),
		Y: int32(p.y),
		W: int32(p.w),
		H: int32(p.h),
	}
	s.FillRect(&rect, 0xff00FFFF)
}
