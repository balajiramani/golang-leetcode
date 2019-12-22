package main

func minTimeToVisitAllPoints(points [][]int) int {
	var sum = 0
	for i := 1; i < len(points); i++ {
		var xdiff= points[i][0] - points[i-1][0]
		var ydiff= points[i][1] - points[i-1][1]
		sum += max(abs(xdiff), abs(ydiff))
	}
	return sum
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
