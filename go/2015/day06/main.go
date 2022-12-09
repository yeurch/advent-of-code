package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

var re *regexp.Regexp

func Part1(input string) int {
	return doPuzzle(input, false)
}

func Part2(input string) int {
	return doPuzzle(input, true)
}

func doPuzzle(input string, isPart2 bool) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	result := 0
	grid := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make([]int, 1000)
	}

	for _, line := range lines {
		processLine(grid, line, isPart2)
	}

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			result += grid[j][i]
		}
	}
	return result
}

func processLine(grid [][]int, line string, isPart2 bool) {
	var op func(grid [][]int, x int, y int)

	if isPart2 {
		if strings.HasPrefix(line, "turn on") {
			op = func(grid [][]int, x int, y int) {
				grid[y][x] += 1
			}
		} else if strings.HasPrefix(line, "turn off") {
			op = func(grid [][]int, x int, y int) {
				if grid[y][x] > 0 {
					grid[y][x] -= 1
				}
			}
		} else {
			op = func(grid [][]int, x int, y int) {
				grid[y][x] += 2
			}
		}
	} else {
		if strings.HasPrefix(line, "turn on") {
			op = func(grid [][]int, x int, y int) {
				grid[y][x] = 1
			}
		} else if strings.HasPrefix(line, "turn off") {
			op = func(grid [][]int, x int, y int) {
				grid[y][x] = 0
			}
		} else {
			op = func(grid [][]int, x int, y int) {
				grid[y][x] = 1 - grid[y][x]
			}
		}
	}

	matches := re.FindStringSubmatch(line)
	x1, _ := strconv.Atoi(matches[1])
	y1, _ := strconv.Atoi(matches[2])
	x2, _ := strconv.Atoi(matches[3])
	y2, _ := strconv.Atoi(matches[4])

	for i := x1 - 1; i < x2; i++ {
		for j := y1 - 1; j < y2; j++ {
			op(grid, i, j)
		}
	}
}

func main() {
	fmt.Println("*** Advent of Code 2015, day 06 ***")

	re = regexp.MustCompile("(\\d+),(\\d+) through (\\d+),(\\d+)")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
