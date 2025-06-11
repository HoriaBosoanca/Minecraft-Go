package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
	"runtime"
)

func displayCoords(posX, posY float64) {
	camPos := camera3D.Position
	camWorldPos := Position2{X: int(camPos.X), Z: int(camPos.Z)}
	canvasDrawText(fmt.Sprintf("Coordinates:\nX: %d, Y: %d, Z: %d", camWorldPos.X, int(camPos.Y), camWorldPos.Z), posX, posY, 20.0, rl.Black)
	chunkPos := worldPos2ToChunkPos2(camWorldPos)
	canvasDrawText(fmt.Sprintf("Chunk positions:\nX: %d, Z: %d", chunkPos.X, chunkPos.Z), posX, posY+15.0, 20.0, rl.Black)
}

func displayFPS(posX, posY float64) {
	canvasDrawText(fmt.Sprintf("FPS: %d", rl.GetFPS()), posX, posY, 20.0, rl.Black)
}

func displayMemUsage(posX, posY float64) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	canvasDrawText(fmt.Sprintf("Memory: %d MB", mem.Alloc/1024/1024), posX, posY, 20.0, rl.Black)
}

func displayTargetCoordinates(posX, posY float64) {
	block, blockPos, _ := world.getClosestBlockHit(getPlayerRay(), MAX_CONTINUOUS_CHUNK_TARGET_SEARCH)
	if block != nil {
		canvasDrawText(fmt.Sprintf("Target:\nX: %d, Y: %d, Z: %d", blockPos.X, blockPos.Y, blockPos.Z), posX, posY, 20.0, rl.Black)
	} else {
		canvasDrawText("Target: nil", posX, posY, 20.0, rl.Black)
	}
}

func canvasDrawText(text string, posX, posY float64, fs int32, color color.RGBA) {
	posX = posX * float64(rl.GetScreenWidth()) / 100.0
	posY = posY * float64(rl.GetScreenHeight()) / 100.0
	rl.DrawText(text, int32(posX), int32(posY), fs, color)
}
