package main

// composite from GameObject
type Player struct {
	GameObj
}

//Delete two functions because we have emedded  GameObj

func (p *Player) Update() {
	p.x -= 1
}

func (p *Player) Clean() {}
