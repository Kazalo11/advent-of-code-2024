package main

func part1() {
	matrix := NewMatrix("day6/day6.txt")
	guardLocation := matrix.FindGuard()
	obstacles := matrix.FindObstacles()
}
