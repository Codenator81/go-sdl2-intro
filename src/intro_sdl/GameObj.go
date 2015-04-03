package main
import "github.com/veandco/go-sdl2/sdl"

type Gamer interface {
	Draw(g *Game, tm *TextureManager)
	Update()
}

type GameObj struct {
	position Vector2d
	velocity Vector2d
	acceleration Vector2d
	width int32
	height int32
	textureID string
	currentFrame int32
	currentRow   int32
}

func (gb *GameObj) Draw(g *Game, tm *TextureManager) {
	tm.Draw(*gb, g, sdl.FLIP_NONE)
}

func (gb *GameObj) Update() {
	gb.velocity.Add(gb.acceleration)
	gb.position.Add(gb.velocity)
}

//func NewGameObj(x float64, y float64, width int32, height int32, textureID string) *GameObj{
//	position := Vector2d{}
//	velocity  := Vector2d{0,0}
//	acceleration := Vector2d{0,0}
//	position.SetX(x)
//	position.SetY(y)
//	return &GameObj{position, velocity, acceleration, width, height, textureID, 1, 1}
//}