package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ojrac/opensimplex-go"
	"time"
)

type ChunkPos struct {
	X int
	Z int
}

var chunks map[ChunkPos]*Chunk
var noise = opensimplex.New(time.Now().Unix())

// actual number of chunks is (2*WORLD_SIZE)^2
const WORLD_SIZE = 10

func genWorld() {
	chunks = make(map[ChunkPos]*Chunk, WORLD_SIZE)
	for x := -WORLD_SIZE; x < WORLD_SIZE; x++ {
		for z := -WORLD_SIZE; z < WORLD_SIZE; z++ {
			chunk := &Chunk{}
			chunk.Generate(noise, x, z)
			chunks[ChunkPos{X: x, Z: z}] = chunk
		}
	}
}

// the amount of chunks loaded is (2*RENDER_DISTANCE+1)^2
const RENDER_DISTANCE = 1

func renderWorld(renderDistance int) {
	for coords, chunk := range chunks {
		if coords.X-int(camera3D.Position.X)/16 <= renderDistance &&
			coords.X-int(camera3D.Position.X)/16 >= -renderDistance &&
			coords.Z-int(camera3D.Position.Z)/16 <= renderDistance &&
			coords.Z-int(camera3D.Position.Z)/16 >= -renderDistance {
			chunk.Render(coords.X, coords.Z)
		}
	}
	rl.DrawGrid(2*WORLD_SIZE, 16)
}
