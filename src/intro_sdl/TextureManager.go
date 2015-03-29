package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
	"os"
)

var textureMap = map[string]*sdl.Texture{}

func Load(fileName string, id string) {
	var pTempSurface *sdl.Surface
	var pTexture *sdl.Texture
	pTempSurface, err = img.Load(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load BMP: %s\n", err)
		os.Exit(3)
	}
	pTexture, err = renderer.CreateTextureFromSurface(pTempSurface)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		os.Exit(4)
	}
	defer pTempSurface.Free()
	textureMap[id] = pTexture
}

func Draw(id string, x int32, y int32, width int32, height int32, flip sdl.RendererFlip) {
	var srcRect sdl.Rect
	var destRect sdl.Rect

	srcRect.X = 0
	srcRect.Y = 0
	srcRect.W, destRect.W = width, width
	srcRect.H, destRect.H = height, height
	destRect.X = x
	destRect.Y = y
	renderer.CopyEx(textureMap[id], &srcRect, &destRect, 0, zeroPoint, flip)
}

func DrawFrame(id string, x int32, y int32, width int32, height int32, currentRow int32, currentFrame int32, flip sdl.RendererFlip) {
	var srcRect sdl.Rect
	var destRect sdl.Rect
	srcRect.X = width * currentFrame
	srcRect.Y = height * (currentRow - 1)
	srcRect.W, destRect.W = width, width
	srcRect.H, destRect.H = height, height
	destRect.X = x
	destRect.Y = y

	renderer.CopyEx(textureMap[id], &srcRect, &destRect, 0, zeroPoint, flip)
}

//func for default flip_none
func DrawFN(id string, x int32, y int32, width int32, height int32) {
	Draw(id, x, y, width, height, sdl.FLIP_NONE)
}

//func for default flip_none
func DrawFrameFN(id string, x int32, y int32, width int32, height int32, currentRow, currentFrame int32) {
	DrawFrame(id, x, y, width, height, currentRow, currentFrame, sdl.FLIP_NONE)
}
