package game

import (
	"image/color"
	"math/rand"
	"os"
	"sync"
	"time"

	"code.rocketnine.space/tslocum/fishfightback/asset"
	"code.rocketnine.space/tslocum/fishfightback/component"
	"code.rocketnine.space/tslocum/fishfightback/entity"
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

	sectionA *world.Section
	sectionB *world.Section

	firstSectionB bool

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

	g.sectionA = world.NewSection(0, 0)
	g.sectionB = world.NewSection(0, 0)

	g.updateCursor()

	return g, nil
}

func (g *game) tileToGameCoords(x, y int) (float64, float64) {
	return float64(x) * 32, float64(y) * 32
}

func (g *game) updateCursor() {
	if g.activeGamepad != -1 || g.gameWon {
		ebiten.SetCursorMode(ebiten.CursorModeHidden)
		return
	}
	ebiten.SetCursorMode(ebiten.CursorModeVisible)
	ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
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

			if world.World.Player == 0 {
				world.World.Player = entity.NewPlayer()
			}

			const playerStartOffset = 128
			const camStartOffset = 480

			/*	w := float64(world.World.Map.Width * world.World.Map.TileWidth)
				h := float64(world.World.Map.Height * world.World.Map.TileHeight)*/

			world.World.Player.With(func(position *component.Position) {
				//position.X, position.Y = w/2, h-playerStartOffset

				position.X, position.Y = 200, 3500
			})

			world.World.CamX, world.World.CamY = 1, 3700-camStartOffset
		}

		// TODO Seed is configurable
		rand.Seed(time.Now().UnixNano())

		world.World.ResetGame = false
		world.World.GameOver = false
	}

	world.World.Tick++

	// Generate next section by repositioning and regenerating previous section.
	if world.World.CamX+world.ScreenWidth >= g.nextSectionX {
		s := g.sectionA
		last := g.sectionB
		if g.firstSectionB {
			s = g.sectionB
			last = g.sectionA
		}

		s.SetPosition(g.nextSectionX, world.World.CamY)
		s.Regenerate(last.ShoreDepth)

		g.nextSectionX += world.SectionWidth
		g.firstSectionB = !g.firstSectionB
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

	gohan.AddSystem(system.NewPlayerMoveSystem(world.World.Player, g.movementSystem))
	gohan.AddSystem(system.NewplayerFireSystem())
	gohan.AddSystem(g.movementSystem)
	gohan.AddSystem(system.NewCreepSystem())
	gohan.AddSystem(system.NewCameraSystem())
	gohan.AddSystem(system.NewRailSystem())

	for layer := -level.NumLayers + 1; layer <= level.LayerDefault; layer++ {
		gohan.AddSystem(system.NewRenderSystem(layer))
	}

	gohan.AddSystem(system.NewRenderMessageSystem())
	gohan.AddSystem(system.NewRenderDebugTextSystem(world.World.Player))
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
