package main

type Structure struct {
	blocks [][][]int8
}

var (
	tree = Structure{
		blocks: [][][]int8{
			{
				{OakLogBlock, OakLogBlock, OakLogBlock, OakLogBlock, OakLogBlock, OakLeafBlock},
			},
		},
	}
)

func (chunk *Chunk) addStructure(structure Structure, chunkPos Position3) {
	for x := range structure.blocks {
		for z := range structure.blocks[x] {
			for y, block := range structure.blocks[x][z] {
				chunk.addBlock(block, Position3{X: chunkPos.X + x, Y: chunkPos.Y + y, Z: chunkPos.Z + z})
			}
		}
	}
}
