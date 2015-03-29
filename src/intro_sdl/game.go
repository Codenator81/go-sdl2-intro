package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
	"os"
	"path/filepath"
)

var pTexture *sdl.Texture
var srcRect sdl.Rect
var dstRect sdl.Rect

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
	//load animate sprite from book in PNG
	pTempSurface, err = img.Load(AssetsPath() + "/animate.png")
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

	dstRect.X, srcRect.X = 0, 0
	dstRect.Y, srcRect.Y = 0, 0
	//get first picture in sprite
	dstRect.W, srcRect.W = 128, 128
	dstRect.H, srcRect.H = 82, 82
}

func Render() {
	renderer.Clear()
	var zeroPoint *sdl.Point
	renderer.CopyEx(pTexture, &srcRect, &dstRect, 0, zeroPoint, sdl.FLIP_HORIZONTAL)
	renderer.Present()
}

func Update() {
	//move sprite for animation
	srcRect.X = 128 * int32(((sdl.GetTicks() / 100) % 6))
}

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
