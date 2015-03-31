package main
import "github.com/veandco/go-sdl2/sdl"

// composite from GameObject
type Enemy struct {
    GameObj
}

func (e *Enemy) Update() {
    e.x += 1
    e.y += 1
    e.currentFrame = int32(((sdl.GetTicks() / 100) % 6))
}
