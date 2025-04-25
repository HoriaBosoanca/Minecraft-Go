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

var chunk Chunk

func Init() {
	rl.InitWindow(screenX, screenY, "My game")
	rl.SetTargetFPS(fps)
	rl.DisableCursor()

	chunk.xPos = 0
	chunk.zPos = 0
	chunk.Generate(noise)
}

func preDraw() {
	rl.UpdateCamera(&camera3D, rl.CameraFirstPerson)
	rl.ClearBackground(rl.RayWhite)
	if rl.IsKeyDown(rl.KeySpace) {
		camera3D.Position.Y += ascendSpeed
		camera3D.Target.Y += ascendSpeed
	}
	if rl.IsKeyDown(rl.KeyLeftShift) {
		camera3D.Position.Y -= ascendSpeed
		camera3D.Target.Y -= ascendSpeed
	}
}

func draw3D() {
	rl.DrawCube(Origin(), 1, 1, 1, rl.Yellow)
	rl.DrawCubeWires(Origin(), 1, 1, 1, rl.Black)
	rl.DrawGrid(100, 1)

	chunk.Render()
}

func draw2D() {
	canvasDrawText("My game", 5, 5, 20, rl.Black)
}
