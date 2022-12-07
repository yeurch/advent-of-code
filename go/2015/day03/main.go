package main

import (
	_ "embed"
	"fmt"
	"github.com/yeurch/advent-of-code/go/ysl"
	"strconv"
	"time"
)

//go:embed input.txt
var inputDay string

func Run(input string) (int, int) {
	houses := ysl.NewSet[string]()
	houses2 := ysl.NewSet[string]()
	x, y := 0, 0
	xa, ya := 0, 0
	xb, yb := 0, 0
	houses.Add("0,0")
	houses2.Add("0,0")

	for i, c := range input {
		var robotX, robotY *int
		if i%2 == 0 {
			robotX = &xa
			robotY = &ya
		} else {
			robotX = &xb
			robotY = &yb
		}

		switch c {
		case '^':
			y += 1
			*robotY += 1
		case 'v':
			y -= 1
			*robotY -= 1
		case '<':
			x -= 1
			*robotX -= 1
		case '>':
			x += 1
			*robotX += 1
		}
		coords := strconv.Itoa(x) + "," + strconv.Itoa(y)
		robotCoords := strconv.Itoa(*robotX) + "," + strconv.Itoa(*robotY)
		houses.Add(coords)
		houses2.Add(robotCoords)
	}
	result1 := houses.Cardinality()
	result2 := houses2.Cardinality()

	return result1, result2
}

func main() {
	fmt.Println("*** Advent of Code 2015, day 03 ***")

	start := time.Now()
	p1, p2 := Run(inputDay)
	fmt.Println("part1: ", p1)
	fmt.Println("part2: ", p2)
	fmt.Println(time.Since(start))
}
