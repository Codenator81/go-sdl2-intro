package main
import (
	"github.com/veandco/go-sdl2/sdl"
)

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

func (p *Player) Update() {
	var in InputHandler
	if in.isKeyDown(sdl.SCANCODE_RIGHT) {
		p.velocity.SetX(2)
		p.GameObj.Update()
	}
	if in.isKeyDown(sdl.SCANCODE_LEFT) {
		p.velocity.SetX(-2)
		p.GameObj.Update()
	}
	if in.isKeyDown(sdl.SCANCODE_UP) {
		p.velocity.SetY(-2)
		p.GameObj.Update()
	}
	if in.isKeyDown(sdl.SCANCODE_DOWN) {
		p.velocity.SetY(2)
		p.GameObj.Update()
	}
	p.velocity.ClearXY()
}
