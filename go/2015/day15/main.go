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

type ingredient struct {
	capacity   int
	durability int
	flavor     int
	texture    int
}

func Part1(input string) int {
	//ingredients := parseInput(input)
	return 0
}

func Part2(input string) int {
	return 0
}

func parseInput(input string) []ingredient {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	// Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
	r := regexp.MustCompile("(\\w+): capacity (.+), durability (.+), flavor (.+), texture (.+), calories (.+)")

	result := make([]ingredient, len(lines))
	for i, line := range lines {
		m := r.FindStringSubmatch(line)
		capacity, _ := strconv.Atoi(m[2])
		durability, _ := strconv.Atoi(m[3])
		flavor, _ := strconv.Atoi(m[4])
		texture, _ := strconv.Atoi(m[5])
		ing := ingredient{capacity, durability, flavor, texture}
		result[i] = ing
	}
	return result
}

func main() {
	fmt.Println("*** Advent of Code 2015, day 15 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
