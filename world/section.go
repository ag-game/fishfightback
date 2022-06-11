package world

import (
	"math/rand"

	"code.rocketnine.space/tslocum/fishfightback/entity"

	"code.rocketnine.space/tslocum/fishfightback/level"

	"code.rocketnine.space/tslocum/fishfightback/asset"
	"code.rocketnine.space/tslocum/fishfightback/component"
	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
)

// TODO tweak

const SectionWidth = ScreenWidth

type Section struct {
	X float64
	Y float64

	Entities []gohan.Entity
}

func NewSection(x, y float64) *Section {
	s := &Section{
		X: x,
		Y: y,
	}

	return s
}

func (s *Section) SetPosition(x float64, y float64) {
	s.X = x
	s.Y = y
}

func (s *Section) Regenerate() {
	for _, e := range s.Entities {
		e.Remove()
	}
	s.Entities = s.Entities[:0]

	maxShoreDepth := 2 + rand.Intn(5)

	// Fill with sea water.
	// TODO optimize
	for y := 0; y < ScreenHeight/16; y++ {
		for x := 0; x < ScreenWidth/16; x++ {
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

	// Generate land.
	var tile *ebiten.Image
	for y := 0; y < maxShoreDepth; y++ {
		for x := 0; x < SectionWidth/16; x++ {
			if y == maxShoreDepth-1 {
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

			e := gohan.NewEntity()
			e.AddComponent(&component.Position{
				X: s.X + float64(x)*16,
				Y: s.Y + float64(y)*16,
				Z: level.LayerLand,
			})
			e.AddComponent(&component.Sprite{
				Image: tile,
			})

			s.Entities = append(s.Entities, e)
		}
	}

	// Generate buildings.

	// TODO bag of random buildings

	addBuildings := rand.Intn(7)

	nextX := s.X
	for j := 0; j < addBuildings; j++ {
		building := buildingFishingShackA
		switch rand.Intn(6) {
		case 0:
			building = buildingFishingShackB
		case 1:
			building = buildingFishingShackC
		case 2:
			building = buildingBar
		case 3:
			building = buildingCrabShackA
		case 4:
			building = buildingCrabShackB
		}
		buildingHeight := len(building)
		buildingWidth := len(building[0])

		remaining := SectionWidth - int(nextX-s.X) - buildingWidth*16
		if remaining <= 0 {
			break
		}

		var bx, by float64
		const attempts = 3
		const buffer = 50.0
		for attempt := 0; attempt < attempts; attempt++ {
			bx, by = nextX+float64(rand.Intn(remaining)), s.Y+float64(rand.Intn(maxShoreDepth*16)-buildingHeight*16)-12
			if bx-nextX >= buffer {
				break
			}
		}

		s.build(building, bx, by, level.LayerBuilding)

		nextX = bx + float64(buildingWidth*16)
	}

	for i := 0; i < 10; i++ {
		x, y := s.X+float64(i)*16, s.Y+100
		creep := entity.NewCreep(0, x, y)

		s.Entities = append(s.Entities, creep)
	}
}

func (s *Section) build(tiles [][][2]int, x, y float64, z int) {
	for ty := range tiles {
		for tx, tile := range tiles[ty] {
			e := gohan.NewEntity()
			e.AddComponent(&component.Position{
				X: x + float64(tx)*16,
				Y: y + float64(ty)*16,
				Z: z,
			})
			e.AddComponent(&component.Sprite{
				Image: asset.FishTileXY(tile[0], tile[1]),
			})
			s.Entities = append(s.Entities, e)
		}
	}
}
