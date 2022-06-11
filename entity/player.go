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

	weapon := &component.Weapon{
		Damage:      1,
		FireRate:    144 / 16,
		BulletSpeed: 8,
	}
	player.AddComponent(weapon)

	player.AddComponent(&component.Sprite{
		Image: asset.ImgBat,
	})

	player.AddComponent(&component.Rail{})

	return player
}
