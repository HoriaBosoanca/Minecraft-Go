package main

import (
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
			if chunk, exists := chunks[ChunkPos{X: x, Z: z}]; exists {
				chunk.Generate(noise)
			} else {
				chunk := &Chunk{xPos: x, zPos: z}
				chunk.Generate(noise)
				chunks[ChunkPos{X: x, Z: z}] = chunk
			}
		}
	}
}

// the amount of chunks loaded is (2*RENDER_DISTANCE+1)^2
const RENDER_DISTANCE = 1

func renderWorld(renderDistance int) {
	for _, chunk := range chunks {
		if chunk.xPos-int(camera3D.Position.X)/16 <= renderDistance &&
			chunk.xPos-int(camera3D.Position.X)/16 >= -renderDistance &&
			chunk.zPos-int(camera3D.Position.Z)/16 <= renderDistance &&
			chunk.zPos-int(camera3D.Position.Z)/16 >= -renderDistance {
			chunk.Render()
		}
	}
}
