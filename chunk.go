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
				pos := rl.Vector3{X: float32(xChunkPos*CHUNK_SIZE + x), Y: float32(y), Z: float32(zChunkPos*CHUNK_SIZE + z)}
				if block == AirBlock {
					continue
				}
				if chunk.isBlockSurrounded(x, y, z) {
					continue
				}
				switch block {
				case GrassBlock:
					drawCube(pos, rl.DarkGreen)
				case DirtBlock:
					drawCube(pos, rl.Brown)
				case StoneBlock:
					drawCube(pos, rl.Gray)
				default:
					continue
				}
			}
		}
	}
}

func (chunk *Chunk) isBlockSurrounded(xPos, yPos, zPos int) bool {
	if xPos == 0 || xPos == CHUNK_SIZE-1 || yPos == 0 || yPos == CHUNK_HEIGHT-1 || zPos == 0 || zPos == CHUNK_SIZE-1 {
		return false
	}
	if chunk.blocks[xPos-1][zPos][yPos] == AirBlock {
		return false
	}
	if chunk.blocks[xPos+1][zPos][yPos] == AirBlock {
		return false
	}
	if chunk.blocks[xPos][zPos][yPos-1] == AirBlock {
		return false
	}
	if chunk.blocks[xPos][zPos][yPos+1] == AirBlock {
		return false
	}
	if chunk.blocks[xPos][zPos-1][yPos] == AirBlock {
		return false
	}
	if chunk.blocks[xPos][zPos+1][yPos] == AirBlock {
		return false
	}
	return true
}
