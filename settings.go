package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// laptop screen sizes
var screenX int32 = 1850
var screenY int32 = 1010

var fps int32 = 20

var camera3D = rl.Camera{
	Position:   rl.NewVector3(1.0, 40.0, 1.0),
	Target:     rl.NewVector3(0.0, 0.0, 0.0),
	Up:         rl.NewVector3(0.0, 1.0, 0.0),
	Fovy:       90.0,
	Projection: rl.CameraPerspective,
}
