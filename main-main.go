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
	if SHOW_PLAYER_TARGET {
		drawPlayerTarget()
	}
}

func draw2D() {
	canvasDrawText("F10 = show cursor, F11 = fullscreen, F9 = show target block", 5.0, 5.0, 20.0, rl.Black)
	displayCoords(5.0, 20.0)
	displayFPS(5.0, 50.0)
	displayMemUsage(5.0, 65.0)
	drawCrosshair()
	if SHOW_PLAYER_TARGET {
		displayTargetCoordinates(5.0, 80.0)
	}
}
