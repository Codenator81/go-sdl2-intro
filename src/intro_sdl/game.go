package main

import (
    "github.com/veandco/go-sdl2/sdl"
    "fmt"
    "os"
)

func HandleEvents() {
    for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
        switch event.(type) {
            // sdl.QuitEvent quit on window close
            case *sdl.QuitEvent:
            gameRunning = false
        }
    }
}

func InitGraph(title string, xpos int, ypos int, height int, width int, fullscreen bool) {
    // initialize SDL
    sdl.Init(sdl.INIT_EVERYTHING)

    var flags uint32 = sdl.WINDOW_SHOWN
    if fullscreen {
        flags = sdl.WINDOW_FULLSCREEN_DESKTOP
    }
    window, err = sdl.CreateWindow(title, xpos, ypos,
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

func Render() {
    renderer.SetDrawColor(0, 0, 0, 255)
    renderer.Present()
    renderer.Clear()
}

func Update() {}

func Clean() {
    window.Destroy()
    renderer.Destroy()
    sdl.Quit()
}