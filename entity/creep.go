package entity

import (
	"math/rand"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"

	"code.rocketnine.space/tslocum/fishfightback/asset"
	"code.rocketnine.space/tslocum/fishfightback/component"
	"code.rocketnine.space/tslocum/fishfightback/level"
	"code.rocketnine.space/tslocum/gohan"
)

var newestCreepID int64
var creepIDLock sync.Mutex

func newCreepID() int64 {
	creepIDLock.Lock()
	defer creepIDLock.Unlock()

	newestCreepID++
	return newestCreepID
}

func randCreepType() int {
	return rand.Intn(8)
}

func NewCreep(creepType int, x, y float64) gohan.Entity {
	creepID := newCreepID()

	creep := gohan.NewEntity()

	creep.AddComponent(&component.Position{
		X: x,
		Y: y,
		Z: level.LayerCreep,
	})

	images := []*ebiten.Image{
		asset.PeepImage(asset.ImgPeepBody, randCreepType(), 0),
	}
	if rand.Intn(3) == 0 {
		if rand.Intn(2) == 0 {
			images = append(images, asset.PeepImage(asset.ImgPeepGlasses, randCreepType(), 0))
		} else {
			images = append(images, asset.PeepImage(asset.ImgPeepSunglasses, randCreepType(), 0))
		}
	}
	if rand.Intn(2) == 0 {
		if rand.Intn(2) == 0 {
			images = append(images, asset.PeepImage(asset.ImgPeepShirt, randCreepType(), 0))
		} else {
			images = append(images, asset.PeepImage(asset.ImgPeepSpaghetti, randCreepType(), 0))
		}
	}
	images = append(images, asset.PeepImage(asset.ImgPeepPants, randCreepType(), 0))

	creep.AddComponent(&component.Sprite{
		Images:  images,
		OffsetX: -16,
		OffsetY: -16,
	})

	creep.AddComponent(&component.Creep{
		Type:       creepType,
		Health:     1,
		FireAmount: 1,
		FireRate:   144 * 1,
		Rand:       rand.New(rand.NewSource(creepID)),
	})

	return creep
}
