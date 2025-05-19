package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	world := &World{}

	Init(world)
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		preDraw()
		rl.BeginMode3D(camera3D)
		draw3D(world)
		rl.EndMode3D()
		draw2D()
		rl.EndDrawing()
	}
}

func Init(world *World) {
	rl.InitWindow(X_SCREEN, Y_SCREEN, "My game")
	rl.SetTargetFPS(fps)
	rl.DisableCursor()

	loadTextures()
	world.generateWorldBlocks()
	world.generateWorldMeshes()
}

func preDraw() {
	rl.UpdateCamera(&camera3D, rl.CameraFirstPerson)
	rl.ClearBackground(rl.RayWhite)
	move()
}

func draw3D(world *World) {
	world.renderWorld(RENDER_DISTANCE)
}

func draw2D() {
	displayCoords()
	displayFPS()
}
