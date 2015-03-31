package main

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

func (gb *GameObj) Load(x int32, y int32, width int32, height int32, textureID string) {
	gb.x = x
	gb.y = y
	gb.width = width
	gb.height = height
	gb.textureID = textureID
	gb.currentRow = 1
	gb.currentFrame = 1
}

func (gb *GameObj) Draw(g *Game, tm *TextureManager) {
	tm.DrawFN(gb.textureID, gb.x, gb.y, gb.width, gb.height, g)
}

func (gb *GameObj) Update() {
	gb.x += 1
}

func (gb *GameObj) Clean() {}
