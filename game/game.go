package game

import (
	"image/color"
	"math/rand"
	"os"
	"sync"

	"code.rocketnine.space/tslocum/fishfightback/component"
	"code.rocketnine.space/tslocum/fishfightback/entity"

	"code.rocketnine.space/tslocum/fishfightback/asset"
	"code.rocketnine.space/tslocum/fishfightback/level"
	"code.rocketnine.space/tslocum/fishfightback/system"
	"code.rocketnine.space/tslocum/fishfightback/world"
	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

const sampleRate = 44100

type game struct {
	audioContext *audio.Context

	op *ebiten.DrawImageOptions

	disableEsc bool

	debugMode  bool
	cpuProfile *os.File

	movementSystem *system.MovementSystem

	addedSystems bool

	nextSectionX float64

	activeGamepad int

	gameWon bool

	sync.Mutex
}

func NewGame() (*game, error) {
	g := &game{
		audioContext:  audio.NewContext(sampleRate),
		op:            &ebiten.DrawImageOptions{},
		activeGamepad: -1,
	}

	err := g.loadAssets()
	if err != nil {
		panic(err)
	}

	const numEntities = 5000 // TODO tweak
	gohan.Preallocate(numEntities)

	world.World.SectionA = world.NewSection(0, 0)
	world.World.SectionB = world.NewSection(0, 0)

	g.updateCursor()

	return g, nil
}

func (g *game) tileToGameCoords(x, y int) (float64, float64) {
	return float64(x) * 32, float64(y) * 32
}

func (g *game) updateCursor() {
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
}

// Layout is called when the game's layout changes.
func (g *game) Layout(_, _ int) (int, int) {
	return world.ScreenWidth, world.ScreenHeight
}

func (g *game) Update() error {
	if ebiten.IsWindowBeingClosed() {
		g.Exit()
		return nil
	}

	if world.World.ResetGame {
		world.Reset()

		if !g.addedSystems {
			g.addSystems()
			g.addedSystems = true
		}

		if world.World.Player == 0 {
			world.World.Player = entity.NewPlayer()
			world.SetFish(level.FishParrot)
		}

		world.World.Player.With(func(position *component.Position) {
			position.X, position.Y = float64(rand.Intn(world.ScreenWidth/2)), world.ScreenHeight-system.TileWidth*2
		})

		world.World.CamX, world.World.CamY = 0, 0

		g.nextSectionX = 0

		world.World.ResetGame = false
		world.World.GameOver = false
	}

	world.World.Tick++

	s := world.World.SectionA
	last := world.World.SectionB
	if world.World.FirstSectionB {
		s = world.World.SectionB
		last = world.World.SectionA
	}

	// Generate next section by repositioning and regenerating previous section.
	if world.World.CamX+world.ScreenWidth >= g.nextSectionX {
		s.SetPosition(g.nextSectionX, world.World.CamY)
		s.Regenerate(last.ShoreDepth)

		g.nextSectionX += world.SectionWidth
		world.World.FirstSectionB = !world.World.FirstSectionB
	}

	err := gohan.Update()
	if err != nil {
		return err
	}
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	err := gohan.Draw(screen)
	if err != nil {
		panic(err)
	}
}

func (g *game) addSystems() {
	g.movementSystem = system.NewMovementSystem()

	gohan.AddSystem(system.NewRailSystem())
	gohan.AddSystem(system.NewPlayerMoveSystem(world.World.Player, g.movementSystem))
	gohan.AddSystem(system.NewplayerFireSystem())
	gohan.AddSystem(g.movementSystem)
	gohan.AddSystem(system.NewCreepSystem())
	gohan.AddSystem(system.NewCameraSystem())

	for layer := -level.NumLayers + 1; layer <= level.LayerDefault; layer++ {
		gohan.AddSystem(system.NewRenderSystem(layer))
	}

	gohan.AddSystem(system.NewRenderHUDSystem())
	gohan.AddSystem(system.NewRenderMessageSystem())
	gohan.AddSystem(system.NewRenderDebugTextSystem(world.World.Player))
	gohan.AddSystem(system.NewRenderCursorSystem())
	gohan.AddSystem(system.NewProfileSystem(world.World.Player))
}

func (g *game) loadAssets() error {
	asset.ImgWhiteSquare.Fill(color.White)
	asset.LoadSounds(g.audioContext)
	return nil
}

func (g *game) Exit() {
	os.Exit(0)
}
