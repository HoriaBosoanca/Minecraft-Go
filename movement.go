package main

import rl "github.com/gen2brain/raylib-go/raylib"

const moveSpeed float32 = 1
const ascendSpeed float32 = 1

func move() {
	deltaX := rl.Vector3Subtract(camera3D.Target, camera3D.Position)
	deltaX.Y = 0
	deltaX = rl.Vector3Normalize(deltaX)
	deltaX = rl.Vector3Scale(deltaX, moveSpeed)
	deltaZ := rl.NewVector3(deltaX.Z, 0, -deltaX.X)

	if rl.IsKeyDown(rl.KeyW) {
		camera3D.Position = rl.Vector3Add(camera3D.Position, deltaX)
		camera3D.Target = rl.Vector3Add(camera3D.Target, deltaX)
	}
	if rl.IsKeyDown(rl.KeyA) {
		camera3D.Position = rl.Vector3Add(camera3D.Position, deltaZ)
		camera3D.Target = rl.Vector3Add(camera3D.Target, deltaZ)
	}
	if rl.IsKeyDown(rl.KeyS) {
		camera3D.Position = rl.Vector3Subtract(camera3D.Position, deltaX)
		camera3D.Target = rl.Vector3Subtract(camera3D.Target, deltaX)
	}
	if rl.IsKeyDown(rl.KeyD) {
		camera3D.Position = rl.Vector3Subtract(camera3D.Position, deltaZ)
		camera3D.Target = rl.Vector3Subtract(camera3D.Target, deltaZ)
	}

	if rl.IsKeyDown(rl.KeySpace) {
		camera3D.Position.Y += ascendSpeed
		camera3D.Target.Y += ascendSpeed
	}
	if rl.IsKeyDown(rl.KeyLeftShift) {
		camera3D.Position.Y -= ascendSpeed
		camera3D.Target.Y -= ascendSpeed
	}
}
