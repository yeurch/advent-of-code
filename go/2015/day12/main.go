package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func Part1(input string) int {
	return doPuzzle(input)
}

func Part2(input string) int {
	s := stripRed(input)
	return doPuzzle(s)
}

func doPuzzle(input string) int {
	s := input
	result := 0
	for {
		i := strings.IndexAny(s, "0123456789")
		if i < 0 {
			break
		}
		v := string(s[i])
		signMultiplier := 1
		if s[i-1] == '-' {
			signMultiplier = -1
		}
		for {
			i += 1
			if i == len(s) || s[i] < '0' || s[i] > '9' {
				val, _ := strconv.Atoi(v)
				result += signMultiplier * val
				s = s[i:]
				break
			} else {
				v += string(s[i])
			}
		}

	}
	return result
}

func stripRed(input string) string {
	s := input
	i := strings.Index(s, ":\"red\"")
	for i >= 0 {
		s = removeObjectContents(s, i)
		i = strings.Index(s, ":\"red\"")
	}
	return s
}

func removeObjectContents(s string, i int) string {
	a := search(s, i, -1, '}', '{')
	b := search(s, i, 1, '{', '}')
	return s[:a+1] + s[b:] // keeps the surrounding {} as an empty object
}

func search(s string, i int, di int, incrementer byte, decrementer byte) int {
	bracketCount := 1
	for bracketCount > 0 {
		i += di
		if s[i] == incrementer {
			bracketCount += 1
		} else if s[i] == decrementer {
			bracketCount -= 1
		}
	}
	return i
}

func main() {
	fmt.Println("*** Advent of Code 2015, day 12 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
