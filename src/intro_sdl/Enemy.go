package main
import "github.com/veandco/go-sdl2/sdl"

// composite from GameObject
type Enemy struct {
    GameObj
}

func NewEnemy(x float64, y float64, width int32, height int32, textureID string) *Enemy{
    v := Vector2d{}
    v.SetX(x)
    v.SetY(y)
    return &Enemy{GameObj{v,width,height,textureID,1,1}}
}

func (e *Enemy) Update() {
    e.vec2d.SetX(e.vec2d.X() + 1)
    e.vec2d.SetY(e.vec2d.Y() +1)
    e.currentFrame = int32(((sdl.GetTicks() / 100) % 6))
}
