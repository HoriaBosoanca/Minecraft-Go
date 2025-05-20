package main

import rl "github.com/gen2brain/raylib-go/raylib"

func handleInput() {
	deltaX := rl.Vector3Subtract(camera3D.Target, camera3D.Position)
	deltaX.Y = 0
	deltaX = rl.Vector3Normalize(deltaX)
	deltaX = rl.Vector3Scale(deltaX, MOVE_SPEED)
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
		camera3D.Position.Y += ASCEND_SPEED
		camera3D.Target.Y += ASCEND_SPEED
	}
	if rl.IsKeyDown(rl.KeyLeftShift) {
		camera3D.Position.Y -= ASCEND_SPEED
		camera3D.Target.Y -= ASCEND_SPEED
	}

	if rl.IsKeyDown(rl.KeyF11) {
		rl.ToggleFullscreen()
	}
	if rl.IsKeyDown(rl.KeyF10) {
		if rl.IsCursorHidden() {
			rl.EnableCursor()
		} else {
			rl.DisableCursor()
		}
	}
}
