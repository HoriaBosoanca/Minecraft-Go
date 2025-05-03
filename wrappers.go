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

func drawCube(pos rl.Vector3, color color.RGBA) {
	rl.DrawModel(cubeModel, pos, 1.0, color)
	rl.DrawCubeWires(pos, 1, 1, 1, rl.Black)
}

var cubeModel rl.Model

func loadModels() {
	cubeModel = rl.LoadModelFromMesh(rl.GenMeshCube(1, 1, 1))
}

type ChunkMesh struct {
	Initialized bool

	VertexCount   int32
	Vertices      []float32
	TriangleCount int32
	Indices       []uint16
	Colors        []uint8
}

func (chunkMesh *ChunkMesh) addBlock(position rl.Vector3, color rl.Color) {
	if !chunkMesh.Initialized {
		chunkMesh.Initialized = true
		chunkMesh.VertexCount = 0
		chunkMesh.Vertices = make([]float32, 0)
		chunkMesh.TriangleCount = 0
		chunkMesh.Indices = make([]uint16, 0)
		chunkMesh.Colors = make([]uint8, 0)
	}

	chunkMesh.VertexCount += 36
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
	chunkMesh.TriangleCount += 12
	startIndex := uint16(len(chunkMesh.Indices))
	for i := startIndex; i < startIndex+36; i++ {
		chunkMesh.Indices = append(chunkMesh.Indices, i)
	}
	for i := 0; i < 36; i++ {
		chunkMesh.Colors = append(chunkMesh.Colors, color.R, color.G, color.B, color.A)
	}
}

func (chunkMesh *ChunkMesh) render() {
	var mesh rl.Mesh
	mesh.VertexCount = chunkMesh.VertexCount
	mesh.Vertices = &chunkMesh.Vertices[0]
	mesh.TriangleCount = chunkMesh.TriangleCount
	mesh.Indices = &chunkMesh.Indices[0]
	mesh.Colors = &chunkMesh.Colors[0]

	rl.UploadMesh(&mesh, false)
	rl.DrawModel(rl.LoadModelFromMesh(mesh), rl.Vector3{}, 1.0, rl.White)
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
