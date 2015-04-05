package main
import "github.com/veandco/go-sdl2/sdl"

// composite from GameObject
type Player struct {
	GameObj
	input     InputHandler
}

func NewPlayer(x float64, y float64, width int32, height int32, textureID string) *Player{
	position := Vector2d{}
	velocity  := Vector2d{0,0}
	acceleration := Vector2d{0,0}
	input := InputHandler{}
	position.SetX(x)
	position.SetY(y)
	return &Player{GameObj{position, velocity, acceleration,width, height, textureID, 1, 1},input}
}

//Delete two functions because we have emedded  GameObj

func (p *Player) Update() {
	p.input.keyStates = sdl.GetKeyboardState()
	if p.input.isKeyDown(sdl.SCANCODE_RIGHT) {
		p.velocity.SetX(2)
		p.GameObj.Update()
	}
	if p.input.isKeyDown(sdl.SCANCODE_LEFT) {
		p.velocity.SetX(-2)
		p.GameObj.Update()
	}
	if p.input.isKeyDown(sdl.SCANCODE_UP) {
		p.velocity.SetY(-2)
		p.GameObj.Update()
	}
	if p.input.isKeyDown(sdl.SCANCODE_DOWN) {
		p.velocity.SetY(2)
		p.GameObj.Update()
	}
	p.velocity.ClearXY()

//	p.currentFrame = int32(((sdl.GetTicks() / 100) % 6))
//	p.acceleration.SetX(0.05)

}
