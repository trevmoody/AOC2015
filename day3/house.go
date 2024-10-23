package main

type House struct {
	x int
	y int
}

func (h House) MoveRight() House {
	return House{h.x + 1, h.y}
}

func (h House) MoveLeft() House {
	return House{h.x - 1, h.y}
}

func (h House) MoveUp() House {
	return House{h.x, h.y + 1}
}

func (h House) MoveDown() House {
	return House{h.x, h.y - 1}
}
