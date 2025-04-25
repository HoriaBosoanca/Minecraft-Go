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
	rl.SetTargetFPS(fps)
}

func preDraw() {
	rl.UpdateCamera(&camera3D, rl.CameraOrbital)
	rl.ClearBackground(rl.RayWhite)
}

func draw3D() {
	rl.DrawCube(Origin(), 1, 1, 1, rl.Yellow)
	rl.DrawCubeWires(Origin(), 1, 1, 1, rl.Black)
	rl.DrawGrid(100, 1)
}

func draw2D() {
	canvasDrawText("My game", 5, 5, 20, rl.Black)
}
