package main

func part1() {
	matrix := NewMatrix("day6/day6.txt")
	guardLocation := matrix.FindGuard()
	guard := NewCharacter(guardLocation)
	for {
		switch guard.GetDirection() {
		case FORWARD:
			guard.Up()
		case BACKWARD:
			guard.Down()
		case LEFT:
			guard.Left()
		case RIGHT:
			guard.Right()

		}
	}
}
