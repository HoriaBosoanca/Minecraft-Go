package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

// 2D

func canvasDrawText(text string, posX, posY float64, fs int32, color color.RGBA) {
	posX = posX * float64(screenX) / 100.0
	posY = posY * float64(screenY) / 100.0
	rl.DrawText(text, int32(posX), int32(posY), fs, color)
}

// 3D

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
	switch block { // if i % 2 == 0 -> U, else -> V
	case GrassBlock:
		for i, v := range cubeTexture {
			if i >= FACE_1_START && i <= FACE_4_END {
				if i%2 == 0 {
					coordinatesUV[i] = v + GRASS_SIDE_U
				} else {
					coordinatesUV[i] = v + GRASS_SIDE_V
				}
			}
			if i >= FACE_5_START && i <= FACE_5_END {
				if i%2 == 0 {
					coordinatesUV[i] = v + GRASS_TOP_U
				} else {
					coordinatesUV[i] = v + GRASS_TOP_V
				}
			}
			if i >= FACE_6_START && i <= FACE_6_END {
				if i%2 == 0 {
					coordinatesUV[i] = v + DIRT_U
				} else {
					coordinatesUV[i] = v + DIRT_V
				}
			}
		}
	default:
	}
	chunkMesh.Texcoords = append(chunkMesh.Texcoords, coordinatesUV...)
}

func (chunkMesh *ChunkMesh) build() {
	var mesh rl.Mesh
	mesh.VertexCount = chunkMesh.VertexCount
	mesh.Vertices = &chunkMesh.Vertices[0]
	mesh.TriangleCount = chunkMesh.TriangleCount
	mesh.Indices = &chunkMesh.Indices[0]
	mesh.Colors = &chunkMesh.Colors[0]
	mesh.Texcoords = &chunkMesh.Texcoords[0]

	rl.UploadMesh(&mesh, false)
	chunkMesh.Model = rl.LoadModelFromMesh(mesh)
	chunkMesh.Model.Materials.Maps.Texture = atlas
}

func (chunkMesh *ChunkMesh) render() {
	rl.DrawModel(chunkMesh.Model, rl.Vector3{}, 1.0, rl.White)
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
