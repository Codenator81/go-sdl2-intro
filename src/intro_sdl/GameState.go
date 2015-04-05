package main

type GameStater interface {
    Update()
    Render()

    OnEnter() bool
    OnExit() bool

    GetStateId() string
}
