package main

import rl "github.com/gen2brain/raylib-go/raylib"

func (world *World) renderWorld(renderDistance int) {
	for chunkPos, chunk := range world.chunks {
		cameraWorldPos := Position{X: int(camera3D.Position.X), Z: int(camera3D.Position.Z)}
		cameraChunkPos := worldToChunkPos(cameraWorldPos)
		if chunkPos.X-cameraChunkPos.X <= renderDistance &&
			chunkPos.X-cameraChunkPos.X >= -renderDistance &&
			chunkPos.Z-cameraChunkPos.Z <= renderDistance &&
			chunkPos.Z-cameraChunkPos.Z >= -renderDistance {
			chunk.render()
		}
	}
	rl.DrawGrid(2*WORLD_SIZE, CHUNK_SIZE)
}

func (chunk *Chunk) render() {
	chunk.chunkMesh.render()
}
