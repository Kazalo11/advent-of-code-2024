package main

type Character struct {
	direction Direction
	location  [2]int
}

func (c *Character) NewCharacter(location [2]int) *Character {
	return &Character{
		direction: 0,
		location:  location,
	}
}

func (c *Character) TurnDirection() {
	c.direction = (c.direction + 1) % 4
}
