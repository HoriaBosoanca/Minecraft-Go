package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

func canvasDrawText(text string, posX, posY float64, fs int32, color color.RGBA) {
	posX = posX * float64(screenX) / 100.0
	posY = posY * float64(screenY) / 100.0
	rl.DrawText(text, int32(posX), int32(posY), fs, color)
}
