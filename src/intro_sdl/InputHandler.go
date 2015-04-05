package main
import "github.com/veandco/go-sdl2/sdl"

type InputHandler struct {
    keyState uint8
    keyStates []uint8
}

func (in InputHandler) isKeyDown(key sdl.Scancode) bool {
    if len(in.keyStates) != 0 {
        if in.keyStates[key] == 1 {
            return true;
        } else {
            return false;
        }
    }
    return false;
}

func (in InputHandler) Update() {

}



