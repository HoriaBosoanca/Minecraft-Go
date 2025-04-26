package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ojrac/opensimplex-go"
)

const chunkSize = 16
const chunkHeight = 32

type Chunk struct {
	xPos   int
	zPos   int
	blocks [][][]int8 // x z y
}

var craziness = 0.05

func (chunk *Chunk) Generate(noise opensimplex.Noise) {
	chunk.blocks = make([][][]int8, chunkSize)
	for x := 0; x < chunkSize; x++ {
		chunk.blocks[x] = make([][]int8, chunkSize)
		for z := 0; z < chunkSize; z++ {
			chunk.blocks[x][z] = make([]int8, chunkHeight)
			ground := (noise.Eval2(float64(chunk.xPos*chunkSize+x)*craziness, float64(chunk.zPos*chunkSize+z)*craziness) + 1) / 2 * chunkHeight
			for y := 0; y < chunkHeight; y++ {
				if y == int(ground) {
					chunk.blocks[x][z][y] = 2
				} else if y < int(ground) {
					chunk.blocks[x][z][y] = 1
				} else {
					chunk.blocks[x][z][y] = 0
				}
			}
		}
	}
}

func (chunk *Chunk) Render() {
	for x, plane := range chunk.blocks {
		for z, col := range plane {
			for y, block := range col {
				pos := rl.Vector3{X: float32(chunk.xPos*chunkSize + x), Y: float32(y), Z: float32(chunk.zPos*chunkSize + z)}
				if block == 1 {
					drawCube(pos, rl.Brown)
				}
				if block == 2 {
					drawCube(pos, rl.DarkGreen)
				}
			}
		}
	}
}
