package game

import (
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type Trunk struct {
	top, left int32
	topPart   sdl.Rect
	botPart   sdl.Rect
	gap, w    int32
	tex       *sdl.Texture

	Active bool
}

type TrunkPool []*Trunk

func NewTrunk(sh, sw int32, tex *sdl.Texture) *Trunk {

	gap := sh / 10 * 4
	w := sh / 4 / 100 * 90

	return &Trunk{
		top:    getTopPosition(sh, gap),
		left:   sw,
		Active: false,
		w:      w,
		gap:    gap,
		tex:    tex,
	}
}

func getTopPosition(sh, gap int32) int32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int31n(sh-gap-60) + 30
}

func (t *Trunk) Update(sw, sh, velocity int32) {
	if !t.Active {
		return
	}

	t.left -= velocity
	t.topPart = sdl.Rect{X: t.left, Y: 0, H: t.top, W: t.w}
	t.botPart = sdl.Rect{X: t.left, Y: t.top + t.gap, H: sh - (t.top + t.gap), W: t.w}
}

func (t *Trunk) Draw(r *sdl.Renderer) {
	if !t.Active {
		return
	}
	src := &sdl.Rect{0, 0, 1048, 1920}
	r.Copy(t.tex, src, &t.botPart)
	r.Copy(t.tex, src, &t.topPart)
}

func (t *Trunk) OffScreen() bool {
	return t.left+t.w < 0
}

func (t *Trunk) Reset(sh, sw int32) {
	t.Active = false
	t.left = sw
	t.top = getTopPosition(sh, t.gap)
}

func (tt TrunkPool) Next() (*Trunk, bool) {
	for _, t := range tt {
		if !t.Active {
			return t, true
		}
	}

	return nil, false
}
