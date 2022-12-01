package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func Part1(input string) int {
	lines := strings.Split(input, "\n")

	result := 0
	thisElf := 0
	for _, line := range lines {
		if len(line) == 0 {
			if thisElf > result {
				result = thisElf
			}
			thisElf = 0
		} else {
			value, _ := strconv.Atoi(line)
			thisElf += value
		}
	}
	return result
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")

	var elves []int
	thisElf := 0
	for _, line := range lines {
		if len(line) == 0 {
			elves = append(elves, thisElf)
			thisElf = 0
		} else {
			value, _ := strconv.Atoi(line)
			thisElf += value
		}
	}
	sort.Ints(elves)
	result := sum(elves[len(elves)-3:])
	return result
}

func sum(values []int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func main() {
	fmt.Println("*** Advent of Code 2022, day 01 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(string(inputDay)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(string(inputDay)))
	fmt.Println(time.Since(start))
}
