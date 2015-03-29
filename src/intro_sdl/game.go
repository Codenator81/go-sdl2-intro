package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"path/filepath"
)

var pTexture *sdl.Texture
var sourceRectangle sdl.Rect
var destinationRectangle sdl.Rect

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
	var pTempSurface *sdl.Surface
	pTempSurface, err = sdl.LoadBMP(AssetsPath() + "/rider.bmp")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load BMP: %s\n", err)
		os.Exit(3)
	}
	defer pTempSurface.Free()
	pTexture, err = renderer.CreateTextureFromSurface(pTempSurface)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		os.Exit(4)
	}
	//count width and height of image
	//Query return int but we need int32
	var _, _, rWidth, rHeight, err = pTexture.Query()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to query texture: %s\n", err)
		os.Exit(5)
	}
	//convert(cast) int to int32
	sourceRectangle.W, sourceRectangle.H = int32(rWidth), int32(rHeight)

	destinationRectangle.X, sourceRectangle.X = 0, 0
	destinationRectangle.Y, sourceRectangle.Y = 0, 0
	destinationRectangle.W = sourceRectangle.W
	destinationRectangle.H = sourceRectangle.H
    //test to make picture smaller
    sourceRectangle.W = 50
    sourceRectangle.H = 50
}

func Render() {
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	renderer.Copy(pTexture, &sourceRectangle, &destinationRectangle)
	renderer.Present()
}

func Update() {}

func Clean() {
	window.Destroy()
	renderer.Destroy()
	pTexture.Destroy()
	sdl.Quit()
}

// get path to assets dir in current GOPATH
func AssetsPath() string {
	assetsdirectory := filepath.Dir("assets/")
	imagePath, err := filepath.Abs(assetsdirectory)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to find assets dir: %s\n", err)
		os.Exit(6)
	}
	return imagePath
}
