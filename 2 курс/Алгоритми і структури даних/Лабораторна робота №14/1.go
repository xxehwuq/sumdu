package main

import (
	"fmt"
)

const i, j = 8, 8

var prizes = [i][j]int{
	{2, -1, 3, 3, -2, -3, 4, 2},
	{-1, 2, -2, -1, 4, 2, -3, 2},
	{2, -1, 5, 3, -1, 2, -1, 3},
	{1, -2, 5, 3, 3, -2, -4, 6},
	{-2, 4, -2, 3, -5, 1, -3, -4},
	{-3, -2, -1, -5, 3, 1, -3, -2},
	{-2, -3, -2, 3, -2, 3, -1, -5},
	{3, -1, -2, -3, -2, -1, 4, 5},
}

var dp [i][j]int

func calcReversePath() {
	dp[i-1][j-1] = prizes[i-1][j-1]

	for i := i - 2; i >= 0; i-- {
		dp[i][j-1] = prizes[i][j-1] + dp[i+1][j-1]
	}

	for j := j - 2; j >= 0; j-- {
		dp[i-1][j] = prizes[i-1][j] + dp[i-1][j+1]
	}

	for i := i - 2; i >= 0; i-- {
		for j := j - 2; j >= 0; j-- {
			dp[i][j] = prizes[i][j] + getMaxNum(dp[i+1][j], dp[i+1][j+1], dp[i][j+1])
		}
	}
}

func buildPath() (path [][2]int, totalPrize int) {
	x, y := 0, 0
	totalPrize = dp[0][0]
	path = append(path, [2]int{x + 1, y + 1})

	for x < i-1 || y < j-1 {
		if x < i-1 && y < j-1 && dp[x+1][y+1] >= dp[x+1][y] && dp[x+1][y+1] >= dp[x][y+1] {
			x, y = x+1, y+1
		} else if x < i-1 && dp[x+1][y] >= dp[x][y+1] {
			x++
		} else {
			y++
		}
		path = append(path, [2]int{x + 1, y + 1})
	}

	return path, totalPrize
}

func getMaxNum(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= c {
		return b
	}
	return c
}

func main() {
	calcReversePath()

	path, totalPrize := buildPath()

	fmt.Println("Максимальний приз:", totalPrize)
	fmt.Println("Оптимальний шлях:", path)
}
