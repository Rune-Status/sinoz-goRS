package net

var BitMasks = getBitMasks()

func getBitMasks() []int {
	masks := make([]int, 32)

	for i := 0; i < len(masks); i++ {
		masks[i] = (1 << uint(i)) - 1
	}

	return masks
}