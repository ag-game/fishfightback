package system

import (
	"code.rocketnine.space/tslocum/fishfightback/component"
	"code.rocketnine.space/tslocum/fishfightback/world"
	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
)

type RailSystem struct {
	Rail     *component.Rail
	Position *component.Position
}

func NewRailSystem() *RailSystem {
	s := &RailSystem{}

	return s
}

func (s *RailSystem) Update(e gohan.Entity) error {
	if !world.World.GameStarted || world.World.GameOver || !world.World.CamMoving {
		return nil
	}

	s.Position.X += CameraMoveSpeed
	return nil
}

func (_ *RailSystem) Draw(_ gohan.Entity, _ *ebiten.Image) error {
	return gohan.ErrUnregister
}
