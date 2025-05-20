package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// TYPES
type World struct {
	chunks map[Position]*Chunk
}

// TODO: Rewrite type
type Chunk struct {
	blocks    [][][]int8           // x z y
	boxes     [][][]rl.BoundingBox // temp
	chunkMesh *ChunkMesh
}

type ChunkMesh struct {
	Initialized bool

	VertexCount    int32
	Vertices       []float32
	TriangleCount  int32
	Indices        []uint16
	Colors         []uint8
	TexcoordsCount int32
	Texcoords      []float32

	Model rl.Model
}

type Position struct {
	X int
	Z int
}

// GENERATION

func (world *World) generateWorldMeshes() {
	for x := -WORLD_SIZE; x <= WORLD_SIZE; x++ {
		for z := -WORLD_SIZE; z <= WORLD_SIZE; z++ {
			world.chunks[Position{X: x, Z: z}].generateChunkMesh(Position{X: x, Z: z}, world)
		}
	}
}

// TODO: make separate func for boxes
// the world is taken as a parameter for some much-needed optimizations, but can be removed
func (chunk *Chunk) generateChunkMesh(chunkPos Position, world *World) {
	chunk.chunkMesh = &ChunkMesh{}
	chunk.boxes = make([][][]rl.BoundingBox, len(chunk.blocks))
	for x, plane := range chunk.blocks {
		chunk.boxes[x] = make([][]rl.BoundingBox, len(plane))
		for z, col := range plane {
			chunk.boxes[x][z] = make([]rl.BoundingBox, len(col))
			for y, block := range col {
				xBlockWorld := chunkPos.X*CHUNK_SIZE + x
				zBlockWorld := chunkPos.Z*CHUNK_SIZE + z
				if block == AirBlock || world.isBlockSurrounded(xBlockWorld, y, zBlockWorld) {
					continue
				}
				drawPos := rl.Vector3{X: float32(xBlockWorld), Y: float32(y), Z: float32(zBlockWorld)}
				chunk.chunkMesh.addBlock(drawPos, block)

				// temp
				chunk.boxes[x][z][y] = rl.NewBoundingBox(drawPos, rl.Vector3Add(drawPos, rl.Vector3{1, 1, 1}))
			}
		}
	}
	chunk.chunkMesh.buildChunkMesh()
}

func (chunkMesh *ChunkMesh) addBlock(position rl.Vector3, block int8) {
	// Initialization
	if !chunkMesh.Initialized {
		chunkMesh.Initialized = true
		chunkMesh.VertexCount = 0
		chunkMesh.Vertices = make([]float32, 0)
		chunkMesh.TriangleCount = 0
		chunkMesh.Indices = make([]uint16, 0)
		chunkMesh.Colors = make([]uint8, 0)
		chunkMesh.TexcoordsCount = 0
	}

	// Vertices
	chunkMesh.VertexCount += int32(len(cubeVertices) / 3)
	translatedVertices := make([]float32, len(cubeVertices))
	copy(translatedVertices, cubeVertices)
	for i := range translatedVertices {
		switch i % 3 {
		case 0:
			translatedVertices[i] += position.X
		case 1:
			translatedVertices[i] += position.Y
		case 2:
			translatedVertices[i] += position.Z
		}
	}
	chunkMesh.Vertices = append(chunkMesh.Vertices, translatedVertices...)

	// Indices
	chunkMesh.TriangleCount += 12
	startIndex := uint16(len(chunkMesh.Indices))
	for i := startIndex; i < startIndex+36; i++ {
		chunkMesh.Indices = append(chunkMesh.Indices, i)
	}

	// Colors
	for i := 0; i < 36; i++ {
		white := rl.White
		chunkMesh.Colors = append(chunkMesh.Colors, white.R, white.G, white.B, white.A)
	}

	// Textures
	coordinatesUV := make([]float32, len(cubeTexture))
	for i, v := range cubeTexture {
		// add the offset corresponding to the block's face's texture:

		// there are 36 UV points (3 per triangle; there are 12 triangles), so 72 float32s in the standard cube texture
		// each face has 12 float32s (72 total float32s / 6 faces = 12 float32s per face)
		// with i being in the range 1...71, i/12 is the index of each face (0...5)

		if i%2 == 0 { // even -> U
			coordinatesUV[i] = v + textureMap[block][i/12].X
		} else { // odd -> V
			coordinatesUV[i] = v + textureMap[block][i/12].Y
		}
	}
	chunkMesh.Texcoords = append(chunkMesh.Texcoords, coordinatesUV...)
}

func (chunkMesh *ChunkMesh) buildChunkMesh() {
	var mesh rl.Mesh
	mesh.VertexCount = chunkMesh.VertexCount
	mesh.Vertices = &chunkMesh.Vertices[0]
	mesh.TriangleCount = chunkMesh.TriangleCount
	mesh.Indices = &chunkMesh.Indices[0]
	mesh.Colors = &chunkMesh.Colors[0]
	mesh.Texcoords = &chunkMesh.Texcoords[0]

	rl.UploadMesh(&mesh, false)
	chunkMesh.Model = rl.LoadModelFromMesh(mesh)
	if chunkMesh.Model.Materials != nil {
		chunkMesh.Model.Materials.Maps.Texture = atlas
	}
}

// HELPER FUNCTIONS

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

	return chunk.blocks[localPos.X][localPos.Z][y]
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

var cubeVertices = []float32{
	// face 1
	0.0, 1.0, 0.0,
	0.0, 0.0, 0.0,
	0.0, 1.0, 1.0,
	0.0, 0.0, 1.0,
	0.0, 1.0, 1.0,
	0.0, 0.0, 0.0,

	// face 2
	0.0, 0.0, 0.0,
	0.0, 1.0, 0.0,
	1.0, 1.0, 0.0,
	1.0, 1.0, 0.0,
	1.0, 0.0, 0.0,
	0.0, 0.0, 0.0,

	// face 3
	1.0, 0.0, 0.0,
	1.0, 1.0, 0.0,
	1.0, 1.0, 1.0,
	1.0, 1.0, 1.0,
	1.0, 0.0, 1.0,
	1.0, 0.0, 0.0,

	// face 4
	0.0, 1.0, 1.0,
	0.0, 0.0, 1.0,
	1.0, 1.0, 1.0,
	1.0, 0.0, 1.0,
	1.0, 1.0, 1.0,
	0.0, 0.0, 1.0,

	// face 5
	1.0, 1.0, 0.0,
	0.0, 1.0, 0.0,
	0.0, 1.0, 1.0,
	0.0, 1.0, 1.0,
	1.0, 1.0, 1.0,
	1.0, 1.0, 0.0,

	// face 6
	0.0, 0.0, 0.0,
	1.0, 0.0, 0.0,
	0.0, 0.0, 1.0,
	1.0, 0.0, 1.0,
	0.0, 0.0, 1.0,
	1.0, 0.0, 0.0,
}
