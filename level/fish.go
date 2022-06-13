package level

type FishType int

const (
	FishParrot FishType = iota
	FishMackerel
	FishClown
	FishPlaice
	FishSilverEel
	FishSeahorse
	FishLionfish
	FishCowfish
	FishTuna
	FishBandedButterflyfish
	FishAtlanticBass
	FishBlueTang
	FishPollock
	FishBallanWrasse
	FishWeaverFish
	FishBream
	FishPufferfish
	FishCod
	FishDab
	FishFlounder
	FishWhiting
	FishHalibut
	FishHerring
	FishStingray
	FishWolfish
	FishBonefish
	FishCobia
	FishBlackDrum
	FishBlobfish
	FishPompano
	FishSardine
	FishAngelfish
	FishRedSnapper
	FishSalmon
	FishAngler
)

type FishInfo struct {
	Name        string
	Damage      int
	FireRate    int
	BulletSpeed float64
}

const baseFireRate = 120
const fireRateIncrement = 10

const bulletSpeedIncrement = 0.1

var AllFish = map[FishType]*FishInfo{
	FishParrot: {
		Name:        "PARROT",
		Damage:      1,
		FireRate:    baseFireRate,
		BulletSpeed: 2,
	},
	FishMackerel: {
		Name:        "MACKEREL",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 1),
		BulletSpeed: 2,
	},
	FishClown: {
		Name:        "CLOWN FISH",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 2),
		BulletSpeed: 2,
	},
	FishPlaice: {
		Name:        "PLAICE",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 3),
		BulletSpeed: 2,
	},
	FishSilverEel: {
		Name:        "SILVER EEL",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 4),
		BulletSpeed: 2,
	},
	FishSeahorse: {
		Name:        "SEAHORSE",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 5),
		BulletSpeed: 2,
	},
	FishLionfish: {
		Name:        "LION FISH",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 6),
		BulletSpeed: 2,
	},
	FishCowfish: {
		Name:        "COW FISH",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 7),
		BulletSpeed: 2,
	},
	FishTuna: {
		Name:        "TUNA",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 8),
		BulletSpeed: 2,
	},
	FishBandedButterflyfish: {
		Name:        "BUTTERFLY FISH",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 9),
		BulletSpeed: 2,
	},
	FishAtlanticBass: {
		Name:        "ATLANTIC BASS",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2,
	},
	FishBlueTang: {
		Name:        "BLUE TANG",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2 - (bulletSpeedIncrement * 1),
	},
	FishPollock: {
		Name:        "POLLOCK",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2 - (bulletSpeedIncrement * 2),
	},
	FishBallanWrasse: {
		Name:        "BALLAN WRASSE",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2 - (bulletSpeedIncrement * 3),
	},
	FishWeaverFish: {
		Name:        "WEAVER FISH",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2 - (bulletSpeedIncrement * 4),
	},
	FishBream: {
		Name:        "BREAM",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2 - (bulletSpeedIncrement * 5),
	},
	FishPufferfish: {
		Name:        "PUFFERFISH",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2 - (bulletSpeedIncrement * 6),
	},
	FishCod: {
		Name:        "COD",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2 - (bulletSpeedIncrement * 7),
	},
	FishDab: {
		Name:        "DAB",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2 - (bulletSpeedIncrement * 8),
	},
	FishFlounder: {
		Name:        "FLOUNDER",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2 - (bulletSpeedIncrement * 9),
	},
	FishWhiting: {
		Name:        "WHITING",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2 - (bulletSpeedIncrement * 10),
	},
	FishHalibut: {
		Name:        "HALIBUT",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2 - (bulletSpeedIncrement * 11),
	},
	FishHerring: {
		Name:        "HERRING",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2 - (bulletSpeedIncrement * 12),
	},
	FishStingray: {
		Name:        "STINGRAY",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2 - (bulletSpeedIncrement * 13),
	},
	FishWolfish: {
		Name:        "WOLFISH",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2 - (bulletSpeedIncrement * 14),
	},
	FishBonefish: {
		Name:        "BONE FISH",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2 - (bulletSpeedIncrement * 15),
	},
	FishCobia: {
		Name:        "COBIA",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2 - (bulletSpeedIncrement * 16),
	},
	FishBlackDrum: {
		Name:        "BLACK DRUM",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2 - (bulletSpeedIncrement * 17),
	},
	FishBlobfish: {
		Name:        "BLOB FISH",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2 - (bulletSpeedIncrement * 18),
	},
	FishPompano: {
		Name:        "POMPANO",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2 - (bulletSpeedIncrement * 19),
	},
	FishSardine: {
		Name:        "SARDINE",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2 - (bulletSpeedIncrement * 20),
	},
	FishAngelfish: {
		Name:        "ANGEL FISH",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2 - (bulletSpeedIncrement * 21),
	},
	FishRedSnapper: {
		Name:        "RED SNAPPER",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2 - (bulletSpeedIncrement * 22),
	},
	FishSalmon: {
		Name:        "SALMON",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2 - (bulletSpeedIncrement * 23),
	},
	FishAngler: {
		Name:        "ANGLER",
		Damage:      1,
		FireRate:    baseFireRate - (fireRateIncrement * 10),
		BulletSpeed: 2 - (bulletSpeedIncrement * 24),
	},
}
