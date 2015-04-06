package main
import "fmt"

type GameStater interface {
    Update()
    Render()

    OnEnter() bool
    OnExit() bool

    GetStateId() string
}

type GSM struct{
    gStates []GameStater
}

func (g *GSM) pushState(state GameStater) {
    // push back our new state
    g.gStates = append(g.gStates, state)
    // initialise it
    g.gStates[len(g.gStates)-1].OnEnter()
}
func (g *GSM) changeState(state GameStater) {
    fmt.Println("gStates", g.gStates)
    if len(g.gStates) != 0 {
        if g.gStates[len(g.gStates)-1].GetStateId() == state.GetStateId() {
            return;
        }
        if g.gStates[len(g.gStates)-1].OnExit() {
            tmp, states := g.gStates[len(g.gStates)-1], g.gStates[:len(g.gStates)-1]
            g.gStates = append(states, tmp)
        }
    }
    g.pushState(state)
    fmt.Println("Change status", state.GetStateId())
}
func (g *GSM) popState() {
    if len(g.gStates) != 0 {
        if g.gStates[len(g.gStates)-1].OnExit() {
            tmp, states := g.gStates[len(g.gStates)-1], g.gStates[:len(g.gStates)-1]
            g.gStates = append(states, tmp)
        }
    }
}
func (g *GSM) Update() {
    if len(g.gStates) != 0 {
        g.gStates[len(g.gStates)-1].Update();
    }
}

func (g *GSM) Render() {
    if len(g.gStates) != 0 {
        g.gStates[len(g.gStates)-1].Render();
    }
}
