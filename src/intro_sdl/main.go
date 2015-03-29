package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

var window *sdl.Window
var err error
var event sdl.Event
var gameRunning bool

func main() {
	// last parameter bool fullscreen
	InitGraph("Game SDL 2", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		640, 480, false)
	gameRunning = true
	//Game loop
	for gameRunning {
		HandleEvents()
		Update()
		Render()
	}
	Clean()
}
