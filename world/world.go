package world

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"code.rocketnine.space/tslocum/fishfightback/asset"

	"code.rocketnine.space/tslocum/fishfightback/component"
	"code.rocketnine.space/tslocum/fishfightback/level"

	"code.rocketnine.space/tslocum/gohan"
)

const (
	ScreenWidth  = 400
	ScreenHeight = 225
)

const (
	StartingRailSpeed = 0.6
	FishSpeedIncrease = 0.05
)

var RailSpeed = StartingRailSpeed

var NumberPrinter = message.NewPrinter(language.English)

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

	Debug      int
	StartMuted bool
	NoClip     bool
	GodMode    bool

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

	ForceSeed int64

	Tick int

	SectionA      *Section
	SectionB      *Section
	FirstSectionB bool

	Score        int
	ScoreUpdated bool

	Fish         level.FishType
	Kills        int
	NeedKills    int
	LevelUpTicks int

	KillInfoUpdated bool
}

func Reset() {
	for _, e := range gohan.AllEntities() {
		e.Remove()
	}

	World.MessageVisible = false
	World.FirstSectionB = false
	World.Player = 0
	World.Score = 0
	World.ScoreUpdated = true

	World.SectionA.ShoreDepth = 0
	World.SectionB.ShoreDepth = 0

	RailSpeed = 0.4

	World.Kills = 0
	World.Fish = level.FishParrot
	World.NeedKills = NeededKills()
	World.KillInfoUpdated = true

	seed := World.ForceSeed
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	rand.Seed(seed)

	log.Printf("Starting game with seed %d", seed)
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

	log.Printf("Game over - score %d", w.Score)
}

func StartGame() {
	if World.GameStarted {
		return
	}
	World.GameStarted = true

	if !World.StartMuted {
		asset.SoundMusic.Play()
	}
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

	World.Fish = fish

	World.Player.With(func(weapon *component.Weapon, sprite *component.Sprite) {
		weapon.Damage = fishInfo.Damage
		weapon.FireRate = fishInfo.FireRate
		weapon.BulletSpeed = fishInfo.BulletSpeed

		sprite.Image = asset.FishImage(int(fish))
	})

	RailSpeed = StartingRailSpeed + (FishSpeedIncrease * float64(fish))

	World.NeedKills = NeededKills()
}

func MaxCreeps() int {
	const minCreeps = 4
	const levelUpSeconds = 7
	level := World.Tick / (144 * levelUpSeconds)

	maxCreeps := minCreeps + math.Pow(2, float64(level)/4)
	log.Println("level", level, maxCreeps)
	return int(maxCreeps)
}

func NeededKills() int {
	const minCreeps = 7
	level := int(World.Fish)

	maxCreeps := minCreeps + level*7
	log.Println("need creep", level, maxCreeps)
	return int(maxCreeps)
}

func LevelUp() {
	SetFish(World.Fish + 1)

	World.Kills = 0
	World.LevelUpTicks = 144 * 2

	World.Score += int(World.Fish) * 1000
	World.ScoreUpdated = true
}
