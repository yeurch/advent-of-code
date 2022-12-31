package main

import (
	"container/list"
	_ "embed"
	"fmt"
	"github.com/yeurch/advent-of-code/go/ysl"
	"math"
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

type GameState struct {
	*Point
	n int
}
type Blizzard struct {
	*Point
	direction Point // used as a 2D vector
}

type Valley struct {
	entry  Point
	exit   Point
	width  int
	height int

	spacesCache []ysl.Set[Point]
	blizzards   []Blizzard
}

func (v *Valley) GetSpacesAfterN(n int) ysl.Set[Point] {
	if len(v.spacesCache) > n {
		return v.spacesCache[n]
	}

	for i := len(v.spacesCache); i <= n; i++ {
		v.spacesCache = append(v.spacesCache, v.SimulateMove())
	}
	return v.spacesCache[n]
}

// Returns a set of free spaces after simulating another move
func (v *Valley) SimulateMove() ysl.Set[Point] {
	for i := 0; i < len(v.blizzards); i++ {
		b := &v.blizzards[i]
		b.x += b.direction.x
		b.y += b.direction.y
		if b.x < 0 {
			b.x = v.width - 1
		} else if b.x >= v.width {
			b.x = 0
		}
		if b.y < 0 {
			b.y = v.height - 1
		} else if b.y >= v.height {
			b.y = 0
		}
	}

	spaces := ysl.NewSet[Point]()
	for y := 0; y < v.height; y++ {
		for x := 0; x < v.width; x++ {
			spaces.Add(Point{x, y})
		}
	}
	for i := 0; i < len(v.blizzards); i++ {
		b := &v.blizzards[i]
		spaces.Remove(Point{b.x, b.y})
	}
	return spaces
}

func Part1(input string) int {
	valley := parseInput(input)
	return solve(&valley, valley.entry, valley.exit, 0)
}

func Part2(input string) int {
	valley := parseInput(input)
	timeTaken := solve(&valley, valley.entry, valley.exit, 0)
	timeTaken = solve(&valley, valley.exit, valley.entry, timeTaken)
	timeTaken = solve(&valley, valley.entry, valley.exit, timeTaken)
	return timeTaken
}

func solve(valley *Valley, start Point, end Point, startTimeOffset int) int {
	routeQueue := list.New()
	routeQueue.PushBack(GameState{&start, startTimeOffset})
	//visited := ysl.NewSet[GameState]()
	visited := ysl.NewSet[string]()

	bestPathLen := math.MaxInt
	for el := routeQueue.Front(); el != nil; el = el.Next() {
		gameState := el.Value.(GameState)

		//dupeDetectionState := GameState{gameState.Point, gameState.n % (valley.width * valley.height)}
		dupeDetectionState := fmt.Sprintf("%d:%d:%d", gameState.x, gameState.y, gameState.n%(valley.width*valley.height))
		if visited.Contains(dupeDetectionState) {
			continue
		}
		visited.Add(dupeDetectionState)

		if gameState.n >= bestPathLen {
			continue // We already have a better solution
		}

		if *gameState.Point == end {
			if gameState.n < bestPathLen {
				bestPathLen = gameState.n
			}
			continue
		}

		spaces := valley.GetSpacesAfterN(gameState.n + 1)
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				if ysl.Abs(dy)+ysl.Abs(dx) == 2 {
					continue // eliminate diagonal moves
				}
				candidatePos := gameState.Point.Add(Point{dx, dy})
				if spaces.Contains(candidatePos) || candidatePos == valley.exit || candidatePos == valley.entry {
					routeQueue.PushBack(GameState{&candidatePos, gameState.n + 1})
				}
			}
		}
	}
	return bestPathLen
}

func parseInput(input string) Valley {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	var entry, exit Point
	entryRow := lines[0]
	width := len(entryRow) - 2
	height := len(lines) - 2
	for x := 0; x < width; x++ {
		if entryRow[x+1] == '.' {
			entry = Point{x, len(lines) - 2}
			break
		}
	}
	exitRow := lines[len(lines)-1]
	for x := 0; x < width; x++ {
		if exitRow[x+1] == '.' {
			exit = Point{x, -1}
			break
		}
	}

	directions := map[uint8]Point{
		'>': {1, 0},
		'<': {-1, 0},
		'^': {0, 1},
		'v': {0, -1},
	}

	blizzards := make([]Blizzard, 0)
	for i := 1; i < len(lines)-1; i++ {
		line := lines[i]
		y := len(lines) - i - 2
		for j := 1; j < len(line)-1; j++ {
			if line[j] == '.' {
				continue
			}
			x := j - 1
			blizVec := directions[line[j]]
			blizzards = append(blizzards, Blizzard{&Point{x, y}, blizVec})
		}
	}

	spaces := ysl.NewSet[Point]()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			spaces.Add(Point{x, y})
		}
	}
	for i := 0; i < len(blizzards); i++ {
		b := blizzards[i]
		spaces.Remove(Point{b.x, b.y})
	}
	spacesCache := []ysl.Set[Point]{spaces}
	return Valley{
		entry, exit, width, height, spacesCache, blizzards,
	}
}

func main() {
	fmt.Println("*** Advent of Code 2022, day 24 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
