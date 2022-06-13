package entity

import (
	"code.rocketnine.space/tslocum/fishfightback/asset"
	"code.rocketnine.space/tslocum/fishfightback/component"
	"code.rocketnine.space/tslocum/fishfightback/level"
	"code.rocketnine.space/tslocum/gohan"
)

func NewCreepBullet(x, y, xSpeed, ySpeed float64) gohan.Entity {
	bullet := gohan.NewEntity()

	bullet.AddComponent(&component.Position{
		X: x,
		Y: y,
		Z: level.LayerBullet,
	})

	bullet.AddComponent(&component.Velocity{
		X: xSpeed,
		Y: ySpeed,
	})

	bullet.AddComponent(&component.Sprite{
		Image: asset.TrashImage(),
	})

	bullet.AddComponent(&component.CreepBullet{})

	bullet.AddComponent(&component.Rail{})

	return bullet
}
