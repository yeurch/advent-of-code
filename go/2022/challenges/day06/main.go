package main

import (
	_ "embed"
	"fmt"
	"github.com/yeurch/advent-of-code/2022/ysl"
	"time"
)

//go:embed input.txt
var inputDay string

func Part1(input string) int {
	return detectUnique(input, 4)
}

func Part2(input string) int {
	return detectUnique(input, 14)
}

func detectUnique(input string, uniqueLen int) int {
	var result int

	for i := 0; i < len(input)-uniqueLen+1; i++ {
		slice := input[i : i+uniqueLen]
		if isUnique([]byte(slice)) {
			result = i + uniqueLen
			break
		}
	}
	return result
}

func isUnique(slice []byte) bool {
	s := ysl.NewSet[byte]()
	for _, c := range slice {
		s.Add(c)
	}
	return len(slice) == s.Cardinality()
}

func main() {
	fmt.Println("*** Advent of Code 2022, day 06 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
