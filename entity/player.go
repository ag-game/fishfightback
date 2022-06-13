package entity

import (
	"code.rocketnine.space/tslocum/fishfightback/asset"
	"code.rocketnine.space/tslocum/fishfightback/component"
	"code.rocketnine.space/tslocum/fishfightback/level"
	"code.rocketnine.space/tslocum/gohan"
)

func NewPlayer() gohan.Entity {
	player := gohan.NewEntity()

	player.AddComponent(&component.Position{
		Z: level.LayerPlayer,
	})

	player.AddComponent(&component.Velocity{})

	weapon := &component.Weapon{}
	player.AddComponent(weapon)

	player.AddComponent(&component.Sprite{
		Image:          asset.FishImage(int(level.FishParrot)),
		HorizontalFlip: true,
	})

	player.AddComponent(&component.Rail{})

	return player
}
