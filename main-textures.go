package main

import rl "github.com/gen2brain/raylib-go/raylib"

var atlas rl.Texture2D
var cursor rl.Texture2D

var ( // textures begin at the top left of the atlas
	BLOCKS_PER_ATLAS_WIDTH = float32(8.0)
	ATLAS_UNIT             = 1.0 / BLOCKS_PER_ATLAS_WIDTH

	GRASS_SIDE   = rl.Vector2{X: ATLAS_UNIT * 0.0, Y: ATLAS_UNIT * 6.0}
	GRASS_TOP    = rl.Vector2{X: ATLAS_UNIT * 0.0, Y: ATLAS_UNIT * 7.0}
	DIRT         = rl.Vector2{X: ATLAS_UNIT * 1.0, Y: ATLAS_UNIT * 7.0}
	STONE        = rl.Vector2{X: ATLAS_UNIT * 4.0, Y: ATLAS_UNIT * 6.0}
	OAK_LOG_SIDE = rl.Vector2{X: ATLAS_UNIT * 6.0, Y: ATLAS_UNIT * 6.0}
	OAK_LOG_TOP  = rl.Vector2{X: ATLAS_UNIT * 6.0, Y: ATLAS_UNIT * 7.0}
	OAK_LEAF     = rl.Vector2{X: ATLAS_UNIT * 7.0, Y: ATLAS_UNIT * 6.0}
)

func loadTextures() {
	atlas = rl.LoadTexture("assets/atlas.png")
	cursor = rl.LoadTexture("assets/crosshair.png")
	for i := range cubeTexture {
		cubeTexture[i] /= BLOCKS_PER_ATLAS_WIDTH
	}

	textureMap = map[int8][]rl.Vector2{
		GrassBlock: {
			GRASS_SIDE,
			GRASS_SIDE,
			GRASS_SIDE,
			GRASS_SIDE,
			GRASS_TOP,
			DIRT,
		},
		DirtBlock: {
			DIRT,
			DIRT,
			DIRT,
			DIRT,
			DIRT,
			DIRT,
		},
		StoneBlock: {
			STONE,
			STONE,
			STONE,
			STONE,
			STONE,
			STONE,
		},
		OakLogBlock: {
			OAK_LOG_SIDE,
			OAK_LOG_SIDE,
			OAK_LOG_SIDE,
			OAK_LOG_SIDE,
			OAK_LOG_TOP,
			OAK_LOG_TOP,
		},
		OakLeafBlock: {
			OAK_LEAF,
			OAK_LEAF,
			OAK_LEAF,
			OAK_LEAF,
			OAK_LEAF,
			OAK_LEAF,
		},
	}
}

var textureMap map[int8][]rl.Vector2

// UV coordinates for a block
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
