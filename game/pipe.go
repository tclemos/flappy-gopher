package game

import (
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type Pipe struct {
	top, left int32
	topPart   sdl.Rect
	botPart   sdl.Rect
	gap, w    int32

	Active bool
}

type PipePool []*Pipe

func NewPipe(sh, sw int32) *Pipe {

	gap := sh / 10 * 4
	w := sh / 4 / 100 * 90

	return &Pipe{
		top:    getTopPosition(sh, gap),
		left:   sw,
		Active: false,
		w:      w,
		gap:    gap,
	}
}

func getTopPosition(sh, gap int32) int32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int31n(sh-gap-60) + 30
}

func (p *Pipe) Update(sw, sh, velocity int32) {
	if !p.Active {
		return
	}

	p.left -= velocity
	p.topPart = sdl.Rect{X: p.left, Y: 0, H: p.top, W: p.w}
	p.botPart = sdl.Rect{X: p.left, Y: p.top + p.gap, H: sh - (p.top + p.gap), W: p.w}
}

func (p *Pipe) Draw(r *sdl.Renderer) {
	if !p.Active {
		return
	}

	r.SetDrawColor(139, 69, 19, 255)
	r.FillRect(&p.topPart)
	r.FillRect(&p.botPart)
}

func (p *Pipe) OffScreen() bool {
	return p.left+p.w < 0
}

func (p *Pipe) Reset(sh, sw int32) {
	p.Active = false
	p.left = sw
	p.top = getTopPosition(sh, p.gap)
}

func (pp PipePool) Next() (*Pipe, bool) {
	for _, p := range pp {
		if !p.Active {
			return p, true
		}
	}

	return nil, false
}
