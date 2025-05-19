package main

import (
	"fmt"
	"github.com/ojrac/opensimplex-go"
	"time"
)

var noise = opensimplex.New(time.Now().Unix())

const craziness = 0.03

const (
	AirBlock int8 = iota
	GrassBlock
	DirtBlock
	StoneBlock
)

func (world *World) generateWorldBlocks() {
	world.chunks = make(map[Position]*Chunk, WORLD_SIZE)
	for x := -WORLD_SIZE; x <= WORLD_SIZE; x++ {
		for z := -WORLD_SIZE; z <= WORLD_SIZE; z++ {
			chunk := &Chunk{}
			world.chunks[Position{X: x, Z: z}] = chunk
			chunk.generateBlocks(Position{X: x, Z: z})
		}
	}
}

func (chunk *Chunk) generateBlocks(chunkPos Position) {
	chunk.blocks = make([][][]int8, CHUNK_SIZE)
	for x := 0; x < CHUNK_SIZE; x++ {
		chunk.blocks[x] = make([][]int8, CHUNK_SIZE)
		for z := 0; z < CHUNK_SIZE; z++ {
			chunk.blocks[x][z] = make([]int8, CHUNK_HEIGHT)
			worldPos := chunkAndLocalToWorldPos(chunkPos, Position{X: x, Z: z})
			ground := (noise.Eval2(float64(worldPos.X)*craziness, float64(worldPos.Z)*craziness) + 1) / 2 * CHUNK_HEIGHT
			if ground < 3 {
				fmt.Println("this only prints for larger chunk sizes somehow")
			}
			for y := 0; y < CHUNK_HEIGHT; y++ {
				if y == int(ground) {
					chunk.blocks[x][z][y] = GrassBlock
				} else if y < int(ground) && y >= int(ground-5) {
					chunk.blocks[x][z][y] = DirtBlock
				} else if y < int(ground) {
					chunk.blocks[x][z][y] = StoneBlock
				} else {
					chunk.blocks[x][z][y] = AirBlock
				}
			}
		}
	}
}
