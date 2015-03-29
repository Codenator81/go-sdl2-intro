package main

import "fmt"

//// inherit from GameObject  class Player extend GameObject
type Player struct {
	GameObject
}

func (p *Player) Load(x int32, y int32, width int32, height int32, textureID string) {
	p.GameObject.Load(x, y, width, height, textureID)
}
func (p *Player) Draw(g *Game, tm *TextureManager) {
	p.GameObject.Draw(g, tm)
}

func (p *Player) Update() {
	p.x -= 1
}

func (p *Player) Clean() {
	fmt.Println("clean player")
}
