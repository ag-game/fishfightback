package system

import (
	"code.rocketnine.space/tslocum/fishfightback/component"
	"code.rocketnine.space/tslocum/fishfightback/entity"
	"code.rocketnine.space/tslocum/fishfightback/world"
	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
)

type CreepSystem struct {
	Creep    *component.Creep
	Position *component.Position

	Sprite *component.Sprite `gohan:"?"`
	Weapon *component.Weapon `gohan:"?"`
}

func NewCreepSystem() *CreepSystem {
	s := &CreepSystem{}

	return s
}

func (s *CreepSystem) Update(e gohan.Entity) error {
	if !world.World.GameStarted || world.World.GameOver {
		return nil
	}

	creep := s.Creep
	position := s.Position

	if creep.Health <= 0 {
		// TODO replace with dead sprite.
		e.Remove()
		return nil
	}

	// Skip inactive creeps.
	sx, sy := world.LevelCoordinatesToScreen(position.X, position.Y)
	inactive := sx < 0 || sy < 0 || sx > float64(world.ScreenWidth)+TileWidth/2 || sy > float64(world.ScreenHeight)
	if creep.Active != !inactive {
		creep.Active = !inactive
	}
	if inactive {
		return nil
	}

	l := len(creep.Movements)
	if l > creep.Movement {
		if creep.MovementTicks == 0 {
			m := creep.Movements[creep.Movement]
			position.X, position.Y = m[0], m[1]
			creep.Movement++

			creep.MovementTicks = int(m[2])
		}
		creep.MovementTicks--
	}

	randVelocity := func() (float64, float64) {
		for {
			const minSpeed = 0.1
			vx := creep.Rand.Float64()*0.5 + (0.5 - creep.Rand.Float64())
			vy := creep.Rand.Float64() * 0.5
			if (vx > minSpeed || vx < -minSpeed) && (vy > minSpeed || vy < -minSpeed) {
				return vx, vy
			}
		}
	}

	if creep.FireTicks == 0 {
		for i := 0; i < creep.FireAmount; i++ {
			vx, vy := randVelocity()

			if creep.Rand.Intn(2) == 0 {
				vx *= -1
			}
			entity.NewCreepBullet(position.X, position.Y, vx, vy)
		}
		creep.FireTicks = creep.FireRate
	}
	creep.FireTicks--

	if creep.DamageTicks > 0 {
		creep.DamageTicks--

		sprite := s.Sprite
		if sprite != nil {
			if creep.DamageTicks > 0 {
				if creep.DamageTicks%2 == 0 {
					sprite.ColorScale = 100
				} else {
					sprite.ColorScale = .01
				}
				sprite.OverrideColorScale = true
			} else {
				sprite.OverrideColorScale = false
			}
		}
	}

	return nil
}

func (_ *CreepSystem) Draw(_ gohan.Entity, _ *ebiten.Image) error {
	return gohan.ErrUnregister
}
