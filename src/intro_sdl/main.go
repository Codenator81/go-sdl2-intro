package main
import (
    "os"
    "fmt"

    "github.com/veandco/go-sdl2/sdl"
)

var window *sdl.Window
var renderer *sdl.Renderer
var err error
var event sdl.Event
var gameRunning bool

func main() {
    initGraph("Game SDL 2", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
    640, 480, sdl.WINDOW_SHOWN)
    defer window.Destroy()
    defer renderer.Destroy()
    gameRunning = true
    for gameRunning {
        gameEvent()
        render()
    }
    sdl.Quit()
}

    func initGraph(title string, xpos int, ypos int, height int, width int, flags uint32) {
        // initialize SDL
        sdl.Init(sdl.INIT_EVERYTHING)
        window, err := sdl.CreateWindow(title, xpos, ypos,
        height, width, flags)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
            os.Exit(1)
        }

        renderer, err = sdl.CreateRenderer(window, -1, 0)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
            os.Exit(2)
        }
    }

    func render() {
        renderer.SetDrawColor(0, 0, 0, 255)
        renderer.Present()
        renderer.Clear()
    }

    func gameEvent() {
        for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            switch event.(type) {
                // sdl.QuitEvent quit on window close
                case *sdl.QuitEvent:
                    gameRunning = false
            }
        }
    }