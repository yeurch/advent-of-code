package main

import (
	_ "embed"
	"fmt"
	"github.com/yeurch/advent-of-code/go/ysl"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type route struct {
	from string
	to   string
}

type collector struct {
	val int
	f   func(int, int) bool
}

func (c *collector) apply(x int) {
	if c.f(x, c.val) {
		c.val = x
	}
}

func Part1(input string) int {
	minFinder := collector{math.MaxInt, func(i int, res int) bool {
		return i < res
	}}
	doPuzzle(input, &minFinder)
	return minFinder.val
}

func Part2(input string) int {
	maxFinder := collector{0, func(i int, res int) bool {
		return i > res
	}}
	doPuzzle(input, &maxFinder)
	return maxFinder.val
}

func doPuzzle(input string, col *collector) {
	routes, places := parseInput(input)

	ysl.Perm[string](places, func(p []string) {
		dist := 0
		for i := 0; i < len(p)-1; i++ {
			dist += routes[route{p[i], p[i+1]}]
		}
		col.apply(dist)
	})
}

func parseInput(input string) (map[route]int, []string) {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	routes := make(map[route]int)
	placeSet := ysl.NewSet[string]()
	r := regexp.MustCompile("(\\w+) to (\\w+) = (\\d+)")
	for _, line := range lines {
		m := r.FindStringSubmatch(line)
		dist, _ := strconv.Atoi(m[3])
		routes[route{m[1], m[2]}] = dist
		routes[route{m[2], m[1]}] = dist
		placeSet.Add(m[1])
		placeSet.Add(m[2])
	}

	places := make([]string, 0)
	for true {
		p, ok := placeSet.Pop()
		if !ok {
			break
		}
		places = append(places, p)
	}
	return routes, places
}

func main() {
	fmt.Println("*** Advent of Code 2015, day 09 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
