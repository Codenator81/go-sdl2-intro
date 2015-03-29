package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"path/filepath"
)

var currentFrame int32
var renderer *sdl.Renderer
var zeroPoint *sdl.Point

func HandleEvents() {
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		// sdl.QuitEvent quit on window close
		case *sdl.QuitEvent:
			gameRunning = false
		}
	}
}

func InitGraph(title string, xpos int, ypos int, height int, width int, fullscreen bool) {
	// initialize SDL
	sdl.Init(sdl.INIT_EVERYTHING)

	var flags uint32 = sdl.WINDOW_SHOWN
	if fullscreen {
		flags = sdl.WINDOW_FULLSCREEN_DESKTOP
	}
	window, err = sdl.CreateWindow(title, xpos, ypos,
		height, width, flags)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		os.Exit(1)
	}

	renderer, err = sdl.CreateRenderer(window, -1, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(2)
	}
	Load(AssetsPath("animate-alpha.png"), "animate")
}

func Render() {
	renderer.Clear()
	DrawFN("animate", 0, 0, 128, 82)
	DrawFrameFN("animate", 100, 100, 128, 82, 1, currentFrame)
	renderer.Present()
}

func Update() {
	//move sprite for animation
	currentFrame = int32(((sdl.GetTicks() / 100) % 6))
}

func Clean() {
	window.Destroy()
	renderer.Destroy()
	sdl.Quit()
}

// get path to assets dir in current GOPATH
func AssetsPath(filename string) string {
	assetsdirectory := filepath.Dir("assets/")
	imagePath, err := filepath.Abs(assetsdirectory)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to find assets dir: %s\n", err)
		os.Exit(6)
	}
	return imagePath + "/" + filename
}
