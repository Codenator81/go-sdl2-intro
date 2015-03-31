package main
import "github.com/veandco/go-sdl2/sdl"

// composite from GameObject
type Enemy struct {
    GameObj
}

func (e *Enemy) Load(x int32, y int32, width int32, height int32, textureID string) {
    e.GameObj.Load(x, y, width, height, textureID)
}
func (e *Enemy) Draw(g *Game, tm *TextureManager) {
    e.GameObj.Draw(g, tm)
}

func (e *Enemy) Update() {
    e.x += 1
    e.y += 1
    e.currentFrame = int32(((sdl.GetTicks() / 100) % 6))
}

func (p *Enemy) Clean() {}
