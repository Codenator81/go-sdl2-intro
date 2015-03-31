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
	gb           *GameObj
	player       *Player
	gameObjs     []Gamer
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
	g.gb = &GameObj{}
	g.player = &Player{}
	g.gb.Load(100, 100, 128, 82, "animate")
	g.player.Load(300, 300, 128, 82, "animate")
	//add GameObj and Player to Interface
	g.gameObjs = append(g.gameObjs, g.gb)
	g.gameObjs = append(g.gameObjs, g.player)
}

func (g *Game) Render(tm *TextureManager) {
	g.renderer.Clear()
	// loop through our objects and draw them
	for n, _ := range g.gameObjs {
		g.gameObjs[n].Draw(g, tm)
	}
	g.renderer.Present()
}

func (g *Game) Update() {
	// loop through and update our objects
	for n, _ := range g.gameObjs {
		g.gameObjs[n].Update()
	}
}

func (g *Game) Clean() {
	g.window.Destroy()
	g.renderer.Destroy()
	sdl.Quit()
}
