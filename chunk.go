package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ojrac/opensimplex-go"
)

const CHUNK_SIZE = 16
const CHUNK_HEIGHT = 32

type Chunk struct {
	blocks [][][]int8 // x z y
}

const craziness = 0.05

func (chunk *Chunk) Generate(noise opensimplex.Noise, xChunkPos, zChunkPos int) {
	chunk.blocks = make([][][]int8, CHUNK_SIZE)
	for x := 0; x < CHUNK_SIZE; x++ {
		chunk.blocks[x] = make([][]int8, CHUNK_SIZE)
		for z := 0; z < CHUNK_SIZE; z++ {
			chunk.blocks[x][z] = make([]int8, CHUNK_HEIGHT)
			ground := (noise.Eval2(float64(xChunkPos*CHUNK_SIZE+x)*craziness, float64(zChunkPos*CHUNK_SIZE+z)*craziness) + 1) / 2 * CHUNK_HEIGHT
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

func (chunk *Chunk) Render(xChunkPos, zChunkPos int) {
	for x, plane := range chunk.blocks {
		for z, col := range plane {
			for y, block := range col {
				xWorld := xChunkPos*CHUNK_SIZE + x
				zWorld := zChunkPos*CHUNK_SIZE + z
				drawPos := rl.Vector3{X: float32(xWorld), Y: float32(y), Z: float32(zWorld)}
				if block == AirBlock {
					continue
				}
				if chunk.isBlockSurrounded(xWorld, y, zWorld) {
					continue
				}
				switch block {
				case GrassBlock:
					drawCube(drawPos, rl.DarkGreen)
				case DirtBlock:
					drawCube(drawPos, rl.Brown)
				case StoneBlock:
					drawCube(drawPos, rl.Gray)
				default:
					continue
				}
			}
		}
	}
}

func (chunk *Chunk) isBlockSurrounded(x, y, z int) bool {
	if worldGetBlock(x-1, y, z) == AirBlock {
		return false
	}
	if worldGetBlock(x+1, y, z) == AirBlock {
		return false
	}
	if worldGetBlock(x, y-1, z) == AirBlock {
		return false
	}
	if worldGetBlock(x, y+1, z) == AirBlock {
		return false
	}
	if worldGetBlock(x, y, z-1) == AirBlock {
		return false
	}
	if worldGetBlock(x, y, z+1) == AirBlock {
		return false
	}
	return true
}
