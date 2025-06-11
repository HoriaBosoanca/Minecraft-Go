package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

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

func (world *World) generateMeshes() {
	for _, chunk := range world.chunks {
		chunk.generateChunkMesh()
	}
}

// this uses world for some much-needed optimizations, but can be removed
func (chunk *Chunk) generateChunkMesh() {
	chunk.mesh = &ChunkMesh{}
	for x := range chunk.blocks {
		for z := range chunk.blocks[x] {
			for y, block := range chunk.blocks[x][z] {
				worldPos := chunkPos2AndLocalPos2ToWorldPos2(chunk.position, Position2{X: x, Z: z})
				if block.data == AirBlock || world.isBlockSurrounded(worldPos.X, y, worldPos.Z) {
					continue
				}
				drawPos := rl.Vector3{X: float32(worldPos.X), Y: float32(y), Z: float32(worldPos.Z)}
				chunk.mesh.addBlock(drawPos, block.data)
			}
		}
	}
	chunk.mesh.buildChunkMesh()
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
