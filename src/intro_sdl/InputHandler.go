package main
import (
    "github.com/veandco/go-sdl2/sdl"
)

type InputHandler struct {
    keyState uint8
}
var KeyStates []uint8

    func (in *InputHandler) Update() {
        var event sdl.Event
        for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            switch event.(type) {
                // sdl.QuitEvent quit on window close
                case *sdl.QuitEvent:
                Quit()
                case *sdl.KeyDownEvent:
                println("poll event")
                in.onKeyDown()
                case *sdl.KeyUpEvent:
                in.onKeyUp()
            }
        }
    }

    func (in *InputHandler) isKeyDown(key sdl.Scancode) bool {
        if len(KeyStates) != 0 {
            if KeyStates[key] == 1 {
                println("key pressed")
                return true;
            } else {
                return false;
            }
        }
        return false;
    }

    func (in *InputHandler) onKeyDown() {
        KeyStates = sdl.GetKeyboardState()
    }

    func (in *InputHandler) onKeyUp() {
        KeyStates = sdl.GetKeyboardState()
    }



