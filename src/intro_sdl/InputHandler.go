package main
import (
    "github.com/veandco/go-sdl2/sdl"
    "fmt"
)

const (
    M_E_LEFT uint8 = 1
    M_E_MIDDLE uint8 = 2
    M_E_RIGHT uint8 = 3
    M_B_LEFT uint8 = 0
    M_B_MIDDLE uint8 = 1
    M_B_RIGHT uint8 = 2
)

type InputHandler struct {
    keyState uint8
}
var KeyStates []uint8
// Mouse button States
var MsBtnSts [3]bool
// Mouse Position
var msPst Vector2d

    func (in *InputHandler) newIH() {
        // make 3 buttons mouse false default
        println("Initialize mouse ..")
        for i,_ := range MsBtnSts{
            MsBtnSts[i] = false
        }
        // default mouse position 0
        msPst = Vector2d{0,0}
    }

    func (in *InputHandler) Update() {
        var event sdl.Event
        for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            switch t := event.(type) {
                // sdl.QuitEvent quit on window close
                case *sdl.QuitEvent:
                    Quit()
                case *sdl.KeyDownEvent:
                    OnKeyDown()
                    fmt.Println(event)
                case *sdl.KeyUpEvent:
                    OnKeyUp()
                case *sdl.MouseMotionEvent:
                    OnMouseMove(t)
                case *sdl.MouseButtonEvent:
                    if t.State == 1 {
                        OnBMouseDown(t)
                    }
                    if t.State == 0 {
                        OnBMouseUp(t)
                    }
            }
        }
    }

    func IsKeyDown(key sdl.Scancode) bool {
        if len(KeyStates) != 0 {
            if KeyStates[key] == 1 {
                return true;
            } else {
                return false;
            }
        }
        return false;
    }

    func getMouseButtonState(btnNum uint8) bool {
        return MsBtnSts[btnNum]
    }

    func getMousePosition() Vector2d {
        return msPst
    }

    func OnKeyDown() {
        KeyStates = sdl.GetKeyboardState()
    }

    func OnKeyUp() {
        KeyStates = nil
    }

    func OnBMouseDown(t *sdl.MouseButtonEvent) {
        switch t.Button {
            case M_E_LEFT:
                MsBtnSts[M_B_LEFT] = true
            case M_E_MIDDLE:
                MsBtnSts[M_B_MIDDLE] = true
            case M_E_RIGHT:
                MsBtnSts[M_B_RIGHT] = true
        }
    }
    func OnBMouseUp(t *sdl.MouseButtonEvent) {
        switch t.Button {
            case M_E_LEFT:
            MsBtnSts[M_B_LEFT] = false
            case M_E_MIDDLE:
            MsBtnSts[M_B_MIDDLE] = false
            case M_E_RIGHT:
            MsBtnSts[M_B_RIGHT] = false
        }
    }

    func OnMouseMove(t *sdl.MouseMotionEvent) {
        msPst.SetX(float64(t.X))
        msPst.SetY(float64(t.Y))
    }



