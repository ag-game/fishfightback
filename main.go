package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"code.rocketnine.space/tslocum/fishfightback/game"
	"code.rocketnine.space/tslocum/fishfightback/world"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowTitle("Fish Fight Back")
	ebiten.SetWindowResizable(true)
	// ebiten.SetFullscreen(true) // TODO
	ebiten.SetWindowSize(world.ScreenWidth, world.ScreenHeight)
	ebiten.SetMaxTPS(144)
	ebiten.SetRunnableOnUnfocused(true) // Note - this currently does nothing in ebiten
	ebiten.SetWindowClosingHandled(true)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOn)
	ebiten.SetCursorMode(ebiten.CursorModeHidden)

	g, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	parseFlags()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGINT,
		syscall.SIGTERM)
	go func() {
		<-sigc

		g.Exit()
	}()

	if world.World.Debug == 0 {
		world.SetMessage("POLLUTION... DESTRUCTION...\nTHE FISH HAVE HAD ENOUGH!\nIT'S PAYBACK TIME!\nPRESS <ENTER> TO GET REVENGE!", 144)
	} else {
		world.StartGame()
	}

	err = ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}
}
