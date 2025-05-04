package main

import rl "github.com/gen2brain/raylib-go/raylib"

var atlas rl.Texture2D

func loadTextures() {
	atlas = rl.LoadTexture("assets/atlas.png")
}

var cubeTexture = []float32{
	// face 1
	0.0, 0.0,
	0.0, 1.0,
	1.0, 0.0,
	1.0, 1.0,
	1.0, 0.0,
	0.0, 1.0,

	// face 2
	1.0, 1.0,
	1.0, 0.0,
	0.0, 0.0,
	0.0, 0.0,
	0.0, 1.0,
	1.0, 1.0,

	// face 3
	1.0, 1.0,
	1.0, 0.0,
	0.0, 0.0,
	0.0, 0.0,
	0.0, 1.0,
	1.0, 1.0,

	// face 4
	0.0, 0.0,
	0.0, 1.0,
	1.0, 0.0,
	1.0, 1.0,
	1.0, 0.0,
	0.0, 1.0,

	// face 5
	0.0, 0.0,
	0.0, 1.0,
	1.0, 1.0,
	1.0, 1.0,
	1.0, 0.0,
	0.0, 0.0,

	// face 6
	0.0, 0.0,
	0.0, 1.0,
	1.0, 0.0,
	1.0, 1.0,
	1.0, 0.0,
	0.0, 1.0,
}
