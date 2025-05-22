package main

import rl "github.com/gen2brain/raylib-go/raylib"

var world = &World{}

func main() {
	Init()
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		preDraw()
		rl.BeginMode3D(camera3D)
		draw3D()
		rl.EndMode3D()
		draw2D()
		rl.EndDrawing()
	}
}

func Init() {
	rl.InitWindow(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), "Raygocraft")
	rl.ToggleFullscreen()
	rl.SetTargetFPS(fps)
	rl.DisableCursor()

	loadTextures()
	world.memoryInit()
	world.colliderInit()
	world.generateBlocks()
	world.generateMeshes()
}

func preDraw() {
	rl.UpdateCamera(&camera3D, rl.CameraFirstPerson)
	rl.ClearBackground(rl.RayWhite)
	handleInput()
}

func draw3D() {
	world.renderWorld(RENDER_DISTANCE)
	drawPlayerTarget()
}

func draw2D() {
	displayCoords(5.0, 20.0)
	displayFPS(5.0, 50.0)
	canvasDrawText("Press F10 to enable cursor and F11 to disable fullscreen", 5.0, 5.0, 20.0, rl.Black)
}
