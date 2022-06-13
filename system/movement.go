package system

import (
	"image"
	"math"
	"math/rand"

	"code.rocketnine.space/tslocum/fishfightback/asset"

	"code.rocketnine.space/tslocum/fishfightback/component"
	"code.rocketnine.space/tslocum/fishfightback/world"
	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
)

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

func (s *MovementSystem) Update(e gohan.Entity) error {
	if !world.World.GameStarted {
		return nil
	}

	if world.World.GameOver {
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

		shoreY := func(s *world.Section) float64 {
			return float64(s.ShoreDepth * TileWidth)
		}

		var currentSection = world.World.SectionA
		if position.X >= world.World.SectionA.X+world.SectionWidth {
			currentSection = world.World.SectionB
		}

		const shoreBuffer = 4
		px, py := world.LevelCoordinatesToScreen(position.X, position.Y)
		if py-shoreY(currentSection) < -shoreBuffer {
			if !world.World.NoClip {
				if !world.World.GodMode {
					asset.SoundDie.Rewind()
					asset.SoundDie.Play()

					world.World.SetGameOver()
				} else {
					_, newY := world.ScreenToLevelCoordinates(px, shoreY(currentSection)-shoreBuffer)
					position.Y = newY
				}
			}
			return nil
		}
	} else {
		position.X, position.Y = position.X+vx, position.Y+vy
	}

	// Check creepBullet collision.
	if world.World.NoClip {
		return nil
	}

	creepBullet := s.CreepBullet
	playerBullet := s.PlayerBullet

	if creepBullet != nil {
		playerRect := image.Rect(int(world.World.PlayerX), int(world.World.PlayerY), int(world.World.PlayerX+world.World.PlayerWidth), int(world.World.PlayerY+world.World.PlayerHeight))
		bulletSize := 4.0

		var currentSection = world.World.SectionA
		if world.World.SectionB.X != 0 && position.X >= world.World.SectionB.X && position.X < world.World.SectionB.X+world.SectionWidth {
			currentSection = world.World.SectionB
		}
		tx, ty := int((position.X-currentSection.X)/TileWidth), int((position.Y-currentSection.Y)/TileWidth)
		if tx < 0 || ty < 0 || tx >= world.SectionWidth/TileWidth || ty >= world.ScreenHeight/TileWidth {
			e.Remove()
			return nil
		}

		bulletRect := image.Rect(int(position.X), int(position.Y), int(position.X+bulletSize), int(position.Y+bulletSize))
		if bulletRect.Overlaps(playerRect) {
			if !world.World.GodMode {
				asset.SoundDie.Rewind()
				asset.SoundDie.Play()

				world.World.SetGameOver()
			}
		}
		return nil
	}

	if playerBullet != nil {
		var offscreen bool
		for oy := -2; oy < 5; oy++ {
			for ox := -2; ox < 5; ox++ {
				var currentSection = world.World.SectionA
				if world.World.SectionB.X != 0 && position.X >= world.World.SectionB.X && position.X < world.World.SectionB.X+world.SectionWidth {
					currentSection = world.World.SectionB
				}
				tx, ty := int((position.X+float64(ox)-currentSection.X)/TileWidth), int((position.Y+float64(oy)-currentSection.Y)/TileWidth)
				if tx < 0 || ty < 0 || tx >= world.SectionWidth/TileWidth || ty >= world.ScreenHeight/TileWidth {
					continue // Offscreen
				}
				offscreen = false

				// Hit creep.
				for offset := 0; offset < 2; offset++ {
					if ty+offset >= world.ScreenHeight/TileWidth {
						break
					}
					creepEntity := currentSection.Creeps[ty+offset][tx]
					if creepEntity != 0 {
						var hitCreep bool
						creepEntity.With(func(creep *component.Creep) {
							if creep == nil || !creep.Active {
								return
							}
							creep.Health--
							creep.DamageTicks = 6

							hitSound := rand.Intn(3)
							switch hitSound {
							case 0:
								asset.SoundHit1.Rewind()
								asset.SoundHit1.Play()
							case 1:
								asset.SoundHit2.Rewind()
								asset.SoundHit2.Play()
							case 2:
								asset.SoundHit3.Rewind()
								asset.SoundHit3.Play()
							}

							if !world.World.GameOver {
								world.World.Kills++
								world.World.Score += 25
								world.World.ScoreUpdated = true
								if world.World.Kills == world.World.NeedKills {
									world.LevelUp()
								}
								world.World.KillInfoUpdated = true
							}

							hitCreep = true
						})
						if hitCreep {
							e.Remove()
							currentSection.Creeps[ty+offset][tx] = 0
						}
					}
				}
			}
		}

		if offscreen {
			e.Remove()
			return nil
		}
	}

	return nil
}

func (s *MovementSystem) levelCoordinatesToScreen(x, y float64) (float64, float64) {
	return (x - world.World.CamX) * world.World.CamScale, (y - world.World.CamY) * world.World.CamScale
}

func (_ *MovementSystem) Draw(_ gohan.Entity, _ *ebiten.Image) error {
	return gohan.ErrUnregister
}
