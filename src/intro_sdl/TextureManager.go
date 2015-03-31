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

func (tm *TextureManager) Draw(gb GameObj, g *Game, flip sdl.RendererFlip) {
	var srcRect sdl.Rect
	var destRect sdl.Rect

	srcRect.X = 0
	srcRect.Y = 0
	srcRect.W, destRect.W = gb.width, gb.width
	srcRect.H, destRect.H = gb.height, gb.height
	destRect.X = gb.x
	destRect.Y = gb.y
	g.renderer.CopyEx(tm.textureMap[gb.textureID], &srcRect, &destRect, 0, tm.zeroPoint, flip)
}

func (tm *TextureManager) DrawFrame(gb GameObj, g *Game, flip sdl.RendererFlip) {
	var srcRect sdl.Rect
	var destRect sdl.Rect
	srcRect.X = gb.width * gb.currentFrame
	srcRect.Y = gb.height * (gb.currentRow - 1)
	srcRect.W, destRect.W = gb.width, gb.width
	srcRect.H, destRect.H = gb.height, gb.height
	destRect.X = gb.x
	destRect.Y = gb.y

	g.renderer.CopyEx(tm.textureMap[gb.textureID], &srcRect, &destRect, 0, tm.zeroPoint, flip)
}
