package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type World struct {
	chunks map[Position]*Chunk
}

type Chunk struct {
	position Position
	blocks   [][][]*Block // x z y
	collider rl.BoundingBox
	mesh     *ChunkMesh
}

type Block struct {
	data     int8
	collider rl.BoundingBox
}

func (world *World) memoryInit() {
	world.chunks = make(map[Position]*Chunk)
	for xChunk := -WORLD_SIZE; xChunk <= WORLD_SIZE; xChunk++ {
		for zChunk := -WORLD_SIZE; zChunk <= WORLD_SIZE; zChunk++ {
			chunk := &Chunk{}
			world.chunks[Position{xChunk, zChunk}] = chunk
			chunk.position = Position{xChunk, zChunk}
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

type Position struct {
	X int
	Z int
}

type Position3 struct {
	X int
	Z int
	Y int
}

func positionToVector3(position Position) rl.Vector3 {
	return rl.Vector3{X: float32(position.X), Y: 0, Z: float32(position.Z)}
}

func vector3ToPosition(vector3 rl.Vector3) Position {
	return Position{X: int(vector3.X), Z: int(vector3.Z)}
}

func worldToChunkPos(worldPos Position) (chunkPos Position) {
	xChunk := worldPos.X / CHUNK_SIZE
	if worldPos.X < 0 && worldPos.X%CHUNK_SIZE != 0 {
		xChunk--
	}
	zChunk := worldPos.Z / CHUNK_SIZE
	if worldPos.Z < 0 && worldPos.Z%CHUNK_SIZE != 0 {
		zChunk--
	}
	return Position{X: xChunk, Z: zChunk}
}

// local = coordinates of block within chunk
func worldToLocalPos(worldPos Position) (localPos Position) {
	chunkCoords := worldToChunkPos(worldPos)
	localX := worldPos.X - chunkCoords.X*CHUNK_SIZE
	localZ := worldPos.Z - chunkCoords.Z*CHUNK_SIZE
	return Position{X: localX, Z: localZ}
}

func chunkAndLocalToWorldPos(chunkPos, localPos Position) (worldPos Position) {
	return Position{X: chunkPos.X*CHUNK_SIZE + localPos.X, Z: chunkPos.Z*CHUNK_SIZE + localPos.Z}
}

func (world *World) worldGetBlock(x, y, z int) int8 {
	if y < 0 || y >= CHUNK_HEIGHT {
		return AirBlock
	}

	worldPos := Position{X: x, Z: z}
	chunkPos := worldToChunkPos(worldPos)
	chunk, ok := world.chunks[Position{X: chunkPos.X, Z: chunkPos.Z}]
	if !ok {
		return AirBlock
	}

	localPos := worldToLocalPos(worldPos)

	return chunk.blocks[localPos.X][localPos.Z][y].data
}

func (world *World) isBlockSurrounded(x, y, z int) bool {
	if world.worldGetBlock(x-1, y, z) == AirBlock {
		return false
	}
	if world.worldGetBlock(x+1, y, z) == AirBlock {
		return false
	}
	if world.worldGetBlock(x, y-1, z) == AirBlock {
		return false
	}
	if world.worldGetBlock(x, y+1, z) == AirBlock {
		return false
	}
	if world.worldGetBlock(x, y, z-1) == AirBlock {
		return false
	}
	if world.worldGetBlock(x, y, z+1) == AirBlock {
		return false
	}
	return true
}
