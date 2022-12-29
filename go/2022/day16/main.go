package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type location struct {
	code     string
	flowRate int
	isOpen   bool
	exits    []route
}

type route struct {
	dest   string
	length int
}

func Part1(input string) int {
	_, volcano := parseInput(input)
	dumpVolcano(&volcano)

	distances := getDistances(&volcano)
	dumpDistances(&distances)

	var result int
	visit(&volcano, "AA", 0, 30, make([]route, 0), 0, 0, &result)
	return result
}

func Part2(input string) int {
	_, volcano := parseInput(input)
	dumpVolcano(&volcano)

	distances := getDistances(&volcano)
	dumpDistances(&distances)

	var result int
	visit(&volcano, "AA", 0, 30, make([]route, 0), 0, 0, &result)
	return result
}

func visit(volcano *map[string]*location, current string, elapsed int, maxTime int, history []route, pressure int, score int, maxScore *int) {
	var enteredFrom route
	if len(history) > 0 {
		enteredFrom = history[len(history)-1]
	}
	here := (*volcano)[current]

	// enter
	timeElapsed := enteredFrom.length
	deltaScore := timeElapsed * pressure
	score += deltaScore
	elapsed += timeElapsed
	if enteredFrom.dest == here.code {
		// Entering from ourself signifies activating valve
		here.isOpen = true
		pressure += here.flowRate
	}

	// next action
	timeLeft := maxTime - elapsed
	if !here.isOpen && here.flowRate > 0 && timeLeft > 0 {
		history = append(history, route{here.code, 1})
		visit(volcano, here.code, elapsed, maxTime, history, pressure, score, maxScore)
		history = history[:len(history)-1]
	}
	for _, dest := range here.exits {
		if dest.dest != enteredFrom.dest && dest.length <= timeLeft {
			history = append(history, route{here.code, dest.length})
			visit(volcano, dest.dest, elapsed, maxTime, history, pressure, score, maxScore)
			history = history[:len(history)-1]
		}
	}

	// Perhaps we have time left, but not enough to do anything. We should see what our final score would be
	score += timeLeft * pressure
	if score > *maxScore {
		*maxScore = score
	}

	// back out
	if enteredFrom.dest == here.code {
		// Entering from ourself signifies activating valve so we need to deactivate
		here.isOpen = false // this might be the only thing we care about on exit?
	}
}

func parseInput(input string) ([]string, map[string]*location) {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	result := make(map[string]*location)
	keys := make([]string, len(lines))

	// Valve GG has flow rate=0; tunnels lead to valves FF, HH
	// Valve HH has flow rate=22; tunnel leads to valve GG
	r := regexp.MustCompile("Valve (\\w+) has flow rate=([0-9]+); tunnels? leads? to valves? (.+)")
	for i, line := range lines {
		m := r.FindStringSubmatch(line)
		name := m[1]
		keys[i] = name
		flowRate, _ := strconv.Atoi(m[2])
		destinations := strings.Split(m[3], ", ")
		routes := make([]route, len(destinations))
		for i, d := range destinations {
			routes[i] = route{d, 1}
		}
		result[name] = &location{name, flowRate, false, routes}
	}

	simplify(&result)
	return keys, result
}

func getDistances(volcano *map[string]*location) map[string]*map[string]int {
	result := make(map[string]*map[string]int)

	allKeys := make([]string, 0)

	for k, v := range *volcano {
		m := make(map[string]int)
		for _, n := range v.exits {
			m[n.dest] = n.length
		}
		result[k] = &m
		allKeys = append(allKeys, k)
	}

	bFoundRoute := true
	for bFoundRoute {
		bFoundRoute = false
		for i := 0; i < len(allKeys); i++ {
			mapI := result[allKeys[i]]
			for j := i + 1; j < len(allKeys); j++ {
				mapJ := result[allKeys[j]]
				currentBest, ok := (*mapI)[allKeys[j]]
				if !ok {
					currentBest = math.MaxInt
				}
				// try to improve route from i to j
				for _, loc := range allKeys {
					iToLoc, ok := (*mapI)[loc]
					if ok {
						jToLoc, ok := (*mapJ)[loc]
						if ok {
							dist := iToLoc + jToLoc
							if dist < currentBest {
								(*mapI)[allKeys[j]] = dist
								(*mapJ)[allKeys[i]] = dist
								bFoundRoute = true
							}
						}
					}
				}
			}
		}
	}

	// strip out destinations with no valve
	for _, v := range result {
		for k := range *v {
			if (*volcano)[k].flowRate == 0 {
				delete(*v, k)
			}
		}
	}

	// add one second to each time to open valve
	for _, v := range result {
		for k, dist := range *v {
			(*v)[k] = dist + 1
		}
	}
	return result
}

func simplify(volcano *map[string]*location) {
	bModified := true
	for bModified {
		bModified = false
		for _, loc := range *volcano {
			if loc.flowRate == 0 && len(loc.exits) == 2 {
				// This is a straight tunnel section and can be removed
				newLen := loc.exits[0].length + loc.exits[1].length
				for i := 0; i < 2; i++ {
					exits := (*volcano)[loc.exits[i].dest].exits
					for j := range exits {
						e := &exits[j]
						if e.dest == loc.code {
							e.dest = loc.exits[1-i].dest
							e.length = newLen
							break
						}
					}
				}
				delete(*volcano, loc.code)
				bModified = true
			}
		}
	}
}

func dumpVolcano(volcano *map[string]*location) {
	for _, v := range *volcano {
		fmt.Println(v)
	}
	fmt.Println()
}

func dumpDistances(distances *map[string]*map[string]int) {
	for k, v := range *distances {
		fmt.Printf("%s: ", k)
		for k2, v2 := range *v {
			fmt.Printf("%s=%d; ", k2, v2)
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	fmt.Println("*** Advent of Code 2022, day 16 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
