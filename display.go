package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func displayCoords() {
	camPos := camera3D.Position
	camWorldPos := Position{X: int(camPos.X), Z: int(camPos.Z)}
	canvasDrawText(fmt.Sprintf("Coordinates:\nX: %d, Y: %d, Z: %d", camWorldPos.X, int(camPos.Y), camWorldPos.Z),
		5.0, 5.0, 20.0, rl.Black)
	chunkPos := camWorldPos.worldToChunkPos()
	canvasDrawText(fmt.Sprintf("Chunk positions:\nX: %d, Z: %d", chunkPos.X, chunkPos.Z),
		5.0, 20.0, 20.0, rl.Black)
}

func displayFPS() {
	canvasDrawText(fmt.Sprintf("FPS: %d", rl.GetFPS()),
		5.0, 35.0, 20.0, rl.Black)
}
