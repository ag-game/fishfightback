package system

import (
	"math"
	"os"

	"code.rocketnine.space/tslocum/fishfightback/asset"

	"code.rocketnine.space/tslocum/fishfightback/component"
	"code.rocketnine.space/tslocum/fishfightback/world"
	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	moveSpeed      = 0.6
	accelSpeed     = 0.01
	decelRate      = 40
	decelThreshold = 0.01
)

type playerMoveSystem struct {
	Position *component.Position
	Velocity *component.Velocity
	Weapon   *component.Weapon
	Sprite   *component.Sprite

	player       gohan.Entity
	movement     *MovementSystem
	lastWalkDirL bool

	rewindTicks    int
	nextRewindTick int
}

func NewPlayerMoveSystem(player gohan.Entity, m *MovementSystem) *playerMoveSystem {
	return &playerMoveSystem{
		player:   player,
		movement: m,
	}
}

func (s *playerMoveSystem) Update(e gohan.Entity) error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) && !world.World.DisableEsc {
		os.Exit(0)
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyV) {
		v := 1
		if ebiten.IsKeyPressed(ebiten.KeyShift) {
			v = 2
		}
		if world.World.Debug == v {
			world.World.Debug = 0
		} else {
			world.World.Debug = v
		}
		return nil
	}
	if ebiten.IsKeyPressed(ebiten.KeyControl) && inpututil.IsKeyJustPressed(ebiten.KeyN) {
		world.World.NoClip = !world.World.NoClip
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		if asset.SoundMusic.IsPlaying() {
			asset.SoundMusic.Pause()
		} else {
			asset.SoundMusic.Play()
		}
		return nil
	}

	if !world.World.GameStarted {
		if ebiten.IsKeyPressed(ebiten.KeyEnter) || ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			world.StartGame()
		}
		return nil
	}

	if world.World.GameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			world.World.ResetGame = true
		}
		return nil
	}

	pressLeft := ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA)
	pressRight := ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD)
	pressUp := ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW)
	pressDown := ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS)

	if (pressLeft && !pressRight) ||
		(pressRight && !pressLeft) {
		if pressLeft {
			s.Velocity.X -= accelSpeed
			if s.Velocity.X < -moveSpeed {
				s.Velocity.X = -moveSpeed
			}

			s.Sprite.HorizontalFlip = false
		} else {
			s.Velocity.X += accelSpeed
			if s.Velocity.X > moveSpeed {
				s.Velocity.X = moveSpeed
			}

			s.Sprite.HorizontalFlip = true
		}
	} else if s.Velocity.X != 0 {
		s.Velocity.X -= s.Velocity.X / decelRate
		if s.Velocity.X > -decelThreshold && s.Velocity.X < decelThreshold {
			s.Velocity.X = 0
		}
	}

	if (pressUp && !pressDown) ||
		(pressDown && !pressUp) {
		if pressUp {
			s.Velocity.Y -= accelSpeed
			if s.Velocity.Y < -moveSpeed {
				s.Velocity.Y = -moveSpeed
			}
		} else {
			s.Velocity.Y += accelSpeed
			if s.Velocity.Y > moveSpeed {
				s.Velocity.Y = moveSpeed
			}
		}
	} else if s.Velocity.Y != 0 {
		s.Velocity.Y -= s.Velocity.Y / decelRate
		if s.Velocity.Y > -decelThreshold && s.Velocity.Y < decelThreshold {
			s.Velocity.Y = 0
		}
	}

	return nil
}

func (s *playerMoveSystem) Draw(_ gohan.Entity, _ *ebiten.Image) error {
	return gohan.ErrUnregister
}

func deltaXY(x1, y1, x2, y2 float64) (dx float64, dy float64) {
	dx, dy = x1-x2, y1-y2
	if dx < 0 {
		dx *= -1
	}
	if dy < 0 {
		dy *= -1
	}
	return dx, dy
}

func angle(x1, y1, x2, y2 float64) float64 {
	return math.Atan2(y1-y2, x1-x2)
}
