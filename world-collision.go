package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (world *World) colliderInit() {
	for chunkPos, chunk := range world.chunks {
		chunkMin := chunkAndLocalToWorldPos(Position{X: chunkPos.X, Z: chunkPos.Z}, Position{X: 0, Z: 0})
		chunkMax := chunkAndLocalToWorldPos(Position{X: chunkPos.X, Z: chunkPos.Z}, Position{X: CHUNK_SIZE, Z: CHUNK_SIZE})
		chunk.collider = rl.NewBoundingBox(rl.Vector3{X: float32(chunkMin.X), Y: 0, Z: float32(chunkMin.Z)}, rl.Vector3{X: float32(chunkMax.X), Y: CHUNK_HEIGHT, Z: float32(chunkMax.Z)})
		for x := range chunk.blocks {
			for z := range chunk.blocks[x] {
				for y, block := range chunk.blocks[x][z] {
					worldPos := chunkAndLocalToWorldPos(chunkPos, Position{X: x, Z: z})
					worldPosFloat32 := rl.Vector3{X: float32(worldPos.X), Y: float32(y), Z: float32(worldPos.Z)}
					block.collider = rl.NewBoundingBox(worldPosFloat32, rl.Vector3Add(worldPosFloat32, rl.Vector3{X: 1, Y: 1, Z: 1}))
				}
			}
		}
	}
}

func (world *World) getClosestBlockHit(ray rl.Ray) *Block {
	var closest *Block = nil
	for chunkPos, chunk := range world.chunks {
		rayCol := rl.GetRayCollisionBox(ray, chunk.collider)
		if rayCol.Hit && (closest == nil ||
			rl.Vector3Distance(positionToVector3(chunkPos), positionToVector3(worldToChunkPos(vector3ToPosition(ray.Position)))) <
				rl.Vector3Distance(positionToVector3(worldToChunkPos(vector3ToPosition(closest.collider.Min))), positionToVector3(worldToChunkPos(vector3ToPosition(ray.Position))))) {
			for x := range chunk.blocks {
				for z := range chunk.blocks[x] {
					for _, block := range chunk.blocks[x][z] {
						rayCollision := rl.GetRayCollisionBox(ray, block.collider)
						if closest == nil || (rayCollision.Hit && block.data != AirBlock &&
							rl.Vector3Distance(block.collider.Min, ray.Position) < rl.Vector3Distance(closest.collider.Min, ray.Position)) {
							closest = block
						}
					}
				}
			}
		}
	}
	return closest
}
