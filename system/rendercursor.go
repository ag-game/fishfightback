package system

import (
	"code.rocketnine.space/tslocum/fishfightback/asset"
	"code.rocketnine.space/tslocum/fishfightback/component"

	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
)

type RenderCursorSystem struct {
	Position *component.Position
	Velocity *component.Velocity
	Weapon   *component.Weapon

	cx, cy int
}

func NewRenderCursorSystem() *RenderCursorSystem {
	s := &RenderCursorSystem{}

	return s
}

func (s *RenderCursorSystem) Update(_ gohan.Entity) error {
	s.cx, s.cy = ebiten.CursorPosition()
	return nil
}

func (s *RenderCursorSystem) Draw(_ gohan.Entity, screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.cx-2), float64(s.cy-2))
	screen.DrawImage(asset.ImgCrosshair, op)
	return nil
}
