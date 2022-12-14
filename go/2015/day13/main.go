package main

import (
	_ "embed"
	"fmt"
	"github.com/yeurch/advent-of-code/go/ysl"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func Part1(input string) int {
	preferences := parseInput(input)
	return doPuzzle(preferences)
}

func Part2(input string) int {
	preferences := parseInput(input)

	me := make(map[string]int)
	for k, v := range preferences {
		(*v)["me"] = 0
		me[k] = 0
	}
	preferences["me"] = &me

	return doPuzzle(preferences)
}

func doPuzzle(preferences map[string]*map[string]int) int {
	var names []string
	for k := range preferences {
		names = append(names, k)
	}
	numSeats := len(names)

	result := 0
	ysl.Perm[string](names, func(permedNames []string) {
		if permedNames[0] == names[0] { // circular table, so fix one name for efficiency
			score := 0
			for i, name := range permedNames {
				neighbours := []string{permedNames[(i+1)%numSeats], permedNames[(i+numSeats-1)%numSeats]}
				for _, neighbour := range neighbours {
					score += (*preferences[name])[neighbour]
				}
			}
			if score > result {
				result = score
			}
		}
	})
	return result
}

func parseInput(input string) map[string]*map[string]int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	// Mallory would lose 36 happiness units by sitting next to Alice.
	r := regexp.MustCompile("(\\w+) would (gain|lose) (\\d+) happiness units by sitting next to (\\w+).")
	result := make(map[string]*map[string]int)

	for _, line := range lines {
		m := r.FindStringSubmatch(line)
		nameA := m[1]
		nameB := m[4]
		amt, _ := strconv.Atoi(m[3])
		if m[2] == "lose" {
			amt *= -1
		}

		submap, ok := result[nameA]
		if !ok {
			m := make(map[string]int)
			submap = &m
			result[nameA] = submap
		}
		(*submap)[nameB] = amt
	}
	return result
}

func main() {
	fmt.Println("*** Advent of Code 2015, day 13 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
