package main

import (
	_ "embed"
	"fmt"
	"github.com/yeurch/advent-of-code/go/ysl"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type pos struct {
	x int
	y int
}

type move struct {
	dx int
	dy int
	n  int
}

func Part1(input string) int {
	return doPuzzle(input, 2)
}

func Part2(input string) int {
	return doPuzzle(input, 10)
}

func doPuzzle(input string, numSegments int) int {
	moves := parseInput(input)
	visited := ysl.NewSet[pos]()
	visited.Add(pos{0, 0})

	segmentPositions := make([]pos, numSegments)
	tailSegment := &segmentPositions[numSegments-1]

	for _, m := range moves {
		for i := 0; i < m.n; i++ {
			doMove(segmentPositions, m.dx, m.dy)
			visited.Add(*tailSegment)
		}
	}

	return visited.Cardinality()
}

func doMove(segments []pos, dxHead int, dyHead int) {
	// Move head first
	segments[0].x += dxHead
	segments[0].y += dyHead

	// for each non-head segment
	for i := 1; i < len(segments); i++ {
		tailToHeadX := segments[i-1].x - segments[i].x
		tailToHeadY := segments[i-1].y - segments[i].y

		if ysl.Abs(tailToHeadX) <= 1 && ysl.Abs(tailToHeadY) <= 1 {
			// No need for any further moves
			break
		}

		if tailToHeadX != 0 {
			segments[i].x += tailToHeadX / ysl.Abs(tailToHeadX)
		}
		if tailToHeadY != 0 {
			segments[i].y += tailToHeadY / ysl.Abs(tailToHeadY)
		}
	}
}

func parseInput(input string) []move {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	result := make([]move, len(lines))

	for i, line := range lines {
		var dx, dy int
		switch line[0] {
		case 'U':
			dy = -1
		case 'D':
			dy = 1
		case 'R':
			dx = 1
		case 'L':
			dx = -1
		}
		n, _ := strconv.Atoi(line[2:])
		result[i] = move{dx, dy, n}
	}
	return result
}

func main() {
	fmt.Println("*** Advent of Code 2022, day 09 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
