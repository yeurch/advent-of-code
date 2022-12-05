package main

import (
	"container/list"
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func Part1(input string) string {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	moveRegexp := regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")

	crates, moveLines := doSetup(lines)
	for _, line := range moveLines {
		if moveRegexp.MatchString(line) {
			m := moveRegexp.FindStringSubmatch(line)
			qty, _ := strconv.Atoi(m[1])
			from, _ := strconv.Atoi(m[2])
			to, _ := strconv.Atoi(m[3])
			doMove(crates, qty, from, to, false)
		}
	}

	return getResult(crates)
}

func Part2(input string) string {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	moveRegexp := regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")

	crates, moveLines := doSetup(lines)
	for _, line := range moveLines {
		if moveRegexp.MatchString(line) {
			m := moveRegexp.FindStringSubmatch(line)
			qty, _ := strconv.Atoi(m[1])
			from, _ := strconv.Atoi(m[2])
			to, _ := strconv.Atoi(m[3])
			doMove(crates, qty, from, to, true)
		}
	}

	return getResult(crates)
}

func doSetup(lines []string) (map[int]*list.List, []string) {
	crates := make(map[int]*list.List)
	count := 0
	for _, line := range lines {
		count += 1
		if line == "" || !strings.Contains(line, "[") {
			//debugPrint(crates)
			break
		}
		// This is a setup line
		index := 0
		for i := 1; i < len(line); i += 4 {
			index += 1
			if line[i] == ' ' {
				continue
			}
			stack, ok := crates[index]
			if !ok {
				stack = list.New()
				crates[index] = stack
			}
			stack.PushBack(line[i])
		}
	}
	return crates, lines[count:]
}

func doMove(crates map[int]*list.List, qty int, from int, to int, part2 bool) {
	var src *list.List
	if part2 {
		tmp := list.New()
		for i := 0; i < qty; i++ {
			moveCrate(crates[from], tmp)
		}
		src = tmp
	} else {
		src = crates[from]
	}

	for i := 0; i < qty; i++ {
		moveCrate(src, crates[to])
	}
}

func moveCrate(from *list.List, to *list.List) {
	box := from.Front()
	val := from.Remove(box)
	to.PushFront(val)
}

func getResult(crates map[int]*list.List) string {
	var result []byte
	for i := 1; ; i++ {
		stack, ok := crates[i]
		if !ok {
			break
		}

		val := stack.Front().Value
		result = append(result, val.(byte))
	}
	return string(result)
}

func main() {
	fmt.Println("*** Advent of Code 2022, day 05 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
