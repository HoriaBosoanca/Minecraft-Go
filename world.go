package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ojrac/opensimplex-go"
	"time"
)

var world map[Position]*Chunk
var noise = opensimplex.New(time.Now().Unix())

// actual number of chunks is (2*WORLD_SIZE+1)^2
const WORLD_SIZE = 32

func generateWorldBlocks() {
	world = make(map[Position]*Chunk, WORLD_SIZE)
	for x := -WORLD_SIZE; x <= WORLD_SIZE; x++ {
		for z := -WORLD_SIZE; z <= WORLD_SIZE; z++ {
			chunk := &Chunk{}
			world[Position{X: x, Z: z}] = chunk
			chunk.generateBlocks(noise, Position{X: x, Z: z})
		}
	}
}

func generateWorldMeshes() {
	for x := -WORLD_SIZE; x <= WORLD_SIZE; x++ {
		for z := -WORLD_SIZE; z <= WORLD_SIZE; z++ {
			world[Position{X: x, Z: z}].generateMesh(Position{X: x, Z: z})
		}
	}
}

// the amount of chunks loaded is (2*RENDER_DISTANCE+1)^2
const RENDER_DISTANCE = 32

func renderWorld(renderDistance int) {
	for chunkPos, chunk := range world {
		cameraWorldPos := Position{X: int(camera3D.Position.X), Z: int(camera3D.Position.Z)}
		cameraChunkPos := cameraWorldPos.worldToChunkPos()
		if chunkPos.X-cameraChunkPos.X <= renderDistance &&
			chunkPos.X-cameraChunkPos.X >= -renderDistance &&
			chunkPos.Z-cameraChunkPos.Z <= renderDistance &&
			chunkPos.Z-cameraChunkPos.Z >= -renderDistance {
			chunk.render()
		}
	}
	rl.DrawGrid(2*WORLD_SIZE, CHUNK_SIZE)
}

func worldGetBlock(x, y, z int) int8 {
	if y < 0 || y >= CHUNK_HEIGHT {
		return AirBlock
	}

	worldPos := Position{X: x, Z: z}
	chunkPos := worldPos.worldToChunkPos()
	chunk, ok := world[Position{X: chunkPos.X, Z: chunkPos.Z}]
	if !ok {
		return AirBlock
	}

	localPos := worldPos.worldToLocalPos()

	return chunk.blocks[localPos.X][localPos.Z][y]
}
