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
		if isNicePart1(line) {
			result += 1
		}
	}
	return result
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	result := 0

	for _, line := range lines {
		if isNicePart2(line) {
			result += 1
		}
	}
	return result
}

func isNicePart1(line string) bool {
	return !isSubstringNaughty(line) && isVowelNice(line) && isRepeatNice(line)
}

func isNicePart2(line string) bool {
	return isRepeatedPairNice(line) && isAlmostAdjacentNice(line)
}

func isVowelNice(line string) bool {
	vowelCount := 0
	for _, c := range line {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			vowelCount++
			if vowelCount == 3 {
				return true
			}
		}
	}
	return false
}

func isRepeatNice(line string) bool {
	for i := 0; i < len(line)-1; i++ {
		if line[i] == line[i+1] {
			return true
		}
	}
	return false
}

func isSubstringNaughty(line string) bool {
	bannedPairs := []string{"ab", "cd", "pq", "xy"}
	for _, bannedPair := range bannedPairs {
		if strings.Contains(line, bannedPair) {
			return true
		}
	}
	return false
}

func isRepeatedPairNice(line string) bool {
	for i := 0; i < len(line)-3; i++ {
		desiredMatch := line[i : i+2]
		for j := i + 2; j < len(line)-1; j++ {
			if line[j:j+2] == desiredMatch {
				return true
			}
		}
	}
	return false
}

func isAlmostAdjacentNice(line string) bool {
	for i := 0; i < len(line)-2; i++ {
		if line[i] == line[i+2] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("*** Advent of Code 2015, day 05 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
