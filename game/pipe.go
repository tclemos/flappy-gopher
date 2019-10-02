package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	pipeGap   = 420
	pipeWidth = 140
)

type Pipe struct {
	top, left int32
	velocity  int32
	topPart   sdl.Rect
	botPart   sdl.Rect

	Active bool
}

type PipePool []*Pipe

func NewPipe(sh, sw int32) *Pipe {

	return &Pipe{
		top:      getTopPosition(sh),
		velocity: 3,
		left:     sw,
		Active:   false,
	}
}

func getTopPosition(sh int32) int32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int31n(sh-pipeGap-30) + 30
}

func (p *Pipe) Update(sw, sh, gameVelocity int32) {
	if !p.Active {
		return
	}

	fmt.Println("Pipe update:", p.top, p.velocity, p.left)
	p.left -= p.velocity + gameVelocity
	p.topPart = sdl.Rect{X: p.left, Y: 0, H: p.top, W: pipeWidth}
	p.botPart = sdl.Rect{X: p.left, Y: p.top + pipeGap, H: sh - (p.top + pipeGap), W: pipeWidth}
}

func (p *Pipe) Draw(s *sdl.Surface) {
	if !p.Active {
		return
	}

	fmt.Println("Pipe draw")
	fmt.Println("topPart:", p.topPart.Y, p.topPart.X, p.topPart.H, p.topPart.W)
	fmt.Println("botPart:", p.botPart.Y, p.botPart.X, p.botPart.H, p.botPart.W)

	s.FillRect(&p.topPart, 0xff8B4513)
	s.FillRect(&p.botPart, 0xff8B4513)
}

func (p *Pipe) OffScreen() bool {
	return p.left+pipeWidth < 0
}

func (p *Pipe) Reset(sh, sw int32) {
	p.Active = false
	p.left = sw
	p.top = getTopPosition(sh)
}

func (pp PipePool) Next() (*Pipe, bool) {
	for _, p := range pp {
		if !p.Active {
			return p, true
		}
	}

	return nil, false
}
