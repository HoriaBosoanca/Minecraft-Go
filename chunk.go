package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ojrac/opensimplex-go"
)

const chunkSize = 16
const chunkHeight = 127

type Chunk struct {
	xPos   int
	zPos   int
	blocks [][][]int8 // x z y
}

var crazyness float64 = 0.01

func (chunk *Chunk) Generate(noise opensimplex.Noise) {
	chunk.blocks = make([][][]int8, chunkSize)
	for x := 0; x < chunkSize; x++ {
		chunk.blocks[x] = make([][]int8, chunkSize)
		for z := 0; z < chunkSize; z++ {
			chunk.blocks[x][z] = make([]int8, chunkHeight)
			for y := 0; y < chunkHeight; y++ {
				ground := (noise.Eval2(float64(chunk.xPos*chunkSize+x)*crazyness, float64(chunk.zPos+chunkSize+z)*crazyness) + 1) / 2 * 127.0
				if y < int(ground) {
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
				if block == 1 {
					rl.DrawCube(
						rl.Vector3{X: float32(chunk.xPos*chunkSize + x), Y: float32(y), Z: float32(chunk.zPos*chunkSize + z)},
						1, 1, 1, rl.DarkGreen,
					)
					rl.DrawCubeWires(
						rl.Vector3{X: float32(chunk.xPos*chunkSize + x), Y: float32(y), Z: float32(chunk.zPos*chunkSize + z)},
						1, 1, 1, rl.Black,
					)
				}
			}
		}
	}
}
