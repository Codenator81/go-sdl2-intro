package main

import "fmt"

type PlayState struct {
    GameStater
    playId string
}
func NewPlayState() *PlayState{
return &PlayState{playId: "PLAY"}
}
func (ps *PlayState) Update() {}
func (ps *PlayState) Render() {}

func (ps *PlayState) OnEnter() bool {
    fmt.Println("entering PlayState")
    return true
}
func (ps *PlayState) OnExit() bool {
    fmt.Println("exiting PlayState")
    return true
}

func (ps *PlayState) GetStateId() string {
return ps.playId
}
