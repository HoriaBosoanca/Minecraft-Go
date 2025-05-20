package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (world *World) getRayTarget(ray rl.Ray) rl.BoundingBox {
	for _, chunk := range world.chunks {
		for _, plane := range chunk.boxes {
			for _, col := range plane {
				for _, block := range col {
					rayCollision := rl.GetRayCollisionBox(ray, block)
					if rayCollision.Hit {
						return block
					}
				}
			}
		}
	}
	return rl.BoundingBox{}
}
