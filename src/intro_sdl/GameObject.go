package main
import "github.com/veandco/go-sdl2/sdl"

type Gamer interface {
	Draw(g *Game, tm *TextureManager)
	Update()
}

type GameObj struct {
	textureID    string
	currentFrame int32
	currentRow   int32
	x            int32
	y            int32
	width        int32
	height       int32
}

func (gb *GameObj) Load(lo LoadOpt) {
	gb.x = lo.x
	gb.y = lo.y
	gb.width = lo.width
	gb.height = lo.height
	gb.textureID = lo.textureID
	gb.currentRow = 1
	gb.currentFrame = 1
}

func (gb *GameObj) Draw(g *Game, tm *TextureManager) {
	tm.Draw(gb, g, sdl.FLIP_NONE)
}

func (gb *GameObj) Update() {
	gb.x += 1
}
