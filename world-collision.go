package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (world *World) colliderInit() {
	for chunkPos, chunk := range world.chunks {
		chunkMin := chunkAndLocalToWorldPos(Position{X: chunkPos.X, Z: chunkPos.Z}, Position{X: 0, Z: 0})
		chunkMax := chunkAndLocalToWorldPos(Position{X: chunkPos.X, Z: chunkPos.Z}, Position{X: CHUNK_SIZE, Z: CHUNK_SIZE})
		chunk.collider = rl.NewBoundingBox(rl.Vector3{X: float32(chunkMin.X), Y: 0, Z: float32(chunkMin.Z)}, rl.Vector3{X: float32(chunkMax.X), Y: CHUNK_HEIGHT, Z: float32(chunkMax.Z)})
	}
}

func (world *World) getClosestBlockHit(ray rl.Ray, maxDistance float32) (*Block, *Position3, *Chunk) {
	var closestBlock *Block = nil
	var closestBlockPos *Position3 = nil
	var closestChunk *Chunk = nil
	for chunkPos, chunk := range world.chunks {
		chunkCol := rl.GetRayCollisionBox(ray, chunk.collider)
		if chunkCol.Hit && rl.Vector3Distance(positionToVector3(chunkPos), positionToVector3(worldToChunkPos(vector3ToPosition(ray.Position)))) < maxDistance {
			for x := range chunk.blocks {
				for z := range chunk.blocks[x] {
					for y, block := range chunk.blocks[x][z] {
						worldPosVec3 := positionToVector3(chunkAndLocalToWorldPos(chunkPos, Position{X: x, Z: z}))
						worldPosVec3.Y = float32(y)
						worldPos := Position3{X: int(worldPosVec3.X), Y: int(worldPosVec3.Y), Z: int(worldPosVec3.Z)}
						blockCol := rl.GetRayCollisionBox(ray, rl.NewBoundingBox(worldPosVec3, rl.Vector3Add(worldPosVec3, rl.Vector3{X: 1, Y: 1, Z: 1})))
						if closestBlock == nil || closestBlockPos == nil || (blockCol.Hit && block.data != AirBlock &&
							rl.Vector3Distance(worldPosVec3, ray.Position) < rl.Vector3Distance(position3ToVector3(*closestBlockPos), ray.Position)) {
							closestBlock = block
							closestBlockPos = &worldPos
							closestChunk = chunk
						}
					}
				}
			}
		}
	}
	return closestBlock, closestBlockPos, closestChunk
}
