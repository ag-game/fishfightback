package system

import (
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

	lastScore int
}

func NewRenderHUDSystem() *RenderHUDSystem {
	s := &RenderHUDSystem{
		op:        &ebiten.DrawImageOptions{},
		msgImg:    ebiten.NewImage(1, 1),
		tmpImg:    ebiten.NewImage(200, 200),
		lastScore: -1,
	}

	return s
}

func (s *RenderHUDSystem) Update(_ gohan.Entity) error {
	return gohan.ErrUnregister
}

func (s *RenderHUDSystem) Draw(_ gohan.Entity, screen *ebiten.Image) error {
	if !world.World.GameStarted {
		return nil
	}

	if world.World.Score != s.lastScore {
		s.drawScore()
	}

	// Draw score.
	s.op.GeoM.Reset()
	s.op.GeoM.Translate(1, world.ScreenHeight-15)
	screen.DrawImage(s.msgImg, s.op)
	return nil
}

func (s *RenderHUDSystem) drawScore() {
	message := world.NumberPrinter.Sprintf("%d", world.World.Score)

	split := []string{message}
	width := 0
	for _, line := range split {
		lineSize := len(line) * 12
		if lineSize > width {
			width = lineSize
		}
	}
	height := len(split) * 32

	const padding = 8
	width, height = width+padding*2, height+padding*2

	s.msgImg = ebiten.NewImage(width, height)
	//s.msgImg.Fill(color.RGBA{17, 17, 17, 100})

	s.tmpImg.Clear()
	s.tmpImg = ebiten.NewImage(width*2, height*2)
	s.op.GeoM.Reset()
	s.op.GeoM.Scale(1, 1)
	//s.op.GeoM.Translate(float64(padding), float64(padding))
	ebitenutil.DebugPrint(s.tmpImg, message)
	s.msgImg.DrawImage(s.tmpImg, s.op)
	s.op.ColorM.Reset()
}
