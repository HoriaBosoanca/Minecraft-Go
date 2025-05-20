package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

func displayCoords(posX, posY float64) {
	camPos := camera3D.Position
	camWorldPos := Position{X: int(camPos.X), Z: int(camPos.Z)}
	canvasDrawText(fmt.Sprintf("Coordinates:\nX: %d, Y: %d, Z: %d", camWorldPos.X, int(camPos.Y), camWorldPos.Z), posX, posY, 20.0, rl.Black)
	chunkPos := worldToChunkPos(camWorldPos)
	canvasDrawText(fmt.Sprintf("Chunk positions:\nX: %d, Z: %d", chunkPos.X, chunkPos.Z), posX, posY+15.0, 20.0, rl.Black)
}

func displayFPS(posX, posY float64) {
	canvasDrawText(fmt.Sprintf("FPS: %d", rl.GetFPS()), posX, posY, 20.0, rl.Black)
}

func canvasDrawText(text string, posX, posY float64, fs int32, color color.RGBA) {
	posX = posX * float64(rl.GetScreenWidth()) / 100.0
	posY = posY * float64(rl.GetScreenHeight()) / 100.0
	rl.DrawText(text, int32(posX), int32(posY), fs, color)
}
