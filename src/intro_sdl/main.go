package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"runtime"
)

func main() {
	var game Game
	tm := TextureManager{textureMap: map[string]*sdl.Texture{}}
	runtime.GOMAXPROCS(runtime.NumCPU())
	// last parameter bool fullscreen
	game.InitGraph("Game SDL 2", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		640, 480, false, &tm)
	game.gameRunning = true
	//Game loop
	for game.gameRunning {
		game.HandleEvents()
		game.Update()
		game.Render(&tm)
	}
	game.Clean()
}
