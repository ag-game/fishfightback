package world

import (
	"math"

	"code.rocketnine.space/tslocum/gohan"
)

const (
	ScreenWidth  = 400
	ScreenHeight = 225
)

var World = &GameWorld{
	CamScale:     1,
	CamMoving:    true,
	PlayerWidth:  16,
	PlayerHeight: 16,
	ResetGame:    true,
}

type GameWorld struct {
	Player gohan.Entity

	DisableEsc bool

	Debug  int
	NoClip bool

	GameStarted      bool
	GameStartedTicks int
	GameOver         bool

	MessageVisible  bool
	MessageTicks    int
	MessageDuration int
	MessageUpdated  bool
	MessageText     string

	PlayerX, PlayerY float64

	CamX, CamY float64
	CamScale   float64
	CamMoving  bool

	PlayerWidth  float64
	PlayerHeight float64

	ResetGame bool

	resetTipShown bool

	Tick int
}

func Reset() {
	for _, e := range gohan.AllEntities() {
		e.Remove()
	}
	World.Player = 0
	World.MessageVisible = false
}

func LevelCoordinatesToScreen(x, y float64) (float64, float64) {
	return (x - World.CamX) * World.CamScale, (y - World.CamY) * World.CamScale
}

func (w *GameWorld) SetGameOver() {
	if w.GameOver {
		return
	}

	w.GameOver = true

	if !World.resetTipShown {
		SetMessage("  GAME  OVER\n\nRESET: <ENTER>", math.MaxInt)
		World.resetTipShown = true
	} else {
		SetMessage("GAME OVER", math.MaxInt)
	}
}

func StartGame() {
	if World.GameStarted {
		return
	}
	World.GameStarted = true
}

func SetMessage(message string, duration int) {
	World.MessageText = message
	World.MessageVisible = true
	World.MessageUpdated = true
	World.MessageDuration = duration
	World.MessageTicks = 0
}
