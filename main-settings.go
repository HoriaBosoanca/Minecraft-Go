package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// --> GRAPHICS

var fps int32 = 500

var camera3D = rl.Camera{
	Position:   rl.NewVector3(1.0, 40.0, 1.0),
	Target:     rl.NewVector3(0.0, 0.0, 0.0),
	Up:         rl.NewVector3(0.0, 1.0, 0.0),
	Fovy:       90.0,
	Projection: rl.CameraPerspective,
}

// --> WORLD

const (
	CHUNK_SIZE      = 16
	CHUNK_HEIGHT    = 32
	WORLD_SIZE      = 16 // actual number of chunks is (2*WORLD_SIZE+1)^2
	RENDER_DISTANCE = 32 // the number of chunks loaded is (2*RENDER_DISTANCE+1)^2
)

// --> PLAYER

const (
	MOVE_SPEED   float32 = 1
	ASCEND_SPEED float32 = 1
)
