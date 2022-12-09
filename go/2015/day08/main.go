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
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	result := 0
	for _, line := range lines {
		result += 2 // surrounding double quotes
		for i := 0; i < len(line); i++ {
			c := line[i]
			if c == '\\' {
				result += 1
				// We'll assume there's always another character otherwise the input isn't well-formed
				if line[i+1] == 'x' { // a hex escaped char
					result += 2 // there will be a further two "bonus" chars
				}
				i++ // skip the next char, so we don't process a literal backslash as another backslash
			}
		}
	}
	return result
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	result := 0
	for _, line := range lines {
		result += 2 // surrounding double quotes
		for _, c := range line {
			if c == '\\' || c == '"' {
				result += 1
			}
		}
	}
	return result
}

func main() {
	fmt.Println("*** Advent of Code 2015, day 08 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
