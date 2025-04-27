package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var fps int32 = 60

var camera3D = rl.Camera{
	Position:   rl.NewVector3(1.0, 60.0, 1.0),
	Target:     rl.NewVector3(16.0, 30.0, 16.0),
	Up:         rl.NewVector3(0.0, 1.0, 0.0),
	Fovy:       90.0,
	Projection: rl.CameraPerspective,
}

var screenX int32 = 960
var screenY int32 = 540
