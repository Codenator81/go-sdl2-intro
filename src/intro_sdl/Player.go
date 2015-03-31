package main

// composite from GameObject
type Player struct {
	GameObj
}

func (p *Player) Load(x int32, y int32, width int32, height int32, textureID string) {
	p.GameObj.Load(x, y, width, height, textureID)
}
func (p *Player) Draw(g *Game, tm *TextureManager) {
	p.GameObj.Draw(g, tm)
}

func (p *Player) Update() {
	p.x -= 1
}

func (p *Player) Clean() {}
