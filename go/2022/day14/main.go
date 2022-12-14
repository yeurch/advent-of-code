package main

import (
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
}

func Part1(input string) int {
	caves := parseInput(input)

	abyssLevel := 0
	for k := range caves {
		if k.y > abyssLevel {
			abyssLevel = k.y
		}
	}

	result := 0
	for !dropSand(&caves, abyssLevel) {
		result += 1
	}

	return result
}

func Part2(input string) int {
	caves := parseInput(input)

	floorLevel := 0
	for k := range caves {
		if k.y > floorLevel {
			floorLevel = k.y
		}
	}
	floorLevel += 2
	for x := 499 - floorLevel; x <= 501+floorLevel; x++ {
		caves[point{x, floorLevel}] = '#'
	}

	result := 0
	startPoint := point{500, 0}
	for true {
		dropSand(&caves, math.MaxInt)
		result += 1
		_, full := caves[startPoint]
		if full {
			break
		}
	}

	return result
}

// Returns true if the sand drops into the abyss
func dropSand(caves *map[point]byte, abyssLevel int) bool {
	p := point{500, 0}
	ok := true
	for ok {
		p, ok = moveSand(caves, p)
		if p.y > abyssLevel {
			return true
		}
	}
	(*caves)[p] = 'o'
	return false
}

// bool result is true if the sand moved ok
func moveSand(caves *map[point]byte, p point) (point, bool) {
	targets := []point{{p.x, p.y + 1}, {p.x - 1, p.y + 1}, {p.x + 1, p.y + 1}}
	for _, target := range targets {
		_, ok := (*caves)[target]
		if !ok {
			// Spaces is empty
			return target, true
		}
	}
	return p, false
}

func parseInput(input string) map[point]byte {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	result := make(map[point]byte)
	for _, line := range lines {
		bFirst := true
		var x, y int
		for _, p := range strings.Split(line, " -> ") {
			pSplit := strings.Split(p, ",")
			xTarg, _ := strconv.Atoi(pSplit[0])
			yTarg, _ := strconv.Atoi(pSplit[1])
			if bFirst {
				x, y, bFirst = xTarg, yTarg, false
			}
			dx, dy := ysl.Sgn(xTarg-x), ysl.Sgn(yTarg-y)
			result[point{x, y}] = '#'
			for x != xTarg || y != yTarg {
				x += dx
				y += dy
				result[point{x, y}] = '#'
			}
		}

	}
	return result
}

func main() {
	fmt.Println("*** Advent of Code 2022, day 14 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
