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
	Scored bool
}

type TrunkPool []*Trunk

func NewTrunk(sh, sw int32, tex *sdl.Texture) *Trunk {

	gap := int32(float32(sh) * 0.40)

	w := int32(float32(sh) * 0.2 * 1.05)

	return &Trunk{
		top:  getTopPosition(sh, gap),
		left: sw,
		w:    w,
		gap:  gap,
		tex:  tex,

		Active: false,
		Scored: false,
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
	r.Copy(t.tex, src, &t.topPart)
	r.Copy(t.tex, src, &t.botPart)
}

func (t *Trunk) OffScreen() bool {
	return t.left+t.w < 0
}

func (t *Trunk) Reset(sh, sw int32) {
	t.Active = false
	t.Scored = false

	t.left = sw
	t.top = getTopPosition(sh, t.gap)
}

func (t *Trunk) ColidesWith(p *Player) bool {
	for _, hb := range p.hitboxes {
		playerTopLeft := sdl.Point{X: hb.X, Y: hb.Y}
		playerBottomRight := sdl.Point{X: hb.X + hb.W, Y: hb.Y + hb.H}

		trunkTopPartTopLeft := sdl.Point{X: t.topPart.X, Y: t.topPart.Y}
		trunkTopPartBottomRight := sdl.Point{X: t.topPart.X + t.topPart.W, Y: t.topPart.Y + t.topPart.H}

		trunkBotPartTopLeft := sdl.Point{X: t.botPart.X, Y: t.botPart.Y}

		// As both trunk parts moves at the same speed and has the same left/right position
		// we use only one of the trunk parts to make sure the player is not coliding by the sides.
		//
		// The condition below represents: if the player is not crossing the pipes, it is not colliding
		if trunkTopPartTopLeft.X > playerBottomRight.X || playerTopLeft.X > trunkTopPartBottomRight.X {
			continue
		}

		// Since the player can't move out of the screen, we just need to check if it colides
		// from the top of the bottom part of the trunk and from the bottom of the top part of the trunk
		// since the last verification is checking side collisions
		//
		// The condition bellow represents: if the player is fully inside of the gap, it not is not colliding
		if playerTopLeft.Y > trunkTopPartBottomRight.Y && playerBottomRight.Y < trunkBotPartTopLeft.Y {
			continue
		}

		return true
	}

	return false
}

func (tt TrunkPool) Next() (*Trunk, bool) {
	for _, t := range tt {
		if !t.Active {
			return t, true
		}
	}

	return nil, false
}

func (tt TrunkPool) NextToPlayer(p *Player) (*Trunk, bool) {
	var r *Trunk

	for _, t := range tt {
		if !t.Active {
			continue
		}

		if r == nil || (r.topPart.X > t.topPart.X && (t.topPart.X+t.topPart.W) < p.X) {
			r = t
		}
	}

	return r, r != nil
}
