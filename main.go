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

var mesh ChunkMesh

func Init() {
	rl.InitWindow(screenX, screenY, "My game")
	rl.SetTargetFPS(fps)
	rl.DisableCursor()

	loadTextures()

	//generateWorldBlocks()
	//generateWorldMeshes()

	for i := float32(0); i < 100; i++ {
		mesh.addBlock(rl.Vector3{i, i, i}, rl.White)
	}
	mesh.build()
}

func preDraw() {
	rl.UpdateCamera(&camera3D, rl.CameraFirstPerson)
	rl.ClearBackground(rl.RayWhite)
	move()
}

func draw3D() {
	renderWorld(RENDER_DISTANCE)

	mesh.render()
}

func draw2D() {
	displayCoords()
	displayFPS()
}
