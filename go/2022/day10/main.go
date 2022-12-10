package main

import (
	_ "embed"
	"fmt"
	"github.com/yeurch/advent-of-code/go/ysl"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func solvePuzzle(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	currentVal := 1
	values := make([]int, 0)
	values = append(values, currentVal)

	for _, line := range lines {
		values = append(values, currentVal)
		if strings.HasPrefix(line, "addx ") {
			dx, _ := strconv.Atoi(line[5:])
			currentVal += dx
			values = append(values, currentVal)
		}
	}

	// Render our output
	fmt.Println("") // blank line above for clarity
	for i := 0; i < 240; i++ {
		colNum := i % 40
		if ysl.Abs(colNum-values[i]) <= 1 {
			fmt.Print("#")
		} else {
			fmt.Print(" ")
		}
		if colNum == 39 {
			fmt.Println("")
		}
	}
	fmt.Println("")

	result := 0
	for i := 19; i < 220; i += 40 {
		result += (i + 1) * values[i]
	}
	return result
}

func main() {
	fmt.Println("*** Advent of Code 2022, day 10 ***")

	start := time.Now()
	fmt.Println("part1: ", solvePuzzle(inputDay))
	fmt.Println("part2 is shown above in graphical output")
	fmt.Println(time.Since(start))
}
