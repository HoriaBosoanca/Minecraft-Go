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
