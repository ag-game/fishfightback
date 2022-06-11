//go:build !js || !wasm
// +build !js !wasm

package main

import (
	"flag"

	"code.rocketnine.space/tslocum/fishfightback/world"
	"github.com/hajimehoshi/ebiten/v2"
)

func parseFlags() {
	var (
		fullscreen bool
		noSplash   bool
	)
	flag.BoolVar(&fullscreen, "fullscreen", false, "run in fullscreen mode")
	flag.BoolVar(&noSplash, "no-splash", false, "skip splash screen")
	flag.IntVar(&world.World.Debug, "debug", 0, "print debug information")
	flag.Parse()

	if fullscreen {
		ebiten.SetFullscreen(true)
	}

	if noSplash || world.World.Debug > 0 {
		world.StartGame()
		//world.World.MessageVisible = false
	}
}
