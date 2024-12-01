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

type Point struct {
	x int
	y int
}

func (p *Point) Add(other Point) Point {
	return Point{p.x + other.x, p.y + other.y}
}

var facings = []Point{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func Part1(input string) int {
	grid, instructions := parseInput(input)

	startX := strings.Index(grid[0], ".")
	pos := Point{startX, 0}
	facing := 0
	for _, instruction := range instructions {
		n, err := strconv.Atoi(instruction)
		if err != nil {
			// This is a facing change
			switch instruction {
			case "R":
				facing = (facing + 1) % len(facings)
			case "L":
				facing = (facing + len(facings) - 1) % len(facings)
			default:
				panic("Unexpected direction")
			}
			continue
		}

		// Now we have n containing the number of steps to take
		newPos := pos
		for n > 0 {
			newPos = newPos.Add(facings[facing])
			if newPos.y >= len(grid) {
				newPos.y = 0
			} else if newPos.y < 0 {
				newPos.y = len(grid) - 1
			}
			if newPos.x >= len(grid[newPos.y]) {
				newPos.x = 0
			} else if newPos.x < 0 {
				newPos.x = len(grid[newPos.y]) - 1
			}

			tile := grid[newPos.y][newPos.x]
			switch tile {
			case '#':
				// We hit a wall
				n = 0
			case '.':
				n--
				pos = newPos
			case ' ':
			default:
				panic("Unexpected tile found")
			}
		}
	}

	return 1000*(pos.y+1) + 4*(pos.x+1) + facing
}

func Part2(input string) int {
	grid, instructions := parseInput(input)

	startX := strings.Index(grid[0], ".")
	pos := Point{startX, 0}
	facing := 0
	for _, instruction := range instructions {
		n, err := strconv.Atoi(instruction)
		if err != nil {
			// This is a facing change
			switch instruction {
			case "R":
				facing = (facing + 1) % len(facings)
			case "L":
				facing = (facing + len(facings) - 1) % len(facings)
			default:
				panic("Unexpected direction")
			}
			continue
		}

		// Now we have n containing the number of steps to take
		newPos := pos
		for n > 0 {
			newPos = newPos.Add(facings[facing])
			if newPos.y >= len(grid) {
				newPos.y = 0
			} else if newPos.y < 0 {
				newPos.y = len(grid) - 1
			}
			if newPos.x >= len(grid[newPos.y]) {
				newPos.x = 0
			} else if newPos.x < 0 {
				newPos.x = len(grid[newPos.y]) - 1
			}

			tile := grid[newPos.y][newPos.x]
			switch tile {
			case '#':
				// We hit a wall
				n = 0
			case '.':
				n--
				pos = newPos
			case ' ':
			default:
				panic("Unexpected tile found")
			}
		}
	}

	return 1000*(pos.y+1) + 4*(pos.x+1) + facing
}

func parseInput(input string) ([]string, []string) {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	movementInstructions := lines[len(lines)-1]
	gridLines := lines[0 : len(lines)-2]

	// Determine grid width based on longest line
	gridWidth := 0
	for _, l := range gridLines {
		if len(l) > gridWidth {
			gridWidth = len(l)
		}
	}

	// Adjust line lengths
	for i := range gridLines {
		l := &gridLines[i]
		gridLines[i] = *l + strings.Repeat(" ", gridWidth-len(*l))
	}

	// Split our instructions into consecutive series of digits, and letters
	var instructions []string
	processingDigit := true
	var tempNum string
	for _, c := range movementInstructions {
		if c >= '0' && c <= '9' {
			processingDigit = true
			tempNum += string(c)
		} else {
			if processingDigit {
				instructions = append(instructions, tempNum)
				tempNum = ""
				processingDigit = false
			}
			instructions = append(instructions, string(c))
		}
	}
	if processingDigit {
		instructions = append(instructions, tempNum)
	}

	return gridLines, instructions
}

func main() {
	fmt.Println("*** Advent of Code 2022, day nn ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
