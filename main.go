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
	rl.DisableCursor()

	genWorld()
}

func preDraw() {
	// in rcamera.h
	// #define CAMERA_MOVE_SPEED 21.6f // (initially 5.4f)
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
	renderWorld()
}

func draw2D() {
	canvasDrawText("My game", 5, 5, 20, rl.Black)
}
