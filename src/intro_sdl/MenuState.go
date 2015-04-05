package main
import "fmt"

type MenuState struct {
    GameStater
    menuId string
}

func NewMenuState() *MenuState{
    return &MenuState{menuId: "MENU"}
}
func (ms *MenuState) Update() {}
func (ms *MenuState) Render() {}

func (ms *MenuState) OnEnter() bool {
    fmt.Println("entering MenuState")
    return true
}
func (ms *MenuState) OnExit() bool {
    fmt.Println("exiting MenuState")
    return true
}

func (ms *MenuState) GetStateId() string {
    return ms.menuId
}
