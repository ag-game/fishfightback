package world

import (
	"math/rand"
	"sync"

	"code.rocketnine.space/tslocum/fishfightback/entity"

	"code.rocketnine.space/tslocum/fishfightback/level"

	"code.rocketnine.space/tslocum/fishfightback/asset"
	"code.rocketnine.space/tslocum/fishfightback/component"
	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
)

const SectionWidth = ScreenWidth * 2

type Section struct {
	X float64
	Y float64

	Entities []gohan.Entity

	TileOccupied [][]bool

	ShoreDepth int

	Creeps [][]gohan.Entity

	sync.Mutex
}

func NewSection(x, y float64) *Section {
	s := &Section{
		X: x,
		Y: y,
	}

	s.TileOccupied = make([][]bool, ScreenHeight/16)
	for ty := range s.TileOccupied {
		s.TileOccupied[ty] = make([]bool, SectionWidth/16)
	}

	s.Creeps = make([][]gohan.Entity, ScreenHeight/16)
	for ty := range s.Creeps {
		s.Creeps[ty] = make([]gohan.Entity, SectionWidth/16)
	}

	return s
}

func (s *Section) SetPosition(x float64, y float64) {
	s.Lock()
	defer s.Unlock()

	s.X = x
	s.Y = y
}

func (s *Section) tileAvailable(tx, ty int, creep bool) bool {
	if ty < 0 {
		return true // TODO
	} else if ty > s.ShoreDepth || (!creep && ty == s.ShoreDepth-1) {
		return false
	}

	if tx < 0 || tx >= SectionWidth/16 || ty < 0 || ty >= ScreenHeight/16 {
		return false
	}

	return !s.TileOccupied[ty][tx]
}

func (s *Section) Regenerate(lastShoreDepth int) {
	s.Lock()
	defer s.Unlock()

	for _, e := range s.Entities {
		e.Remove()
	}
	s.Entities = s.Entities[:0]

	for ty := range s.TileOccupied {
		for tx := range s.TileOccupied[ty] {
			s.TileOccupied[ty][tx] = false
		}
	}

	for ty := range s.Creeps {
		for tx := range s.Creeps[ty] {
			s.Creeps[ty][tx] = 0
		}
	}

	const minShoreDepth = 3
	const maxShoreDepth = 6

	var firstShore bool
	if lastShoreDepth == 0 {
		lastShoreDepth = minShoreDepth + rand.Intn(maxShoreDepth-minShoreDepth)
		firstShore = true
	}
	s.ShoreDepth = lastShoreDepth + 1 - rand.Intn(3)
	if s.ShoreDepth < minShoreDepth {
		s.ShoreDepth = minShoreDepth
	} else if s.ShoreDepth > maxShoreDepth {
		s.ShoreDepth = maxShoreDepth
	}

	// Fill with sea water.
	// TODO optimize
	for y := 0; y <= ScreenHeight/16; y++ {
		for x := 0; x < SectionWidth/16; x++ {
			e := gohan.NewEntity()
			e.AddComponent(&component.Position{
				X: s.X + float64(x)*16,
				Y: s.Y + float64(y)*16,
				Z: level.LayerSea,
			})
			e.AddComponent(&component.Sprite{
				Image: asset.ImgWater,
			})

			s.Entities = append(s.Entities, e)
		}
	}

	addTile := func(img *ebiten.Image, tx, ty int) {
		e := gohan.NewEntity()
		e.AddComponent(&component.Position{
			X: s.X + float64(tx)*16,
			Y: s.Y + float64(ty)*16,
			Z: level.LayerLand,
		})
		e.AddComponent(&component.Sprite{
			Image: img,
		})

		s.Entities = append(s.Entities, e)
	}

	// Generate land.
	var tile *ebiten.Image
	for y := 0; y < s.ShoreDepth; y++ {
		for x := 0; x < SectionWidth/16; x++ {
			if y == s.ShoreDepth-1 {
				tile = asset.FishTileXY(1, 14)
			} else {
				switch rand.Intn(4) {
				case 0:
					tile = asset.FishTileXY(5, 9)
				case 1:
					tile = asset.FishTileXY(6, 9)
				case 2:
					tile = asset.FishTileXY(5, 10)
				case 3:
					tile = asset.FishTileXY(6, 10)
				}
			}
			if !firstShore {
				if s.ShoreDepth-lastShoreDepth > 0 {
					if x == 0 && y == lastShoreDepth-1 {
						tile = asset.FishTileXY(4, 13)
					} else if x == 0 && y == lastShoreDepth {
						tile = asset.FishTileXY(0, 14)
					}
				} else if s.ShoreDepth-lastShoreDepth < 0 {
					if x == 0 && y == lastShoreDepth-2 {
						tile = asset.FishTileXY(3, 13)

						addTile(asset.FishTileXY(2, 14), x, y+1)
					}
				}
			}

			addTile(tile, x, y)
		}
	}

	// Generate buildings.
	// TODO bag of random buildings
	addBuildings := rand.Intn(14)
	for j := 0; j < addBuildings; j++ {
		specialBuilding := rand.Intn(4) == 0

		var building [][][2]int
		if specialBuilding {
			switch rand.Intn(3) {
			case 0:
				building = buildingBar
			case 1:
				building = buildingCrabShackA
			case 2:
				building = buildingCrabShackB
			}
		} else {
			switch rand.Intn(3) {
			case 0:
				building = buildingFishingShackA
			case 1:
				building = buildingFishingShackB
			case 2:
				building = buildingFishingShackC
			}
		}
		buildingHeight := len(building)

		tx, ty := rand.Intn(SectionWidth/16), rand.Intn(s.ShoreDepth)-buildingHeight
		if !s.canBuild(building, tx, ty) {
			continue
		}

		s.build(building, tx, ty, level.LayerBuilding)
	}

	// Generate creeps.
	const attempts = 10
	numCreeps := MaxCreeps()
	for i := 0; i < numCreeps; i++ {
		for attempt := 0; attempt < attempts; attempt++ {
			tx, ty := rand.Intn(SectionWidth/16), int(float64(rand.Intn(s.ShoreDepth-1)))
			if !s.tileAvailable(tx, ty, true) || (ty != 0 && !s.tileAvailable(tx, ty-1, true)) || !s.tileAvailable(tx, ty+1, true) || !s.tileAvailable(tx-1, ty, true) || !s.tileAvailable(tx+1, ty, true) {
				continue
			}

			x, y := s.X+float64(tx)*16, s.Y+float64(ty)*16
			creep := entity.NewCreep(0, x, y)

			s.Creeps[ty][tx] = creep
			s.TileOccupied[ty][tx] = true

			s.Entities = append(s.Entities, creep)
			break
		}
	}
}

func (s *Section) canBuild(tiles [][][2]int, tx, ty int) bool {
	for tty := range tiles {
		for ttx, _ := range tiles[tty] {
			tileX, tileY := tx+ttx, ty+tty
			if !s.tileAvailable(tileX, tileY, false) {
				return false
			}
		}
	}
	return true
}

func (s *Section) build(tiles [][][2]int, tx, ty int, z int) {
	for tty := range tiles {
		for ttx, tile := range tiles[tty] {
			tileX, tileY := tx+ttx, ty+tty
			if tileY < 0 {
				continue // Off-screen
			}

			s.TileOccupied[tileY][tileX] = true

			e := gohan.NewEntity()
			e.AddComponent(&component.Position{
				X: s.X + float64(tileX)*16,
				Y: s.Y + float64(tileY)*16,
				Z: z,
			})
			e.AddComponent(&component.Sprite{
				Image: asset.FishTileXY(tile[0], tile[1]),
			})
			s.Entities = append(s.Entities, e)
		}
	}
}
