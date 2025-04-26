package main

import (
	"github.com/ojrac/opensimplex-go"
	"time"
)

var noise = opensimplex.New(time.Now().Unix())
var chunks [][]Chunk
var sqrtChunkCount = 9

func genWorld() {
	chunks = make([][]Chunk, sqrtChunkCount)
	for x := 0; x < sqrtChunkCount; x++ {
		chunks[x] = make([]Chunk, sqrtChunkCount)
		for z := 0; z < sqrtChunkCount; z++ {
			chunks[x][z].xPos = x
			chunks[x][z].zPos = z
			chunks[x][z].Generate(noise)
		}
	}
}

func renderWorld(renderDistance int) {
	for _, chunkLine := range chunks {
		for _, chunk := range chunkLine {
			if chunk.xPos-int(camera3D.Position.X)/16 <= renderDistance &&
				chunk.xPos-int(camera3D.Position.X)/16 >= -renderDistance &&
				chunk.zPos-int(camera3D.Position.Z)/16 <= renderDistance &&
				chunk.zPos-int(camera3D.Position.Z)/16 >= -renderDistance {
				chunk.Render()
			}
		}
	}
}
