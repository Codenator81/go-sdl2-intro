package main

import (
	"fmt"
)

//start do OOP thinks ))

type GameObject struct {
	x int
	y int
}

func (gb *GameObject) Draw() {
	fmt.Println("draw game object")
}

func (gb *GameObject) Update() {
	fmt.Println("update game object")
}

func (gb *GameObject) Clean() {
	fmt.Println("clean game object")
}
