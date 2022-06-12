package entity

import (
	"math/rand"
	"sync"

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

func NewCreep(creepType int, x, y float64) gohan.Entity {
	creepID := newCreepID()

	creep := gohan.NewEntity()

	creep.AddComponent(&component.Position{
		X: x,
		Y: y,
		Z: level.LayerCreep,
	})

	creep.AddComponent(&component.Sprite{
		Image: asset.FishImage(int(level.FishMackerel)),
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
