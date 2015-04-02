package main
import "github.com/veandco/go-sdl2/sdl"

type Gamer interface {
	Draw(g *Game, tm *TextureManager)
	Update()
}

type GameObj struct {
	vec2d Vector2d
	width int32
	height int32
	textureID string
	currentFrame int32
	currentRow   int32
}
func NewGameObj(x float64, y float64, width int32, height int32, textureID string) *GameObj{
	v := Vector2d{}
	//play with setters ))
	v.SetX(x)
	v.SetY(y)
	return &GameObj{v,width,height,textureID,1,1}
}

func (gb *GameObj) Draw(g *Game, tm *TextureManager) {
	tm.Draw(*gb, g, sdl.FLIP_NONE)
}

func (gb *GameObj) Update() {
	gb.vec2d.SetX(gb.vec2d.X() + 1)
}
