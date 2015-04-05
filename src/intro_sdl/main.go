package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"runtime"
	"fmt"
)
// Here we define how many frames per second we want our game to run at
const FPS uint32 = 60
const DELAY_TIME uint32 = uint32(1000.0 / FPS)
func main() {
	runtime.GOMAXPROCS(4)
	var frameStart, frameTime uint32
	var game Game
	tm := TextureManager{textureMap: map[string]*sdl.Texture{}}
	// before last parameter bool fullscreen
	game.InitGame("Game SDL 2", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		640, 480, false, &tm)
	fmt.Println(runtime.NumGoroutine())
	//Game loop
	for {
		//fmt.Println(runtime.NumGoroutine())
		frameStart = sdl.GetTicks()
		game.HandleEvents()
		game.Update()
		game.Render(&tm)
		frameTime = sdl.GetTicks() - frameStart
		if frameTime < DELAY_TIME {
			sdl.Delay(uint32(DELAY_TIME - frameTime))
		}
	}
	game.window.Destroy()
	game.renderer.Destroy()
	Quit()
}
