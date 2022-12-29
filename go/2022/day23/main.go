package main

import (
	_ "embed"
	"fmt"
	"github.com/yeurch/advent-of-code/go/ysl"
	"math"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type Point struct {
	x int
	y int
}

func (p *Point) Add(other Point) Point {
	return Point{p.x + other.x, p.y + other.y}
}

type Rule struct {
	direction  Point
	ifAllEmpty []Point
}

var rule0 = Rule{Point{0, 0}, []Point{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}}
var rules = []Rule{
	{Point{0, 1}, []Point{{-1, 1}, {0, 1}, {1, 1}}},
	{Point{0, -1}, []Point{{-1, -1}, {0, -1}, {1, -1}}},
	{Point{-1, 0}, []Point{{-1, -1}, {-1, 0}, {-1, 1}}},
	{Point{1, 0}, []Point{{1, -1}, {1, 0}, {1, 1}}},
}

func Part1(input string) int {
	elves := parseInput(input)

	for i := 0; i < 10; i++ {
		elves, _ = doMove(i, &elves)
	}

	return getNumEmpty(&elves)
}
func Part2(input string) int {
	elves := parseInput(input)

	i := 0
	for {
		var n int
		elves, n = doMove(i, &elves)
		if n == 0 {
			break
		}
		i++
	}

	return i + 1
}

// Processes a move and returns the new positions of elves, along with the number of elves that moved
func doMove(moveNum int, elves *[]Point) ([]Point, int) {
	startPositions := ysl.NewSet[Point]()
	for _, p := range *elves {
		startPositions.Add(p)
	}

	// Determine intentions
	numMoved := 0
	intentions := append([]Point{}, *elves...) // Make a copy of elves, so each initial intention is to move nowhere
	for j, p := range intentions {
		if checkRule(&startPositions, &p, &rule0) {
			continue // This elf won't move, as it has no neighbours
		}
		for ruleNum := 0; ruleNum < len(rules); ruleNum++ {
			rule := rules[(moveNum+ruleNum)%len(rules)]
			if checkRule(&startPositions, &p, &rule) {
				intentions[j] = p.Add(rule.direction)
				numMoved++
				break
			}
		}

	}

	// Duplicate intention detection
	dupes := ysl.NewSet[Point]()
	intentionSet := ysl.NewSet[Point]()
	for _, intention := range intentions {
		ok := intentionSet.Add(intention)
		if !ok {
			dupes.Add(intention)
		}
	}

	for i, p := range intentions {
		if dupes.Contains(p) {
			intentions[i] = (*elves)[i] // Cancel the intent of this elf
			numMoved--
		}
	}

	return intentions, numMoved
}
func checkRule(startPositions *ysl.Set[Point], p *Point, rule *Rule) bool {
	for _, checkDir := range rule.ifAllEmpty {
		if (*startPositions).Contains(p.Add(checkDir)) {
			return false
		}
	}
	return true
}

func parseInput(input string) []Point {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	result := make([]Point, 0)
	for i, line := range lines {
		y := len(lines) - i - 1
		for x, c := range line {
			if c == '#' {
				result = append(result, Point{x, y})
			}
		}
	}

	return result
}

func getBounds(elves *[]Point) [2]Point {
	minX, maxX, minY, maxY := math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	for _, p := range *elves {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	var bounds [2]Point
	bounds[0].x = minX
	bounds[0].y = minY
	bounds[1].x = maxX
	bounds[1].y = maxY
	return bounds
}
func getNumEmpty(elves *[]Point) int {
	bounds := getBounds(elves)
	area := (bounds[1].x - bounds[0].x + 1) * (bounds[1].y - bounds[0].y + 1)
	return area - len(*elves)
}

func main() {
	fmt.Println("*** Advent of Code 2022, day 23 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
