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

type point struct {
	x int
	y int
}

type shape []point

type towerStatus struct {
	brickNum    int
	towerHeight int
}

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

type tower struct {
	cycleRocks  int
	cycleHeight int
	heights     []int
}

func (t *tower) HeightAfterRock(n int64) int64 {
	if n <= 0 {
		return 0
	}
	rocksRemaining := n - int64(len(t.heights))
	if rocksRemaining < 0 {
		return int64(t.heights[n-1])
	}
	cyclesToSkip := rocksRemaining/int64(t.cycleRocks) + 1
	result := cyclesToSkip * int64(t.cycleHeight)
	index := n - cyclesToSkip*int64(t.cycleRocks) - 1
	result += int64(t.heights[index])
	return result
}

func newTower(input string) *tower {
	input = strings.TrimSuffix(input, "\n")
	inputLen := len(input)

	var result tower

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

	rockNum := 0
	for {
		rockIndex := rockNum % numShapes
		rock := shapes[rockIndex]
		var turnIndex int

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
				result.heights = append(result.heights, rockTowerHeight)
			}
			turn++
		}

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
				result.cycleHeight = rockTowerHeight - existingStatus.towerHeight
				result.cycleRocks = candidateCycleSize
				fmt.Printf("Found a cycle: %d --> %d = %d (occurred %d times). Delta height = %d\n",
					existingStatus.brickNum, rockNum, candidateCycleSize,
					candidateCycleConsecutiveOccurrences, result.cycleHeight)

				return &result
			}

		}
		fingerprints[fp] = towerStatus{rockNum, rockTowerHeight}
		rockNum++
	}
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

func Part1(input string) int64 {
	tower := newTower(input)
	return tower.HeightAfterRock(2022)
}

func Part2(input string) int64 {
	tower := newTower(input)
	return tower.HeightAfterRock(1000000000000)
}

func main() {
	fmt.Println("*** Advent of Code 2022, day 17 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
