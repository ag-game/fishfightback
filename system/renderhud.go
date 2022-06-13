package system

import (
	"fmt"

	"code.rocketnine.space/tslocum/fishfightback/level"

	"code.rocketnine.space/tslocum/fishfightback/component"
	"code.rocketnine.space/tslocum/fishfightback/world"
	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type RenderHUDSystem struct {
	Position *component.Position
	Velocity *component.Velocity
	Weapon   *component.Weapon

	op     *ebiten.DrawImageOptions
	msgImg *ebiten.Image
	tmpImg *ebiten.Image

	levelUpImg *ebiten.Image

	killInfoWidth float64

	padding float64
}

func NewRenderHUDSystem() *RenderHUDSystem {
	s := &RenderHUDSystem{
		op:         &ebiten.DrawImageOptions{},
		msgImg:     ebiten.NewImage(200, 200),
		tmpImg:     ebiten.NewImage(200, 200),
		levelUpImg: ebiten.NewImage(200, 200),
		padding:    4,
	}

	return s
}

func (s *RenderHUDSystem) Update(_ gohan.Entity) error {
	if world.World.LevelUpTicks != 0 {
		world.World.LevelUpTicks--
		if world.World.LevelUpTicks == 0 {
			world.World.KillInfoUpdated = true
		}
	}
	return nil
}

func (s *RenderHUDSystem) Draw(_ gohan.Entity, screen *ebiten.Image) error {
	if !world.World.GameStarted {
		return nil
	}

	if world.World.ScoreUpdated {
		s.drawScore()
		world.World.ScoreUpdated = false
	}

	if world.World.KillInfoUpdated {
		s.drawLevelUp()
		world.World.KillInfoUpdated = false
	}

	// Draw score.
	s.op.GeoM.Reset()
	s.op.GeoM.Translate(s.padding, world.ScreenHeight-16-s.padding+2)
	screen.DrawImage(s.msgImg, s.op)

	// Draw level-up info.
	s.op.GeoM.Reset()
	s.op.GeoM.Translate(world.ScreenWidth-s.killInfoWidth-s.padding, world.ScreenHeight-16-s.padding+2)
	screen.DrawImage(s.levelUpImg, s.op)
	return nil
}

func (s *RenderHUDSystem) drawLevelUp() {
	var message string
	if world.World.Fish == level.FishAngler {
		message = "MAX EVOLUTION LEVEL!"
	} else if world.World.LevelUpTicks != 0 {
		message = fmt.Sprintf("YOU EVOLVED INTO A %s!", level.AllFish[world.World.Fish].Name)
	} else {
		message = world.NumberPrinter.Sprintf("%d TO GO!", world.World.NeedKills-world.World.Kills)
	}

	split := []string{message}
	width := 0
	for _, line := range split {
		lineSize := len(line) * 6
		if lineSize > width {
			width = lineSize
		}
	}
	s.killInfoWidth = float64(width)

	s.levelUpImg.Clear()
	s.tmpImg.Clear()
	s.op.GeoM.Reset()
	s.op.GeoM.Scale(1, 1)
	ebitenutil.DebugPrint(s.tmpImg, message)
	s.levelUpImg.DrawImage(s.tmpImg, s.op)
	s.op.ColorM.Reset()
}

func (s *RenderHUDSystem) drawScore() {
	message := world.NumberPrinter.Sprintf("%d", world.World.Score)

	s.msgImg.Clear()
	s.tmpImg.Clear()
	s.op.GeoM.Reset()
	s.op.GeoM.Scale(1, 1)
	ebitenutil.DebugPrint(s.tmpImg, message)
	s.msgImg.DrawImage(s.tmpImg, s.op)
	s.op.ColorM.Reset()
}
