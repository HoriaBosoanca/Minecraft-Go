package main

import rl "github.com/gen2brain/raylib-go/raylib"

var atlas rl.Texture2D

const (
	BLOCKS_PER_ATLAS_WIDTH = 8.0
	ATLAS_UNIT             = 1.0 / BLOCKS_PER_ATLAS_WIDTH

	FACE_1_START = 0
	FACE_1_END   = 11

	FACE_2_START = 12
	FACE_2_END   = 23

	FACE_3_START = 24
	FACE_3_END   = 35

	FACE_4_START = 36
	FACE_4_END   = 47

	FACE_5_START = 48
	FACE_5_END   = 59

	FACE_6_START = 60
	FACE_6_END   = 71
)

const ( // textures begin at the top left of the atlas
	GRASS_SIDE_U = ATLAS_UNIT * 0.0
	GRASS_SIDE_V = ATLAS_UNIT * 6.0

	GRASS_TOP_U = ATLAS_UNIT * 0.0
	GRASS_TOP_V = ATLAS_UNIT * 7.0

	DIRT_U = ATLAS_UNIT * 1.0
	DIRT_V = ATLAS_UNIT * 7.0

	STONE_U = ATLAS_UNIT * 4.0
	STONE_V = ATLAS_UNIT * 6.0
)

func loadTextures() {
	atlas = rl.LoadTexture("atlas.png")

	for i := range cubeTexture {
		cubeTexture[i] /= BLOCKS_PER_ATLAS_WIDTH
	}
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
