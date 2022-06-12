package system

import (
	"image"
	"image/color"
	"math"

	"code.rocketnine.space/tslocum/fishfightback/component"
	"code.rocketnine.space/tslocum/fishfightback/world"
	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
)

const rewindThreshold = 1

type MovementSystem struct {
	Position *component.Position
	Velocity *component.Velocity

	Creep        *component.Creep        `gohan:"?"`
	CreepBullet  *component.CreepBullet  `gohan:"?"`
	PlayerBullet *component.PlayerBullet `gohan:"?"`
	Sprite       *component.Sprite       `gohan:"?"`
}

func NewMovementSystem() *MovementSystem {
	s := &MovementSystem{}

	return s
}

func drawDebugRect(r image.Rectangle, c color.Color, overrideColorScale bool) gohan.Entity {
	rectEntity := gohan.NewEntity()

	rectImg := ebiten.NewImage(r.Dx(), r.Dy())
	rectImg.Fill(c)

	rectEntity.AddComponent(&component.Position{
		X: float64(r.Min.X),
		Y: float64(r.Min.Y),
	})

	rectEntity.AddComponent(&component.Sprite{
		Image:              rectImg,
		OverrideColorScale: overrideColorScale,
	})

	return rectEntity
}

func (s *MovementSystem) Update(e gohan.Entity) error {
	if !world.World.GameStarted {
		return nil
	}

	if world.World.GameOver && e == world.World.Player {
		return nil
	}

	position := s.Position
	velocity := s.Velocity

	vx, vy := velocity.X, velocity.Y
	if e == world.World.Player && (world.World.NoClip || world.World.Debug != 0) && ebiten.IsKeyPressed(ebiten.KeyShift) {
		vx, vy = vx*2, vy*2
	}

	// Force player to remain within the screen bounds.
	// TODO same for bullets
	if e == world.World.Player {
		for math.Abs(vx)+math.Abs(vy) > moveSpeed*1.5 {
			vx -= vx / 100
			vy -= vy / 100
		}

		position.X, position.Y = position.X+vx, position.Y+vy

		screenX, screenY := s.levelCoordinatesToScreen(position.X, position.Y)
		if screenX < 0 {
			diff := screenX / world.World.CamScale
			position.X -= diff
		} else if screenX > float64(world.ScreenWidth)-world.World.PlayerWidth {
			diff := (float64(world.ScreenWidth) - world.World.PlayerWidth - screenX) / world.World.CamScale
			position.X += diff
		}
		if screenY < 0 {
			diff := screenY / world.World.CamScale
			position.Y -= diff
		} else if screenY > float64(world.ScreenHeight)-world.World.PlayerHeight {
			diff := (float64(world.ScreenHeight) - world.World.PlayerHeight - screenY) / world.World.CamScale
			position.Y += diff
		}

		world.World.PlayerX, world.World.PlayerY = position.X, position.Y

		// Check player hazard collision.
		if world.World.NoClip {
			return nil
		}
		/*playerRect := image.Rect(int(position.X), int(position.Y), int(position.X+world.World.PlayerWidth), int(position.Y+world.World.PlayerHeight))
		for _, r := range world.World.HazardRects {
			if playerRect.Overlaps(r) {
				world.World.SetGameOver()
				return nil
			}
		}*/
	} else {
		position.X, position.Y = position.X+vx, position.Y+vy
	}

	// Check creepBullet collision.
	if world.World.NoClip {
		return nil
	}
	bulletSize := 4.0
	bulletRect := image.Rect(int(position.X), int(position.Y), int(position.X+bulletSize), int(position.Y+bulletSize))

	creepBullet := s.CreepBullet
	playerBullet := s.PlayerBullet

	// Check hazard collisions.
	if creepBullet != nil || playerBullet != nil {
		/* TODO
		for _, hazardRect := range world.World.HazardRects {
			if bulletRect.Overlaps(hazardRect) {
				e.Remove()
				return nil
			}
		} */
	}

	if creepBullet != nil {

		playerRect := image.Rect(int(world.World.PlayerX), int(world.World.PlayerY), int(world.World.PlayerX+world.World.PlayerWidth), int(world.World.PlayerY+world.World.PlayerHeight))

		// TODO if not god mode
		if bulletRect.Overlaps(playerRect) {
			if !world.World.GodMode {
				world.World.SetGameOver()
			}

			e.Remove()
		}
		return nil
	}

	if playerBullet != nil {
		/* TODO
		var hitCreep bool
		for i, creepRect := range world.World.CreepRects {
			if bulletRect.Overlaps(creepRect) {
				creepEntity := world.World.CreepEntities[i]
				creepEntity.With(func(creep *component.Creep) {
					if creep.Active {
						creep.Health--
						creep.DamageTicks = 6
						hitCreep = true
					}
				})

				if hitCreep {
					e.Remove()
					return nil
				}
			}
		} */
	}

	return nil
}

func (s *MovementSystem) levelCoordinatesToScreen(x, y float64) (float64, float64) {
	return (x - world.World.CamX) * world.World.CamScale, (y - world.World.CamY) * world.World.CamScale
}

func (_ *MovementSystem) Draw(_ gohan.Entity, _ *ebiten.Image) error {
	return gohan.ErrUnregister
}
