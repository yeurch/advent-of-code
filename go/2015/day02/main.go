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

func Run(input string) (int, int) {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	result1 := 0
	result2 := 0

	for _, line := range lines {
		d := strings.Split(line, "x")
		x, _ := strconv.Atoi(d[0])
		y, _ := strconv.Atoi(d[1])
		z, _ := strconv.Atoi(d[2])

		max := 0
		if x > max {
			max = x
		}
		if y > max {
			max = y
		}
		if z > max {
			max = z
		}
		extra := x * y * z / max
		bowVol := x * y * z
		result1 += 2*(x*y+y*z+z*x) + extra
		result2 += 2*(x+y+z-max) + bowVol
	}
	return result1, result2
}

func main() {
	fmt.Println("*** Advent of Code 2015, day 02 ***")

	start := time.Now()
	p1, p2 := Run(inputDay)
	fmt.Println("part1: ", p1)
	fmt.Println("part2: ", p2)
	fmt.Println(time.Since(start))
}
