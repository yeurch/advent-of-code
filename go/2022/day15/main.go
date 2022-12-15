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

type nullableInt struct {
	value int
	isSet bool
}

func (ni *nullableInt) set(val int) {
	ni.value = val
	ni.isSet = true
}

type point struct {
	x int
	y int
}

func (p point) ManhattanDistance(other point) int {
	return ysl.Abs(other.x-p.x) + ysl.Abs(other.y-p.y)
}

type sensor struct {
	loc                point
	nearestBeacon      point
	_ManhattanDistance nullableInt
}

func (s sensor) ManhattanDistance() int {
	if !s._ManhattanDistance.isSet {
		s._ManhattanDistance.set(s.loc.ManhattanDistance(s.nearestBeacon))
	}
	return s._ManhattanDistance.value
}

func Part1(input string) int {
	sensors := parseInput(input)

	minX, maxX := math.MaxInt, math.MinInt
	for _, s := range sensors {
		md := s.ManhattanDistance()
		if s.loc.x-md < minX {
			minX = s.loc.x - md
		}
		if s.loc.x+md > maxX {
			maxX = s.loc.x + md
		}
	}

	const y = 2000000
	result := 0
	for x := minX; x <= maxX; x++ {
		p := point{x, y}
		for _, s := range sensors {
			if s.loc.ManhattanDistance(p) <= s.ManhattanDistance() && p != s.nearestBeacon {
				result++
				break
			}
		}
	}
	return result
}

func Part2(input string) int64 {
	sensors := parseInput(input)

	const maxCoord = 4000000

	for y := 0; y <= maxCoord; y++ {
		for x := 0; x <= maxCoord; x++ {
			p := point{x, y}
			bFound := true
			for _, s := range sensors {
				if s.loc.ManhattanDistance(p) <= s.ManhattanDistance() {
					maxXForThisSensor := s.ManhattanDistance() - ysl.Abs(y-s.loc.y) + s.loc.x
					x = maxXForThisSensor // Jump to the end of where this sensor blocks
					bFound = false
					break
				}
			}
			if bFound {
				return int64(p.x)*int64(maxCoord) + int64(p.y)
			}
		}
	}

	panic("Could not find a suitable point for beacon")
}

func parseInput(input string) []sensor {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	// Sensor at x=2880363, y=3875503: closest beacon is at x=2729330, y=3697325
	r := regexp.MustCompile("Sensor at x=([0-9\\-]+), y=([0-9\\-]+): closest beacon is at x=([0-9\\-]+), y=([0-9\\-]+)")
	result := make([]sensor, len(lines))

	for i, line := range lines {
		m := r.FindStringSubmatch(line)
		sx, _ := strconv.Atoi(m[1])
		sy, _ := strconv.Atoi(m[2])
		bx, _ := strconv.Atoi(m[3])
		by, _ := strconv.Atoi(m[4])
		var s sensor
		s.loc = point{sx, sy}
		s.nearestBeacon = point{bx, by}
		result[i] = s
	}
	return result
}

func main() {
	fmt.Println("*** Advent of Code 2022, day 15 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
