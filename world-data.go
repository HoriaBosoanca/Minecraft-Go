package main

import (
	"github.com/ojrac/opensimplex-go"
	"math/rand"
	"time"
)

var seed = rand.New(rand.NewSource(time.Now().Unix()))
var terrainNoise = opensimplex.New(seed.Int63())
var treeNoise = opensimplex.New(seed.Int63())

const (
	terrainCraziness = 0.03
	waterLevel       = 9
	treeCraziness    = 0.5
	treeAmount       = 0.1 // 0 to 1
)

const (
	AirBlock int8 = iota
	GrassBlock
	DirtBlock
	StoneBlock
	OakLeafBlock
	OakLogBlock
	WaterBlock
)

func (world *World) generateBlocks() {
	for chunkPos, chunk := range world.chunks {
		chunk.generateBlockData(chunkPos)
	}
}

func (chunk *Chunk) generateBlockData(chunkPos Position2) {
	chunk.generateTerrain(chunkPos)
	chunk.addWater()
	//chunk.addTrees(chunkPos)
}

func (chunk *Chunk) addBlock(block int8, chunkPos Position3) {
	if chunkPos.X >= 0 && chunkPos.X < CHUNK_SIZE && chunkPos.Z >= 0 && chunkPos.Z < CHUNK_SIZE && chunkPos.Y >= 0 && chunkPos.Y < CHUNK_HEIGHT {
		chunk.blocks[chunkPos.X][chunkPos.Z][chunkPos.Y].data = block
	}
}

func getGroundLevel(worldPos Position2) (groundHeight int) {
	return int((terrainNoise.Eval2(float64(worldPos.X)*terrainCraziness, float64(worldPos.Z)*terrainCraziness) + 1) / 2 * CHUNK_HEIGHT)
}

func positionHasTree(worldPos Position2) bool {
	return (treeNoise.Eval2(float64(worldPos.X)*treeCraziness, float64(worldPos.Z)*treeCraziness)+1)/2 < treeAmount
}

func (chunk *Chunk) generateTerrain(chunkPos Position2) {
	for x := range chunk.blocks {
		for z := range chunk.blocks[x] {
			ground := getGroundLevel(chunkPos2AndLocalPos2ToWorldPos2(chunkPos, Position2{X: x, Z: z}))
			for y := range chunk.blocks[x][z] {
				if y == ground {
					chunk.addBlock(GrassBlock, Position3{X: x, Y: y, Z: z})
				} else if y < ground && y >= ground-5 {
					chunk.addBlock(DirtBlock, Position3{X: x, Y: y, Z: z})
				} else if y < ground {
					chunk.addBlock(StoneBlock, Position3{X: x, Y: y, Z: z})
				} else {
					chunk.addBlock(AirBlock, Position3{X: x, Y: y, Z: z})
				}
			}
		}
	}
}

func (chunk *Chunk) addWater() {
	for x := range chunk.blocks {
		for z := range chunk.blocks[x] {
			for y := range chunk.blocks[x][z] {
				if y < waterLevel && chunk.blocks[x][z][y].data == AirBlock {
					chunk.addBlock(WaterBlock, Position3{X: x, Y: y, Z: z})
				}
			}
		}
	}
}

func (chunk *Chunk) addTrees(chunkPos Position2) {
	for x := range chunk.blocks {
		for z := range chunk.blocks[x] {
			worldPos := chunkPos2AndLocalPos2ToWorldPos2(chunkPos, Position2{X: x, Z: z})
			ground := getGroundLevel(worldPos)
			if positionHasTree(worldPos) {
				chunk.addStructure(tree, Position3{X: x, Y: ground, Z: z})
			}
		}
	}
}
