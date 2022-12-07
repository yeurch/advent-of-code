package main

import (
	_ "embed"
	"fmt"
	"github.com/yeurch/advent-of-code/go/ysl"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func Part1(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	result := 0
	for _, line := range lines {
		partA := line[0 : len(line)/2]
		partB := line[len(line)/2:]

		for _, c := range partA {
			if strings.ContainsRune(partB, c) {
				result += getPriority(c)
				break
			}
		}
	}
	return result
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	result := 0
	for i := 0; i < len(lines); i += 3 {
		setA := toCharSet(lines[i])
		setB := toCharSet(lines[i+1])
		setC := toCharSet(lines[i+2])
		common := setA.Intersect(setB).Intersect(setC)
		if common.Cardinality() != 1 {
			panic("Unexpected number of items in intersction")
		}
		commonByte, _ := common.Pop()
		result += getPriority(rune(commonByte))

	}
	return result
}

func toCharSet(s string) ysl.Set[byte] {
	return ysl.NewSet[byte]([]byte(s)...)
}

func getPriority(c rune) int {
	if c >= 'a' {
		return int(c - 'a' + 1)
	}
	return int(c - 'A' + 27)
}

func main() {
	fmt.Println("*** Advent of Code 2022, day 03 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
