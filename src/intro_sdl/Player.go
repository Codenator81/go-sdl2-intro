package main

import "fmt"

//// inherit from GameObject  class Player extend GameObject
type Player struct {
	GameObject
}

func (p *Player) Draw() {
	fmt.Println("draw Player")
}

func (p *Player) Update() {
	fmt.Println("update Player")
	p.x = 10
	p.y = 20
}

func (p *Player) Clean() {
	fmt.Println("clean Player")
}
