package main

import (
	"github.com/ojrac/opensimplex-go"
	"time"
)

var noise = opensimplex.New(time.Now().Unix())
var chunks [][]Chunk
var chunkCount int = 3

func genWorld() {
	chunks = make([][]Chunk, chunkCount)
	for x := 0; x < chunkCount; x++ {
		chunks[x] = make([]Chunk, chunkCount)
		for z := 0; z < chunkCount; z++ {
			chunks[x][z].xPos = x
			chunks[x][z].zPos = z
			chunks[x][z].Generate(noise)
		}
	}
}

func renderWorld() {
	for _, chunkLine := range chunks {
		for _, chunk := range chunkLine {
			chunk.Render()
		}
	}
}
