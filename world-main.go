package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type World struct {
	chunks map[Position2]*Chunk
}

type Chunk struct {
	position Position2
	blocks   [][][]*Block // x z y
	collider rl.BoundingBox
	mesh     *ChunkMesh
}

type Block struct {
	data int8
}

func (world *World) memoryInit() {
	world.chunks = make(map[Position2]*Chunk)
	for xChunk := -WORLD_SIZE; xChunk <= WORLD_SIZE; xChunk++ {
		for zChunk := -WORLD_SIZE; zChunk <= WORLD_SIZE; zChunk++ {
			chunk := &Chunk{}
			world.chunks[Position2{xChunk, zChunk}] = chunk
			chunk.position = Position2{xChunk, zChunk}
			chunk.blocks = make([][][]*Block, CHUNK_SIZE)
			for x := 0; x < CHUNK_SIZE; x++ {
				chunk.blocks[x] = make([][]*Block, CHUNK_SIZE)
				for z := 0; z < CHUNK_SIZE; z++ {
					chunk.blocks[x][z] = make([]*Block, CHUNK_HEIGHT)
					for y := 0; y < CHUNK_HEIGHT; y++ {
						chunk.blocks[x][z][y] = &Block{}
					}
				}
			}
		}
	}
}

type Position2 struct {
	X int
	Z int
}

type Position3 struct {
	X int
	Z int
	Y int
}

// position transformations:

func position3ToVector3(position Position3) rl.Vector3 {
	return rl.Vector3{X: float32(position.X), Y: float32(position.Y), Z: float32(position.Z)}
}

func position2ToVector3(position Position2) rl.Vector3 {
	return rl.Vector3{X: float32(position.X), Y: 0, Z: float32(position.Z)}
}

func vector3ToPosition2(vector3 rl.Vector3) Position2 {
	return Position2{X: int(vector3.X), Z: int(vector3.Z)}
}

func worldPos2ToChunkPos2(worldPos Position2) (chunkPos Position2) {
	xChunk := worldPos.X / CHUNK_SIZE
	if worldPos.X < 0 && worldPos.X%CHUNK_SIZE != 0 {
		xChunk--
	}
	zChunk := worldPos.Z / CHUNK_SIZE
	if worldPos.Z < 0 && worldPos.Z%CHUNK_SIZE != 0 {
		zChunk--
	}
	return Position2{X: xChunk, Z: zChunk}
}

// local = coordinates of block within chunk
func worldPos2ToLocalPos2(worldPos Position2) (localPos Position2) {
	chunkCoords := worldPos2ToChunkPos2(worldPos)
	localX := worldPos.X - chunkCoords.X*CHUNK_SIZE
	localZ := worldPos.Z - chunkCoords.Z*CHUNK_SIZE
	return Position2{X: localX, Z: localZ}
}

func worldPos3ToLocalPos3(worldPos Position3) (localPos Position3) {
	localPos2 := worldPos2ToLocalPos2(Position2{X: worldPos.X, Z: worldPos.Z})
	return Position3{X: localPos2.X, Z: localPos2.Z, Y: worldPos.Y}
}

func chunkPos2AndLocalPos2ToWorldPos2(chunkPos, localPos Position2) (worldPos Position2) {
	return Position2{X: chunkPos.X*CHUNK_SIZE + localPos.X, Z: chunkPos.Z*CHUNK_SIZE + localPos.Z}
}

// utils:

func (world *World) worldGetBlock(x, y, z int) int8 {
	if y < 0 || y >= CHUNK_HEIGHT {
		return AirBlock
	}

	worldPos := Position2{X: x, Z: z}
	chunkPos := worldPos2ToChunkPos2(worldPos)
	chunk, ok := world.chunks[Position2{X: chunkPos.X, Z: chunkPos.Z}]
	if !ok {
		return AirBlock
	}

	localPos := worldPos2ToLocalPos2(worldPos)

	return chunk.blocks[localPos.X][localPos.Z][y].data
}

func (world *World) isBlockSurrounded(x, y, z int) bool {
	if world.worldGetBlock(x-1, y, z) == AirBlock ||
		world.worldGetBlock(x+1, y, z) == AirBlock ||
		world.worldGetBlock(x, y-1, z) == AirBlock ||
		world.worldGetBlock(x, y+1, z) == AirBlock ||
		world.worldGetBlock(x, y, z-1) == AirBlock ||
		world.worldGetBlock(x, y, z+1) == AirBlock {
		return false
	}
	return true
}
