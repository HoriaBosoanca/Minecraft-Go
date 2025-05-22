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

func (world *World) generateBlocks() {
	for chunkPos, chunk := range world.chunks {
		chunk.generateBlockData(chunkPos)
	}
}

func (chunk *Chunk) generateBlockData(chunkPos Position) {
	for x := range chunk.blocks {
		for z := range chunk.blocks[x] {
			worldPos := chunkAndLocalToWorldPos(chunkPos, Position{X: x, Z: z})
			ground := (noise.Eval2(float64(worldPos.X)*craziness, float64(worldPos.Z)*craziness) + 1) / 2 * CHUNK_HEIGHT
			if ground < 3 {
				fmt.Println("this only prints for larger chunk sizes somehow")
			}
			for y := 0; y < CHUNK_HEIGHT; y++ {
				if y == int(ground) {
					chunk.blocks[x][z][y].data = GrassBlock
				} else if y < int(ground) && y >= int(ground-5) {
					chunk.blocks[x][z][y].data = DirtBlock
				} else if y < int(ground) {
					chunk.blocks[x][z][y].data = StoneBlock
				} else {
					chunk.blocks[x][z][y].data = AirBlock
				}
			}
		}
	}
}
