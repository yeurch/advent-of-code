package main

import (
	"container/list"
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type monkey struct {
	items           *list.List
	operation       []string
	divisor         int
	trueTarget      int
	falseTarget     int
	inspectionCount int
}

func newMonkey(items []int, allDivisors []int, op []string, divisor, trueTarget, falseTarget int) monkey {
	l := list.New()
	for _, i := range items {
		l.PushBack(i)
		m := make(map[int]int)
		for _, d := range allDivisors {
			m[d] = i % d
		}
		l.PushBack(m)
	}
	return monkey{l, op, divisor, trueTarget, falseTarget, 0}
}

func (m *monkey) do(allMonkeys []monkey, part2 bool) {
	for el := m.items.Front(); el != nil; el = el.Next() {
		val := el.Value.(int)
		el = el.Next()
		divisorMap := el.Value.(map[int]int)

		target := m.falseTarget
		if part2 {
			for k, v := range divisorMap {
				newVal := m.op(v)
				divisorMap[k] = newVal % k
			}
			if divisorMap[m.divisor] == 0 {
				target = m.trueTarget
			}
		} else {
			val = m.op(val) / 3
			if val%m.divisor == 0 {
				target = m.trueTarget
			}
		}

		allMonkeys[target].giveItem(val, divisorMap)
		m.inspectionCount++
	}
	m.items = list.New()
}

func (m *monkey) op(val int) int {
	var op1, op2 int
	if m.operation[0] == "old" {
		op1 = val
	} else {
		op1, _ = strconv.Atoi(m.operation[0])
	}
	if m.operation[2] == "old" {
		op2 = val
	} else {
		op2, _ = strconv.Atoi(m.operation[2])
	}

	if m.operation[1] == "*" {
		return op1 * op2
	} else if m.operation[1] == "+" {
		return op1 + op2
	} else {
		panic("Unexpected operation: " + m.operation[2])
	}
}

func (m *monkey) giveItem(item int, divisorMap map[int]int) {
	m.items.PushBack(item)
	m.items.PushBack(divisorMap)
}

func Part1(input string) int64 {
	allMonkeys := parseMonkeys(input)

	for roundNum := 0; roundNum < 20; roundNum++ {
		for i := range allMonkeys {
			monkey := &allMonkeys[i]
			monkey.do(allMonkeys, false)
		}
	}

	return calcMonkeyBusiness(allMonkeys)
}

func Part2(input string) int64 {
	allMonkeys := parseMonkeys(input)

	for roundNum := 0; roundNum < 10000; roundNum++ {
		for i := range allMonkeys {
			monkey := &allMonkeys[i]
			monkey.do(allMonkeys, true)
		}
	}

	return calcMonkeyBusiness(allMonkeys)
}

func calcMonkeyBusiness(allMonkeys []monkey) int64 {
	n := len(allMonkeys)
	inspectionCounts := make([]int, n)
	for i, m := range allMonkeys {
		inspectionCounts[i] = m.inspectionCount
	}
	sort.Ints(inspectionCounts)
	return int64(inspectionCounts[n-1]) * int64(inspectionCounts[n-2])
}

func parseMonkeys(input string) []monkey {
	lines := strings.Split(input, "\n")
	n := len(lines) / 7

	// First we'll get all of the divisors, we'll need each monkey to track them all
	allDivisors := make([]int, n)
	for i := 0; i < n; i++ {
		divisorLine := i*7 + 3
		divisor, _ := strconv.Atoi(lines[divisorLine][21:])
		allDivisors[i] = divisor
	}

	monkeys := make([]monkey, n)
	for i := 0; i < n; i++ {
		start := i * 7

		itemStrings := strings.Split(lines[start+1][18:], ", ")
		items := make([]int, len(itemStrings))
		for j, itemString := range itemStrings {
			items[j], _ = strconv.Atoi(itemString)
		}

		ops := strings.Split(lines[start+2][19:], " ")

		divisor, _ := strconv.Atoi(lines[start+3][21:])
		trueTarget, _ := strconv.Atoi(lines[start+4][29:])
		falseTarget, _ := strconv.Atoi(lines[start+5][30:])

		monkeys[i] = newMonkey(items, allDivisors, ops, divisor, trueTarget, falseTarget)
	}
	return monkeys
}

func main() {
	fmt.Println("*** Advent of Code 2022, day 11 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
