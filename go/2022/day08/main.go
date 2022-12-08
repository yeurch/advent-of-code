package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func Part1(input string) int {
	plantation := parseInput(input)
	result := 0

	for i, row := range plantation {
		for j := range row {
			if isVisible(plantation, i, j) {
				result += 1
			}
		}
	}

	return result
}

func Part2(input string) int {
	plantation := parseInput(input)
	result := 0

	for i, row := range plantation {
		for j := range row {
			score := scenicScore(plantation, i, j)
			if score > result {
				result = score
			}
		}
	}

	return result
}

func parseInput(input string) [][]uint8 {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	result := make([][]uint8, len(lines))

	for i, line := range lines {
		row := make([]uint8, len(line))
		for j, c := range line {
			row[j] = (uint8)(c - '0')
		}
		result[i] = row
	}
	return result
}

func isVisible(plantation [][]uint8, x int, y int) bool {
	return isVisibleDir(plantation, x, y, 0, 1) ||
		isVisibleDir(plantation, x, y, 0, -1) ||
		isVisibleDir(plantation, x, y, 1, 0) ||
		isVisibleDir(plantation, x, y, -1, 0)
}

func isVisibleDir(plantation [][]uint8, x int, y int, dx int, dy int) bool {
	height := plantation[y][x]
	visible := true
	i, j := x, y
	for true {
		i += dx
		j += dy
		if i < 0 || j < 0 || j >= len(plantation) || i >= len(plantation[j]) {
			break
		}
		if plantation[j][i] >= height {
			visible = false
			break
		}
	}

	return visible
}

func scenicScore(plantation [][]uint8, x int, y int) int {
	return scenicScoreDir(plantation, x, y, 0, 1) *
		scenicScoreDir(plantation, x, y, 0, -1) *
		scenicScoreDir(plantation, x, y, 1, 0) *
		scenicScoreDir(plantation, x, y, -1, 0)
}

func scenicScoreDir(plantation [][]uint8, x int, y int, dx int, dy int) int {
	height := plantation[y][x]
	result := 0
	i, j := x, y
	for true {
		i += dx
		j += dy
		if i < 0 || j < 0 || j >= len(plantation) || i >= len(plantation[j]) {
			break
		}
		result++
		if plantation[j][i] >= height {
			break
		}
	}

	return result
}

func main() {
	fmt.Println("*** Advent of Code 2022, day 08 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
