package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

	type Game struct {
		currentFrame int32
		renderer     *sdl.Renderer
		window       *sdl.Window
		err          error
		gameObjs     []Gamer
		//GameStateMachine
		gsm GSM
	}
	func (g *Game) InitGame(title string, xpos int, ypos int, height int, width int, fullscreen bool, tm *TextureManager) {
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
		g.gameObjs = append(g.gameObjs, NewPlayer(300, 300, 128, 82, "animate"))
		g.gameObjs = append(g.gameObjs, NewEnemy(0, 0, 128, 82, "animate"))
		//init game states
		g.gsm.changeState(NewMenuState())
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

	func (g *Game) HandleEvents() {
		var in InputHandler
		in.Update()
		if in.isKeyDown(sdl.SCANCODE_RETURN) {
			g.gsm.changeState(NewPlayState())
		}
	}

	func Quit() {
		sdl.Quit()
		fmt.Println("Exit Game")
		os.Exit(0)
	}
