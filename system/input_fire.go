package system

import (
	"math"

	"code.rocketnine.space/tslocum/fishfightback/asset"

	"code.rocketnine.space/tslocum/fishfightback/component"
	"code.rocketnine.space/tslocum/fishfightback/entity"
	"code.rocketnine.space/tslocum/fishfightback/world"
	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
)

type playerFireSystem struct {
	Position *component.Position
	Velocity *component.Velocity
	Weapon   *component.Weapon
	Sprite   *component.Sprite
}

func NewplayerFireSystem() *playerFireSystem {
	return &playerFireSystem{}
}

func (s *playerFireSystem) Update(e gohan.Entity) error {
	if !world.World.GameStarted || world.World.GameOver {
		return nil
	}

	position := s.Position
	weapon := s.Weapon

	cx, cy := ebiten.CursorPosition()
	px, py := world.LevelCoordinatesToScreen(s.Position.X+8, s.Position.Y+8)
	pa := angle(float64(cx), float64(cy), px, py)

	if ebiten.IsKeyPressed(ebiten.KeyZ) || ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if weapon.NextFire == 0 {
			bx := math.Cos(pa) * weapon.BulletSpeed
			by := math.Sin(pa) * weapon.BulletSpeed

			entity.NewPlayerBullet(position.X+6, position.Y+6, bx, by)
			weapon.NextFire = weapon.FireRate

			asset.SoundFire.Rewind()
			asset.SoundFire.Play()
		}
	}
	if weapon.NextFire > 0 {
		weapon.NextFire--
	}
	return nil
}

func (s *playerFireSystem) Draw(_ gohan.Entity, _ *ebiten.Image) error {
	return gohan.ErrUnregister
}
