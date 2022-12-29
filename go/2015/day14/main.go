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

type reindeer struct {
	speed    int
	flyTime  int
	restTime int
}

func (r *reindeer) distAfter(seconds int) int {
	cycleTime := r.flyTime + r.restTime
	fullCycles := seconds / cycleTime
	totalFlyTime := fullCycles * r.flyTime

	extraTime := seconds % cycleTime
	if extraTime > r.flyTime {
		extraTime = r.flyTime
	}
	totalFlyTime += extraTime
	return totalFlyTime * r.speed
}

func Part1(input string) int {
	reindeer := parseInput(input)
	result := 0
	for _, r := range reindeer {
		dist := r.distAfter(2503)
		if dist > result {
			result = dist
		}
	}
	return result
}

func Part2(input string) int {
	reindeer := parseInput(input)
	scores := make([]int, len(reindeer))
	distances := make([]int, len(reindeer))

	for i := 0; i < 2503; i++ {
		maxDist := 0
		for j, r := range reindeer {
			distances[j] = r.distAfter(i + 1)
			if distances[j] > maxDist {
				maxDist = distances[j]
			}
		}
		for j, d := range distances {
			if d == maxDist {
				scores[j] += 1
			}
		}
	}

	result := 0
	for _, s := range scores {
		if s > result {
			result = s
		}
	}
	return result
}

func parseInput(input string) []reindeer {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	// Vixen can fly 18 km/s for 5 seconds, but then must rest for 84 seconds.
	r := regexp.MustCompile("(\\w+) can fly (\\d+) km/s for (\\d+) seconds, but then must rest for (\\d+) seconds.")

	result := make([]reindeer, len(lines))
	for i, line := range lines {
		m := r.FindStringSubmatch(line)
		speed, _ := strconv.Atoi(m[2])
		flyTime, _ := strconv.Atoi(m[3])
		restTime, _ := strconv.Atoi(m[4])
		result[i] = reindeer{speed, flyTime, restTime}
	}
	return result
}

func main() {
	fmt.Println("*** Advent of Code 2015, day 14 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
