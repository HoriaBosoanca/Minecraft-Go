package main

type Position struct {
	X int
	Z int
}

func (worldPos Position) worldToChunkPos() Position {
	xChunk := worldPos.X / CHUNK_SIZE
	if worldPos.X < 0 && worldPos.X%CHUNK_SIZE != 0 {
		xChunk--
	}
	zChunk := worldPos.Z / CHUNK_SIZE
	if worldPos.Z < 0 && worldPos.Z%CHUNK_SIZE != 0 {
		zChunk--
	}
	return Position{X: xChunk, Z: zChunk}
}

// local = coordinates of block within chunk
func (worldPos Position) worldToLocalPos() Position {
	chunkCoords := worldPos.worldToChunkPos()
	localX := worldPos.X - chunkCoords.X*CHUNK_SIZE
	localZ := worldPos.Z - chunkCoords.Z*CHUNK_SIZE
	return Position{X: localX, Z: localZ}
}

func chunkAndLocalToWorldPos(chunkPos, localPos Position) Position {
	return Position{X: chunkPos.X*CHUNK_SIZE + localPos.X, Z: chunkPos.Z*CHUNK_SIZE + localPos.Z}
}
