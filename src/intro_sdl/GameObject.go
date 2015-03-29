package main

import (
	"fmt"
)

//start do OOP thinks ))

type GameObject struct {
	textureID    string
	currentFrame int32
	currentRow   int32
	x            int32
	y            int32
	width        int32
	height       int32
}

func (gb *GameObject) Load(x int32, y int32, width int32, height int32, textureID string) {
	gb.x = x
	gb.y = y
	gb.width = width
	gb.height = height
	gb.textureID = textureID
	gb.currentRow = 1
	gb.currentFrame = 1
}

//todo pRenderer *sdl.Renderer
func (gb *GameObject) Draw(g *Game, tm *TextureManager) {
	tm.DrawFN(gb.textureID, gb.x, gb.y, gb.width, gb.height, g)
}

func (gb *GameObject) Update() {
	gb.x += 1
}

func (gb *GameObject) Clean() {
	fmt.Println("clean game object")
}
