package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

const BASE = 5

var snafuToDecMap = map[byte]int64{
	'=': -2,
	'-': -1,
	'0': 0,
	'1': 1,
	'2': 2,
}

var decToSnafuMap = map[int64]byte{
	-2: '=',
	-1: '-',
	0:  '0',
	1:  '1',
	2:  '2',
}

func Part1(input string) string {
	snafuValues := parseInput(input)

	total := int64(0)
	for _, snafuValue := range snafuValues {
		total += snafuToDec(snafuValue)
	}

	return decToSNAFU(total)
}

func parseInput(input string) []string {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	return lines
}

func snafuToDec(snafu string) int64 {
	result := int64(0)
	posVal := int64(1)
	for i := len(snafu) - 1; i >= 0; i-- {
		c := snafu[i]
		result += snafuToDecMap[c] * posVal
		posVal *= BASE
	}
	return result
}

func decToSNAFU(val int64) string {
	posVal := int64(1)
	for posVal*BASE/2 < val {
		posVal *= BASE
	}

	result := ""
	for posVal > 0 {
		digit := int64(math.Round(float64(val) / float64(posVal)))
		result += string(decToSnafuMap[digit])
		val -= digit * posVal
		posVal /= BASE
	}
	return result
}

func main() {
	fmt.Println("*** Advent of Code 2022, day 25 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))
}
