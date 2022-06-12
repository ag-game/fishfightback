package asset

import (
	"bytes"
	"embed"
	"image"
	"image/color"
	_ "image/png"
	"io"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

const sampleRate = 44100

//go:embed image
var FS embed.FS

var ColorWater = color.RGBA{96, 160, 168, 255}

var (
	ImgWhiteSquare = newFilledImage(4, 4, color.White)
	ImgBlackSquare = newFilledImage(4, 4, color.Black)

	ImgWater = newFilledImage(16, 16, ColorWater)

	ImgCrosshair = LoadImage("image/crosshair.png")

	ImgFishTileset = LoadImage("image/cozy-fishing/global.png")
)

func init() {
	ImgWhiteSquare.Fill(color.White)
	ImgBlackSquare.Fill(color.Black)
}

func newFilledImage(w int, h int, c color.Color) *ebiten.Image {
	img := ebiten.NewImage(w, h)
	img.Fill(c)
	return img
}

func LoadSounds(ctx *audio.Context) {
	// TODO
}

func LoadImage(p string) *ebiten.Image {
	f, err := FS.Open(p)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	baseImg, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(baseImg)
}

func LoadBytes(p string) []byte {
	b, err := FS.ReadFile(p)
	if err != nil {
		panic(err)
	}
	return b
}

func LoadWAV(context *audio.Context, p string) *audio.Player {
	f, err := FS.Open(p)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	stream, err := wav.DecodeWithSampleRate(sampleRate, f)
	if err != nil {
		panic(err)
	}

	player, err := context.NewPlayer(stream)
	if err != nil {
		panic(err)
	}

	// Workaround to prevent delays when playing for the first time.
	player.SetVolume(0)
	player.Play()
	player.Pause()
	player.Rewind()
	player.SetVolume(1)

	return player
}

func LoadOGG(context *audio.Context, p string, loop bool) *audio.Player {
	b := LoadBytes(p)

	stream, err := vorbis.DecodeWithSampleRate(sampleRate, bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	var s io.Reader
	if loop {
		s = audio.NewInfiniteLoop(stream, stream.Length())
	} else {
		s = stream
	}

	player, err := context.NewPlayer(s)
	if err != nil {
		panic(err)
	}

	// Workaround to prevent delays when playing for the first time.
	player.SetVolume(0)
	player.Play()
	player.Pause()
	player.Rewind()
	player.SetVolume(1)

	return player
}

func FishTileAt(i int) *ebiten.Image {
	const tileSize = 16
	const tilesetWidth = 56

	x, y := i%tilesetWidth, i/tilesetWidth

	r := image.Rect(x*tileSize, y*tileSize, (x+1)*tileSize, (y+1)*tileSize)

	return ImgFishTileset.SubImage(r).(*ebiten.Image)
}

func FishTileXY(x, y int) *ebiten.Image {
	const tileSize = 16
	const tilesetWidth = 56

	r := image.Rect(x*tileSize, y*tileSize, (x+1)*tileSize, (y+1)*tileSize)

	return ImgFishTileset.SubImage(r).(*ebiten.Image)
}

func FishImage(i int) *ebiten.Image {
	const tileSize = 16
	const fishTilesetWidth = 10

	x, y := i%fishTilesetWidth, i/fishTilesetWidth

	x += 46

	r := image.Rect(x*tileSize, y*tileSize, (x+1)*tileSize, (y+1)*tileSize)

	return ImgFishTileset.SubImage(r).(*ebiten.Image)
}
