package main
import "github.com/veandco/go-sdl2/sdl"

// composite from GameObject
type Player struct {
	GameObj
}

func NewPlayer(x float64, y float64, width int32, height int32, textureID string) *Player{
	position := Vector2d{}
	velocity  := Vector2d{0,0}
	acceleration := Vector2d{0,0}
	position.SetX(x)
	position.SetY(y)
	return &Player{GameObj{position, velocity, acceleration,width, height, textureID, 1, 1}}
}

//Delete two functions because we have emedded  GameObj

func (p *Player) Update() {
	p.currentFrame = int32(((sdl.GetTicks() / 100) % 6))
	p.acceleration.SetX(0.05)
	p.GameObj.Update()
}
