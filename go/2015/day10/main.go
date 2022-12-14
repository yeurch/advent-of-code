package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed conwayElements.txt
var conwayElements string

type element struct {
	symbol   string
	value    string
	decaysTo []string
}

func Part1(input string) int {
	return len(doPuzzle(input, 40))
}

func Part2(input string) int {
	elements := parseInput(conwayElements)
	var start string
	for k, v := range elements {
		if v.value == input {
			start = k
			break
		}
	}
	if start == "" {
		// If we panic here, we'll need to write code to decompose our input into a number of elements
		panic("Start value is not an element: " + input)
	}

	buckets := make(map[string]int)
	buckets[start] = 1
	for roundNum := 0; roundNum < 50; roundNum++ {
		newBuckets := make(map[string]int)
		for k, v := range buckets {
			targets := elements[k].decaysTo
			for _, t := range targets {
				newBuckets[t] += v
			}
		}
		buckets = newBuckets
	}

	result := 0
	for k, v := range buckets {
		elementLength := len(elements[k].value)
		result += v * elementLength
	}

	return result

}

func doPuzzle(input string, numRounds int) string {
	s := input
	for round := 0; round < numRounds; round++ {
		last := s
		s = ""
		searchChar := byte(0)
		n := 0
		for i := 0; i < len(last); i++ {
			if searchChar == byte(0) {
				searchChar = last[i]
				n = 1
			} else if searchChar == last[i] {
				n += 1
			} else {
				s += strconv.Itoa(n)
				s += string(searchChar)
				searchChar = last[i]
				n = 1
			}
		}
		s += strconv.Itoa(n)
		s += string(searchChar)
		fmt.Print(".")
	}
	fmt.Println()
	return s
}

func parseInput(input string) map[string]element {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	result := make(map[string]element)
	for _, line := range lines {
		parts := strings.Fields(line)
		symbol := parts[1]
		value := parts[2]
		decaysTo := make([]string, 0)
		for _, el := range strings.Split(parts[3], ".") {
			decaysTo = append(decaysTo, el)
		}
		result[symbol] = element{symbol, value, decaysTo}
	}
	return result
}

func main() {
	inputDay := "1321131112"

	fmt.Println("*** Advent of Code 2022, day 10 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
