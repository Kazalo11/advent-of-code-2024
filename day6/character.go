package main

type Character struct {
	direction Direction
	location  [2]int
}

func NewCharacter(location [2]int) *Character {
	return &Character{
		direction: 0,
		location:  location,
	}
}

func (c *Character) TurnDirection() {
	c.direction = (c.direction + 1) % 4
}

func (c *Character) GetLocation() [2]int {
	return c.location
}

func (c *Character) GetDirection() Direction {
	return c.direction
}

func (c *Character) Up() {
	c.location[0] = c.location[0] - 1
}

func (c *Character) Down() {
	c.location[0] = c.location[0] + 1
}

func (c *Character) Left() {
	c.location[1] = c.location[1] - 1
}

func (c *Character) Right() {
	c.location[1] = c.location[1] + 1
}

func (c *Character) LookAhead(m Matrix) string {
	curr := c.GetLocation()
	row := curr[0]
	col := curr[1]
	switch c.GetDirection() {
	case FORWARD:
		if row == 0 {

		}
		return m.data[row-1][col]
	case BACKWARD:
		guard.Down()
	case LEFT:
		guard.Left()
	case RIGHT:
		guard.Right()

	}
}
