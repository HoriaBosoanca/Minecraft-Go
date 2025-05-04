package main

import rl "github.com/gen2brain/raylib-go/raylib"

var dirtTexture rl.Texture2D
var grassTexture rl.Texture2D

func loadTextures() {
	dirtTexture = rl.LoadTexture("assets/dirt.png")
	grassTexture = rl.LoadTexture("assets/grass.png")
}
