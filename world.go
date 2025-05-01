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

var world map[ChunkPos]*Chunk
var noise = opensimplex.New(time.Now().Unix())

// actual number of world is (2*WORLD_SIZE)^2
const WORLD_SIZE = 10

func genWorld() {
	world = make(map[ChunkPos]*Chunk, WORLD_SIZE)
	for x := -WORLD_SIZE; x <= WORLD_SIZE; x++ {
		for z := -WORLD_SIZE; z <= WORLD_SIZE; z++ {
			chunk := &Chunk{}
			chunk.Generate(noise, x, z)
			world[ChunkPos{X: x, Z: z}] = chunk
		}
	}
}

// the amount of world loaded is (2*RENDER_DISTANCE+1)^2
const RENDER_DISTANCE = 3

func renderWorld(renderDistance int) {
	for chunkPos, chunk := range world {
		cameraPos := worldToChunkPos(int(camera3D.Position.X), int(camera3D.Position.Z))
		if chunkPos.X-cameraPos.X <= renderDistance &&
			chunkPos.X-cameraPos.X >= -renderDistance &&
			chunkPos.Z-cameraPos.Z <= renderDistance &&
			chunkPos.Z-cameraPos.Z >= -renderDistance {
			chunk.Render(chunkPos.X, chunkPos.Z)
		}
	}
	rl.DrawGrid(2*WORLD_SIZE, 16)
}

func worldGetBlock(x, y, z int) int8 {
	if y < 0 || y >= CHUNK_HEIGHT {
		return AirBlock
	}

	chunkPos := worldToChunkPos(x, z)
	chunk, ok := world[ChunkPos{X: chunkPos.X, Z: chunkPos.Z}]
	if !ok {
		return AirBlock
	}

	localX := x - chunkPos.X*CHUNK_SIZE
	localZ := z - chunkPos.Z*CHUNK_SIZE

	return chunk.blocks[localX][localZ][y]
}

func worldToChunkPos(x, z int) ChunkPos {
	xChunk := x / CHUNK_SIZE
	if x < 0 && x%CHUNK_SIZE != 0 {
		xChunk--
	}
	zChunk := z / CHUNK_SIZE
	if z < 0 && z%CHUNK_SIZE != 0 {
		zChunk--
	}
	return ChunkPos{X: xChunk, Z: zChunk}
}
