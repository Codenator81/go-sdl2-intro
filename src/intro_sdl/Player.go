package main

// composite from GameObject
type Player struct {
	GameObj
}

func NewPlayer(x float64, y float64, width int32, height int32, textureID string) *Player{
	v := Vector2d{}
	v.SetX(x)
	v.SetY(y)
	return &Player{GameObj{v,width,height,textureID,1,1}}
}

//Delete two functions because we have emedded  GameObj

func (p *Player) Update() {
	p.vec2d.SetX(p.vec2d.X() - 1)

}
