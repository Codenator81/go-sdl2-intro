package main
import "fmt"

type GSM struct{
    gStates []GameStater
}

func (g *GSM) pushState(state GameStater) {
    // push back our new state
    g.gStates = append(g.gStates, state)
    // initialise it
    g.gStates[len(g.gStates)-1].OnEnter()
}
func (g GSM) changeState(state GameStater) {
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
func (g GSM) popState() {
    if len(g.gStates) != 0 {
        if g.gStates[len(g.gStates)-1].OnExit() {
            tmp, states := g.gStates[len(g.gStates)-1], g.gStates[:len(g.gStates)-1]
            g.gStates = append(states, tmp)
        }
    }
}
