package component

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Image          *ebiten.Image
	HorizontalFlip bool
	VerticalFlip   bool

	Angle float64

	Overlay            *ebiten.Image
	OverlayX, OverlayY float64 // Overlay offset

	Frame     int
	Frames    []*ebiten.Image
	FrameTime time.Duration
	LastFrame time.Time
	NumFrames int

	DamageTicks int

	OverrideColorScale bool
	ColorScale         float64

	OffsetX, OffsetY float64

	Images []*ebiten.Image
}
