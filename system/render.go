package system

import (
	_ "image/png"
	"time"

	"code.rocketnine.space/tslocum/fishfightback/component"
	"code.rocketnine.space/tslocum/fishfightback/world"
	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
)

const TileWidth = 16

type RenderSystem struct {
	Position *component.Position
	Sprite   *component.Sprite

	img *ebiten.Image            `gohan:"-"`
	op  *ebiten.DrawImageOptions `gohan:"-"`

	camScale float64 `gohan:"-"`

	renderer gohan.Entity `gohan:"-"`

	z int
}

func NewRenderSystem(z int) *RenderSystem {
	s := &RenderSystem{
		renderer: gohan.NewEntity(),
		img:      ebiten.NewImage(320, 100),
		op:       &ebiten.DrawImageOptions{},
		camScale: 1,
		z:        z, // Only the specified z-level will be rendered by this system.
	}

	return s
}

func (s *RenderSystem) Update(_ gohan.Entity) error {
	return gohan.ErrUnregister
}

func (s *RenderSystem) levelCoordinatesToScreen(x, y float64) (float64, float64) {
	px, py := world.World.CamX, world.World.CamY
	py *= -1
	return ((x - px) * s.camScale), ((y + py) * s.camScale)
}

// renderSprite renders a sprite on the screen.
func (s *RenderSystem) renderSprite(x float64, y float64, offsetx float64, offsety float64, angle float64, geoScale float64, colorScale float64, alpha float64, hFlip bool, vFlip bool, sprite *ebiten.Image, target *ebiten.Image) int {
	if alpha < .01 || colorScale < .01 {
		return 0
	}

	// Skip drawing off-screen tiles.
	drawX, drawY := s.levelCoordinatesToScreen(x, y)
	const padding = TileWidth * 4
	width, height := float64(TileWidth), float64(TileWidth)
	left := drawX
	right := drawX + width
	top := drawY
	bottom := drawY + height
	if (left < -padding || left > float64(world.ScreenWidth)+padding) || (top < -padding || top > float64(world.ScreenHeight)+padding) ||
		(right < -padding || right > float64(world.ScreenWidth)+padding) || (bottom < -padding || bottom > float64(world.ScreenHeight)+padding) {
		return 0
	}

	s.op.GeoM.Reset()

	if hFlip {
		s.op.GeoM.Scale(-1, 1)
		s.op.GeoM.Translate(TileWidth, 0)
	}
	if vFlip {
		s.op.GeoM.Scale(1, -1)
		s.op.GeoM.Translate(0, TileWidth)
	}

	s.op.GeoM.Scale(geoScale, geoScale)
	// Rotate
	s.op.GeoM.Translate(offsetx, offsety)
	s.op.GeoM.Rotate(angle)
	// Move to current isometric position.
	s.op.GeoM.Translate(x, y)
	// Translate camera position.
	s.op.GeoM.Translate(-world.World.CamX, -world.World.CamY)
	// Zoom.
	s.op.GeoM.Scale(s.camScale, s.camScale)
	// Center.
	//s.op.GeoM.Translate(float64(s.ScreenW/2.0), float64(s.ScreenH/2.0))

	// Fade in.
	if world.World.Tick < 72 {
		colorScale = float64(world.World.Tick) / 72
	}

	s.op.ColorM.Scale(colorScale, colorScale, colorScale, alpha)

	target.DrawImage(sprite, s.op)

	s.op.ColorM.Reset()

	return 1
}

func (s *RenderSystem) Draw(e gohan.Entity, screen *ebiten.Image) error {
	position := s.Position
	sprite := s.Sprite

	if position.Z != s.z {
		return nil
	}

	if sprite.NumFrames > 0 && time.Since(sprite.LastFrame) > sprite.FrameTime {
		sprite.Frame++
		if sprite.Frame >= sprite.NumFrames {
			sprite.Frame = 0
		}
		sprite.Image = sprite.Frames[sprite.Frame]
		sprite.LastFrame = time.Now()
	}

	colorScale := 1.0
	if sprite.OverrideColorScale {
		colorScale = sprite.ColorScale
	}

	s.renderSprite(position.X, position.Y, 0, 0, sprite.Angle, 1.0, colorScale, 1.0, sprite.HorizontalFlip, sprite.VerticalFlip, sprite.Image, screen)
	return nil
}
