package main

import "sort"

func NextObstacleUp(curr [2]int, obstacles [][2]int) ([2]int, bool) {

	sort.Slice(obstacles, func(i, j int) bool {
		return obstacles[i][0] > obstacles[j][0]
	})

	for _, item := range obstacles {
		if item[1] == curr[1] && item[0] < curr[0] {
			return item, false
		}
	}

	return [2]int{-1, curr[1]}, true

}

func NextObstacleDown(curr [2]int, obstacles [][2]int) ([2]int, bool) {

	sort.Slice(obstacles, func(i, j int) bool {
		return obstacles[i][0] < obstacles[j][0]
	})

	for _, item := range obstacles {
		if item[1] == curr[1] && item[0] > curr[0] {
			return item, false
		}
	}

	return [2]int{130, curr[1]}, true

}

func NextObstacleLeft(curr [2]int, obstacles [][2]int) ([2]int, bool) {

	sort.Slice(obstacles, func(i, j int) bool {
		return obstacles[i][1] > obstacles[j][1]
	})

	for _, item := range obstacles {
		if item[0] == curr[0] && item[1] < curr[1] {
			return item, false
		}
	}

	return [2]int{curr[0], -1}, true

}

func NextObstacleRight(curr [2]int, obstacles [][2]int) ([2]int, bool) {

	sort.Slice(obstacles, func(i, j int) bool {
		return obstacles[i][1] < obstacles[j][1]
	})

	for _, item := range obstacles {
		if item[0] == curr[0] && item[1] > curr[1] {
			return item, false
		}
	}

	return [2]int{curr[0], 130}, true

}

func GetDistinctItems(results [][2]int) int {
	ans := map[[2]int]bool{}
	for _, item := range results {
		ans[item] = true
	}
	return len(ans)

}
