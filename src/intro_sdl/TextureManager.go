package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
	"os"
)

type TextureManager struct {
	textureMap map[string]*sdl.Texture
	zeroPoint  *sdl.Point
	err        error
}

func (tm *TextureManager) Load(fileName string, id string, g *Game) {
	var pTempSurface *sdl.Surface
	var pTexture *sdl.Texture
	pTempSurface, tm.err = img.Load(fileName)
	if tm.err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load BMP: %s\n", tm.err)
		os.Exit(3)
	}
	pTexture, tm.err = g.renderer.CreateTextureFromSurface(pTempSurface)
	if tm.err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", tm.err)
		os.Exit(4)
	}
	defer pTempSurface.Free()

	tm.textureMap[id] = pTexture
}

func (tm *TextureManager) Draw(id string, x int32, y int32, width int32, height int32, g *Game, flip sdl.RendererFlip) {
	var srcRect sdl.Rect
	var destRect sdl.Rect

	srcRect.X = 0
	srcRect.Y = 0
	srcRect.W, destRect.W = width, width
	srcRect.H, destRect.H = height, height
	destRect.X = x
	destRect.Y = y
	g.renderer.CopyEx(tm.textureMap[id], &srcRect, &destRect, 0, tm.zeroPoint, flip)
}

func (tm *TextureManager) DrawFrame(id string, x int32, y int32, width int32, height int32, currentRow int32, currentFrame int32, g *Game, flip sdl.RendererFlip) {
	var srcRect sdl.Rect
	var destRect sdl.Rect
	srcRect.X = width * currentFrame
	srcRect.Y = height * (currentRow - 1)
	srcRect.W, destRect.W = width, width
	srcRect.H, destRect.H = height, height
	destRect.X = x
	destRect.Y = y

	g.renderer.CopyEx(tm.textureMap[id], &srcRect, &destRect, 0, tm.zeroPoint, flip)
}

//func for default flip_none
func (tm *TextureManager) DrawFN(id string, x int32, y int32, width int32, height int32, g *Game) {
	tm.Draw(id, x, y, width, height, g, sdl.FLIP_NONE)
}

//func for default flip_none
func (tm *TextureManager) DrawFrameFN(id string, x int32, y int32, width int32, height int32, currentRow, currentFrame int32, g *Game) {
	tm.DrawFrame(id, x, y, width, height, currentRow, currentFrame, g, sdl.FLIP_NONE)
}
