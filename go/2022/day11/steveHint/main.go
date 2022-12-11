/******************************************************************************
 * The solution in this file is simplified compared to the solution in the
 * parent directory, thanks to a hint from Steve Robb.
 * His suggestion for part 2  was that instead of maintaining each item's value
 * for ALL monkeys, applying the modulo of each monkey to every operation, we
 * can go back to only keeping one value (as per part 1), but applying the
 * modulo of the products of all monkeys' divisors.
 *******************************************************************************/

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

var moduloProduct int

type monkey struct {
	items           *list.List
	operation       []string
	divisor         int
	trueTarget      int
	falseTarget     int
	inspectionCount int
}

func newMonkey(items []int, op []string, divisor, trueTarget, falseTarget int) monkey {
	l := list.New()
	for _, i := range items {
		l.PushBack(i)
	}
	return monkey{l, op, divisor, trueTarget, falseTarget, 0}
}

func (m *monkey) do(allMonkeys []monkey, part2 bool) {
	for el := m.items.Front(); el != nil; el = el.Next() {
		val := el.Value.(int)

		val = m.op(val)
		if !part2 {
			val = val / 3
		}
		val = val % moduloProduct

		target := m.falseTarget
		if val%m.divisor == 0 {
			target = m.trueTarget
		}

		allMonkeys[target].giveItem(val)
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

func (m *monkey) giveItem(item int) {
	m.items.PushBack(item)
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

	monkeys := make([]monkey, n)
	moduloProduct = 1
	for i := 0; i < n; i++ {
		start := i * 7

		itemStrings := strings.Split(lines[start+1][18:], ", ")
		items := make([]int, len(itemStrings))
		for j, itemString := range itemStrings {
			items[j], _ = strconv.Atoi(itemString)
		}

		ops := strings.Split(lines[start+2][19:], " ")

		divisor, _ := strconv.Atoi(lines[start+3][21:])
		moduloProduct *= divisor

		trueTarget, _ := strconv.Atoi(lines[start+4][29:])
		falseTarget, _ := strconv.Atoi(lines[start+5][30:])

		monkeys[i] = newMonkey(items, ops, divisor, trueTarget, falseTarget)
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
