package main

// composite from GameObject
type Player struct {
	GameObj
}

func NewPlayer(x int32, y int32, width int32, height int32, textureID string) *Player{
	return &Player{GameObj{LoaderParams{x,y,width,height,textureID},1,1}}
}

//Delete two functions because we have emedded  GameObj

func (p *Player) Update() {
	p.x -= 1
}
