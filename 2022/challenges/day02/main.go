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
	return doGame(input, parseRound1)
}

func Part2(input string) int {
	return doGame(input, parseRound2)
}

func doGame(input string, parseFn func(string) (int, int)) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	result := 0

	for _, line := range lines {
		a, b := parseFn(line)
		_, myScore := roundScores(a, b)
		result += myScore
	}

	return result
}

func parseRound1(input string) (a, b int) {
	for _, c := range strings.Fields(input) {
		switch c {
		case "A":
			a = 0
		case "B":
			a = 1
		case "C":
			a = 2
		case "X":
			b = 0
		case "Y":
			b = 1
		case "Z":
			b = 2
		}
	}
	return
}

func parseRound2(input string) (a, b int) {
	a, b = parseRound1(input)
	switch b {
	case 0:
		// Need to lose
		b = (a + 2) % 3
	case 1:
		// Need to draw
		b = a
	case 2:
		// Need to win
		b = (a + 1) % 3
	}
	return
}

func roundScores(a, b int) (scoreA, scoreB int) {
	scoreA = a + 1
	scoreB = b + 1
	if a == b {
		// it's a draw!
		scoreA += 3
		scoreB += 3
	} else if (a+1)%3 == b {
		// b wins
		scoreB += 6
	} else {
		// a wins
		scoreA += 6
	}
	return
}

func main() {
	fmt.Println("*** Advent of Code 2022, day 02 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(string(inputDay)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(string(inputDay)))
	fmt.Println(time.Since(start))
}
