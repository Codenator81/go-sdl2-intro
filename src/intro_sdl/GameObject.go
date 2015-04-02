package main
import "github.com/veandco/go-sdl2/sdl"

type Gamer interface {
	Draw(g *Game, tm *TextureManager)
	Update()
}

type GameObj struct {
	LoaderParams
	currentFrame int32
	currentRow   int32
}
func NewGameObj(x int32, y int32, width int32, height int32, textureID string) *GameObj{
	return &GameObj{LoaderParams{x,y,width,height,textureID},1,1}
}

func (gb *GameObj) Draw(g *Game, tm *TextureManager) {
	tm.Draw(*gb, g, sdl.FLIP_NONE)
}

func (gb *GameObj) Update() {
	gb.x += 1
}
