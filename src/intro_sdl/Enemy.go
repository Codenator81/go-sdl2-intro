package main
import "github.com/veandco/go-sdl2/sdl"

// composite from GameObject
type Enemy struct {
    GameObj
}

func NewEnemy(x float64, y float64, width int32, height int32, textureID string) *Enemy{
    position := Vector2d{}
    velocity  := Vector2d{0,0}
    acceleration := Vector2d{0,0}
    position.SetX(x)
    position.SetY(y)
    return &Enemy{GameObj{position, velocity, acceleration, width, height, textureID, 1, 1}}
}

func (e *Enemy) Update() {
    e.position.SetX(e.position.X() + 1)
    e.position.SetY(e.position.Y() +1)
    e.currentFrame = int32(((sdl.GetTicks() / 100) % 6))
}
