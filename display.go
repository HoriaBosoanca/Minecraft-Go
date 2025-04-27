package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func displayCoords() {
	pos := camera3D.Position
	canvasDrawText(fmt.Sprintf("Coordinates:\nX: %f, Y: %f, Z: %f", pos.X, pos.Y, pos.Z),
		5.0, 5.0, 20.0, rl.Black)
	canvasDrawText(fmt.Sprintf("Chunk positions:\nX: %d, Z: %d", int(pos.X/16), int(pos.Z/16)),
		5.0, 20.0, 20.0, rl.Black)
}

func displayFPS() {
	canvasDrawText(fmt.Sprintf("FPS: %d", rl.GetFPS()),
		5.0, 35.0, 20.0, rl.Black)
}
