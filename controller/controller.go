package controller

import (
    "go-webgl/browser"
)

type ActionState struct {
    Up        bool
    Down      bool
    Left      bool
    Right     bool
    TurnLeft  bool
    TurnRight bool

    ShowMap bool
    LockMap bool
    Action  bool
}

type Interface interface {
    Init(dom browser.Document)
    GetState() ActionState
}
