package world

import (
	"fmt"
	"math"

	"code.rocketnine.space/tslocum/fishfightback/asset"

	"code.rocketnine.space/tslocum/fishfightback/component"
	"code.rocketnine.space/tslocum/fishfightback/level"

	"code.rocketnine.space/tslocum/gohan"
)

const (
	ScreenWidth  = 400
	ScreenHeight = 225
)

const RailSpeed = 0.2

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

	Debug   int
	NoClip  bool
	GodMode bool

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

	SectionA      *Section
	SectionB      *Section
	FirstSectionB bool

	Score int
}

func Reset() {
	for _, e := range gohan.AllEntities() {
		e.Remove()
	}

	World.MessageVisible = false
	World.FirstSectionB = false
	World.Player = 0
	World.Score = 0

	World.SectionA.ShoreDepth = 0
	World.SectionB.ShoreDepth = 0
}

func LevelCoordinatesToScreen(x, y float64) (float64, float64) {
	return (x - World.CamX) * World.CamScale, (y - World.CamY) * World.CamScale
}

func ScreenToLevelCoordinates(x, y float64) (float64, float64) {
	return World.CamX + x, World.CamY + y
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

func SetFish(fish level.FishType) {
	fishInfo := level.AllFish[fish]
	if fishInfo == nil {
		panic(fmt.Sprintf("unknown fish type %d", fish))
	}

	World.Player.With(func(weapon *component.Weapon, sprite *component.Sprite) {
		weapon.Damage = fishInfo.Damage
		weapon.FireRate = fishInfo.FireRate
		weapon.BulletSpeed = fishInfo.BulletSpeed

		sprite.Image = asset.FishImage(int(fish))
	})
}
