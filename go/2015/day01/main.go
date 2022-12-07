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
	result := strings.Count(input, "(") - strings.Count(input, ")")
	return result
}

func Part2(input string) int {
	var result int
	currentFloor := 0
	for i, c := range input {
		if c == '(' {
			currentFloor += 1
		} else {
			currentFloor -= 1
		}
		if currentFloor < 0 {
			result = i + 1
			break
		}
	}
	return result
}

func main() {
	fmt.Println("*** Advent of Code 2015, day 01 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
