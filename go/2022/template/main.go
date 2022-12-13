package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func Part1(input string) int {
	result := parseInput(input)
	return result
}

func Part2(input string) int {
	result := parseInput(input)
	return result
}

func parseInput(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	return len(lines)
}

func main() {
	fmt.Println("*** Advent of Code 2022, day nn ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
