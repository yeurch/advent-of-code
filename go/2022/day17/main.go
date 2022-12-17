package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type point struct {
	x int
	y int
}

type shape []point

func (p point) add(other point) point {
	return point{p.x + other.x, p.y + other.y}
}

func (s *shape) height() int {
	result := 0
	for _, p := range *s {
		if p.y > result {
			result = p.y
		}
	}
	return result + 1
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	inputLen := len(input)

	shapes := []shape{
		{{0, 0}, {1, 0}, {2, 0}, {3, 0}},         // horizontal line,
		{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}}, // plus
		{{2, 2}, {2, 1}, {2, 0}, {1, 0}, {0, 0}}, // reverse L
		{{0, 0}, {0, 1}, {0, 2}, {0, 3}},         // vertical line
		{{0, 0}, {0, 1}, {1, 0}, {1, 1}},         // square block
	}
	numShapes := len(shapes)

	grid := make([][]bool, 0)
	turn := 0
	rockTowerHeight := 0

	for rockNum := 0; rockNum < 2022; rockNum++ {
		rock := shapes[rockNum%numShapes]

		// add rows to top of grid as required
		newGridHeight := rockTowerHeight + 3 + rock.height()
		for i := len(grid); i < newGridHeight; i++ {
			grid = append(grid, make([]bool, 7))
		}

		shapePos := point{2, rockTowerHeight + 3} // bottom left of shape
		bRockComplete := false
		for !bRockComplete {
			windDir := input[turn%inputLen]
			lateralMove := point{1, 0}
			if windDir == '<' {
				lateralMove.x = -1
			}
			desiredLoc := shapePos.add(lateralMove)
			if canPlace(&grid, rock, desiredLoc) {
				shapePos = desiredLoc
			}

			desiredLoc = shapePos.add(point{0, -1}) // drop by one
			if canPlace(&grid, rock, desiredLoc) {
				shapePos = desiredLoc
			} else {
				bRockComplete = true
				place(&grid, rock, shapePos)
				topOfShape := shapePos.y + rock.height()
				if topOfShape > rockTowerHeight {
					rockTowerHeight = topOfShape
				}
			}
			turn++
		}
	}
	return rockTowerHeight
}

type towerStatus struct {
	brickNum    int
	towerHeight int
}

func Part2(input string) int64 {
	input = strings.TrimSuffix(input, "\n")
	inputLen := len(input)

	shapes := []shape{
		{{0, 0}, {1, 0}, {2, 0}, {3, 0}},         // horizontal line,
		{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}}, // plus
		{{2, 2}, {2, 1}, {2, 0}, {1, 0}, {0, 0}}, // reverse L
		{{0, 0}, {0, 1}, {0, 2}, {0, 3}},         // vertical line
		{{0, 0}, {0, 1}, {1, 0}, {1, 1}},         // square block
	}
	numShapes := len(shapes)

	grid := make([][]bool, 0)
	turn := 0
	rockTowerHeight := 0

	fingerprints := make(map[string]towerStatus)
	cycleSize := 0
	candidateCycleConsecutiveOccurrences := 0
	bCycleSizeFound := false
	var heightSkipped int64

	rockNum := 0
	for count := math.MaxInt; count >= 0; count-- {
		rockIndex := rockNum % numShapes
		rock := shapes[rockIndex]
		var turnIndex int

		if rockNum == 2022 {
			fmt.Printf("Completed part 1 equivalent ... %d\n", rockTowerHeight)
		}

		// add rows to top of grid as required
		newGridHeight := rockTowerHeight + 3 + rock.height()
		for i := len(grid); i < newGridHeight; i++ {
			grid = append(grid, make([]bool, 7))
		}

		shapePos := point{2, rockTowerHeight + 3} // bottom left of shape
		bRockComplete := false
		for !bRockComplete {
			turnIndex = turn % inputLen
			windDir := input[turnIndex]

			lateralMove := point{1, 0}
			if windDir == '<' {
				lateralMove.x = -1
			}
			desiredLoc := shapePos.add(lateralMove)
			if canPlace(&grid, rock, desiredLoc) {
				shapePos = desiredLoc
			}

			desiredLoc = shapePos.add(point{0, -1}) // drop by one
			if canPlace(&grid, rock, desiredLoc) {
				shapePos = desiredLoc
			} else {
				bRockComplete = true
				place(&grid, rock, shapePos)
				topOfShape := shapePos.y + rock.height()
				if topOfShape > rockTowerHeight {
					rockTowerHeight = topOfShape
				}
			}
			turn++
		}

		if !bCycleSizeFound {
			fp := fingerprint(&grid, rockTowerHeight, rockIndex, turnIndex)
			existingStatus, ok := fingerprints[fp]
			if ok {
				candidateCycleSize := rockNum - existingStatus.brickNum
				if cycleSize == candidateCycleSize {
					candidateCycleConsecutiveOccurrences++
				} else {
					cycleSize = candidateCycleSize
					candidateCycleConsecutiveOccurrences = 0
				}

				if candidateCycleConsecutiveOccurrences > candidateCycleSize {
					heightDelta := rockTowerHeight - existingStatus.towerHeight
					fmt.Printf("Found a cycle: %d --> %d = %d (occurred %d times). Delta height = %d\n",
						existingStatus.brickNum, rockNum, candidateCycleSize,
						candidateCycleConsecutiveOccurrences, heightDelta)

					rocksRemaining := 1000000000000 - int64(rockNum)
					cyclesToSkip := rocksRemaining / int64(candidateCycleSize)
					heightSkipped = cyclesToSkip * int64(heightDelta)
					count = int(rocksRemaining-(cyclesToSkip*int64(candidateCycleSize))) - 1

					bCycleSizeFound = true
				}

			}
			fingerprints[fp] = towerStatus{rockNum, rockTowerHeight}
		}
		rockNum++
	}

	return int64(rockTowerHeight) + heightSkipped
}

func fingerprint(grid *[][]bool, maxHeight int, rockIndex int, turnIndex int) string {
	columnDepths := make([]string, 7)

	depthsFound := 0
	for y := maxHeight - 1; y >= 0; y-- {
		for x := 0; x < 7; x++ {
			if (*grid)[y][x] && depthsFound&(1<<x) == 0 {
				columnDepths[x] = strconv.Itoa(maxHeight - y)
				depthsFound |= 1 << x
			}
		}
		if depthsFound == 127 {
			// all columns found
			break
		}
	}
	depths := strings.Join(columnDepths, ",")
	return fmt.Sprintf("%s;%d;%d", depths, rockIndex, turnIndex)
}

func canPlace(grid *[][]bool, s shape, sPos point) bool {
	for _, p := range s {
		translatedPos := p.add(sPos)
		if translatedPos.x < 0 || translatedPos.x > 6 ||
			translatedPos.y < 0 || (*grid)[translatedPos.y][translatedPos.x] {
			return false
		}
	}
	return true
}

func place(grid *[][]bool, s shape, pos point) {
	for _, p := range s {
		translatedPos := p.add(pos)
		(*grid)[translatedPos.y][translatedPos.x] = true
	}
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
