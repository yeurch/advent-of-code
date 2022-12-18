package main

import (
	"container/list"
	_ "embed"
	"fmt"
	"github.com/yeurch/advent-of-code/go/ysl"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type point struct {
	x int
	y int
	z int
}

func (p point) Add(other point) point {
	return point{p.x + other.x, p.y + other.y, p.z + other.z}
}

func Part1(input string) int {
	cubes := parseInput(input)
	s := ysl.NewSet[point](cubes...)

	cardinalDirections := []point{
		{1, 0, 0}, {-1, 0, 0},
		{0, 1, 0}, {0, -1, 0},
		{0, 0, 1}, {0, 0, -1},
	}

	result := 0

	for _, cube := range cubes {
		for _, direction := range cardinalDirections {
			pos := cube.Add(direction)
			if !s.Contains(pos) {
				result++
			}
		}
	}

	return result
}

func Part2(input string) int {
	cubeArray := parseInput(input)
	cubes := ysl.NewSet[point](cubeArray...)

	var minX, maxX, minY, maxY, minZ, maxZ int
	minX = math.MaxInt
	minY = math.MaxInt
	minZ = math.MaxInt
	for _, c := range cubeArray {
		if c.x > maxX {
			maxX = c.x
		} else if c.x < minX {
			minX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		} else if c.y < minY {
			minY = c.y
		}
		if c.z > maxZ {
			maxZ = c.z
		} else if c.z < minZ {
			minZ = c.z
		}
	}

	// We want to _surround_ the shape with water, so extend bounds by one
	minX--
	minY--
	minZ--
	maxX++
	maxY++
	maxZ++

	// Populate our initial water queue with any location one outside the max bounds of our shape
	waterQueue := list.New()
	waterQueue.PushBack(point{minX, minY, minZ})

	cardinalDirections := []point{
		{1, 0, 0}, {-1, 0, 0},
		{0, 1, 0}, {0, -1, 0},
		{0, 0, 1}, {0, 0, -1},
	}

	result := 0
	waterVisited := ysl.NewSet[point]()

	for el := waterQueue.Front(); el != nil; el = el.Next() {
		current := el.Value.(point)
		if waterVisited.Contains(current) {
			continue
		}
		waterVisited.Add(current)

		for _, d := range cardinalDirections {
			p := current.Add(d)
			if cubes.Contains(p) {
				result++ // we found an exposed surface
			} else if p.x >= minX && p.x <= maxX && p.y >= minY && p.y <= maxY && p.z >= minZ && p.z <= maxZ {
				waterQueue.PushBack(p) // we discovered space to flood with water to continue our search
			}
		}
	}

	return result
}

func parseInput(input string) []point {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	result := make([]point, len(lines))
	for i, line := range lines {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		result[i] = point{x, y, z}
	}
	return result
}

func main() {
	fmt.Println("*** Advent of Code 2022, day 18 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
