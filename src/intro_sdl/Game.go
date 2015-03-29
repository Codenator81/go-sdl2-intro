package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

type Game struct {
	gameRunning  bool
	currentFrame int32
	renderer     *sdl.Renderer
	window       *sdl.Window
	event        sdl.Event
	err          error
	gb           GameObject
	player       Player
}

func (g *Game) HandleEvents() {
	for g.event = sdl.PollEvent(); g.event != nil; g.event = sdl.PollEvent() {
		switch g.event.(type) {
		// sdl.QuitEvent quit on window close
		case *sdl.QuitEvent:
			g.gameRunning = false
		}
	}
}

func (g *Game) InitGraph(title string, xpos int, ypos int, height int, width int, fullscreen bool, tm *TextureManager) {
	// initialize SDL
	sdl.Init(sdl.INIT_EVERYTHING)

	var flags uint32 = sdl.WINDOW_SHOWN
	if fullscreen {
		flags = sdl.WINDOW_FULLSCREEN_DESKTOP
	}
	g.window, g.err = sdl.CreateWindow(title, xpos, ypos,
		height, width, flags)
	if g.err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", g.err)
		os.Exit(1)
	}

	g.renderer, g.err = sdl.CreateRenderer(g.window, -1, 0)
	if g.err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", g.err)
		os.Exit(2)
	}
	tm.Load(AssetsPath("animate-alpha.png"), "animate", g)
	g.gb.Load(100, 100, 128, 82, "animate")
	g.player.Load(300, 300, 128, 82, "animate")
}

func (g *Game) Render(tm *TextureManager) {
	g.renderer.Clear()
	g.gb.Draw(g, tm)
	g.player.Draw(g, tm)
	g.renderer.Present()
}

func (g *Game) Update() {
	//move sprite for animation
	g.gb.Update()
	g.player.Update()
}

func (g *Game) Clean() {
	g.window.Destroy()
	g.renderer.Destroy()
	sdl.Quit()
}
