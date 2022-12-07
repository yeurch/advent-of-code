package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func Part1(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	r := regexp.MustCompile("(\\d+)-(\\d+),(\\d+)-(\\d+)")
	result := 0

	for _, line := range lines {
		a, b, c, d := parseLine(r, line)
		if a <= c && b >= d || c <= a && d >= b {
			result += 1
		}
	}
	return result
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	r := regexp.MustCompile("(\\d+)-(\\d+),(\\d+)-(\\d+)")
	result := 0

	for _, line := range lines {
		a, b, c, d := parseLine(r, line)
		if a <= c && b >= c || a >= c && d >= a {
			result += 1
		}
	}
	return result
}

func parseLine(r *regexp.Regexp, line string) (int, int, int, int) {
	groups := r.FindStringSubmatch(line)
	a, _ := strconv.Atoi(groups[1])
	b, _ := strconv.Atoi(groups[2])
	c, _ := strconv.Atoi(groups[3])
	d, _ := strconv.Atoi(groups[4])
	return a, b, c, d
}

func main() {
	fmt.Println("*** Advent of Code 2022, day 04 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
