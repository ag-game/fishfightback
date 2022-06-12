package level

type FishType int

const (
	FishParrot FishType = iota
	FishMackerel
	FishClown
	FishPlaice
)

type FishInfo struct {
	Damage      int
	FireRate    int
	BulletSpeed float64
}

var AllFish = map[FishType]*FishInfo{
	FishParrot: {
		Damage:      1,
		FireRate:    60,
		BulletSpeed: 2,
	},
	FishMackerel: { // TODO
		Damage:      1,
		FireRate:    60,
		BulletSpeed: 2,
	},
}
