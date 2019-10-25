package game

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Score struct {
	x, y int32

	color sdl.Color
	font  *ttf.Font

	Value int
}

func NewScore(f *ttf.Font, c sdl.Color) *Score {
	return &Score{
		color: c,
		font:  f,
		Value: 0,
	}
}

func (s *Score) Update(p *Player, t *Trunk) {
	if p.X > (t.topPart.X+t.topPart.W) && !t.Scored {
		s.Value++
		t.Scored = true
	}
}

func (s *Score) Draw(r *sdl.Renderer) {
	t := fmt.Sprintf("SCORE: %d", s.Value)

	sur, _ := s.font.RenderUTF8Blended(t, s.color)
	defer sur.Free()

	tx, _ := r.CreateTextureFromSurface(sur)
	defer tx.Destroy()

	w, h, _ := s.font.SizeUTF8(t)

	src := &sdl.Rect{X: 0, Y: 0, W: int32(w), H: int32(h)}
	dst := &sdl.Rect{X: 40, Y: 0, W: int32(w), H: int32(h)}

	r.Copy(tx, src, dst)
}

func (s *Score) Destroy() error {
	s.font.Close()
	return nil
}
