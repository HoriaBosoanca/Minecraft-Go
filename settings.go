package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ojrac/opensimplex-go"
	"time"
)

var fps int32 = 60

var camera3D = rl.Camera{
	Position:   rl.NewVector3(1.0, 2.0, 1.0),
	Target:     rl.NewVector3(0.0, 1.0, 0.0),
	Up:         rl.NewVector3(0.0, 1.0, 0.0),
	Fovy:       45.0,
	Projection: rl.CameraPerspective,
}

var screenX int32 = 960
var screenY int32 = 540

func Origin() rl.Vector3 {
	return rl.Vector3{X: 0.5, Y: 0.5, Z: 0.5}
}

var ascendSpeed float32 = 0.1

var noise = opensimplex.New(time.Now().Unix())
