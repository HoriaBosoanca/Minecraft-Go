package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func displayCoords() {
	pos := camera3D.Position
	canvasDrawText(fmt.Sprintf("Coordinates:\nX: %f, Y: %f, Z: %f", pos.X, pos.Y, pos.Z),
		5.0, 5.0, 20.0, rl.Black)
	chunkPos := worldToChunkPos(int(pos.X), int(pos.Z))
	canvasDrawText(fmt.Sprintf("Chunk positions:\nX: %d, Z: %d", chunkPos.X, chunkPos.Z),
		5.0, 20.0, 20.0, rl.Black)
}

func displayFPS() {
	canvasDrawText(fmt.Sprintf("FPS: %d", rl.GetFPS()),
		5.0, 35.0, 20.0, rl.Black)
}
