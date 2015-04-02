package main
import "github.com/veandco/go-sdl2/sdl"

// composite from GameObject
type Enemy struct {
    GameObj
}

func NewEnemy(x int32, y int32, width int32, height int32, textureID string) *Enemy{
    return &Enemy{GameObj{LoaderParams{x,y,width,height,textureID},1,1}}
}

func (e *Enemy) Update() {
    e.x += 1
    e.y += 1
    e.currentFrame = int32(((sdl.GetTicks() / 100) % 6))
}
