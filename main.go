package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	Init()
	defer rl.CloseWindow()

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
	rl.InitWindow(screenX, screenY, "My game")
	loadModels()
	rl.SetTargetFPS(fps)
	rl.DisableCursor()

	//genWorld()
}

func preDraw() {
	rl.UpdateCamera(&camera3D, rl.CameraFirstPerson)
	rl.ClearBackground(rl.RayWhite)
	move()
}

func draw3D() {
	renderWorld(RENDER_DISTANCE)
}

func draw2D() {
	displayCoords()
	displayFPS()
}
