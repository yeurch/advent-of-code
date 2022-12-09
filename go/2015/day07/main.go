package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type operation struct {
	opcode   string
	operands []string
	solved   bool
	solution uint16
}

func newOperation(opcode string, operands []string) *operation {
	return &operation{opcode, operands, false, 0}
}

func (o *operation) solve(ops *map[string]*operation) uint16 {
	if o.solved {
		return o.solution
	}

	opValues := make([]uint16, len(o.operands))
	for i, operand := range o.operands {
		val32, err := strconv.Atoi(operand)
		var val uint16
		if err != nil {
			// Value is not numeric (i.e. it's another variable)
			o2 := (*ops)[operand]
			val = o2.solve(ops) // recursively solve our operand
		} else {
			val = uint16(val32)
		}
		opValues[i] = val
	}

	switch o.opcode {
	case "AND":
		o.solution = opValues[0] & opValues[1]
	case "OR":
		o.solution = opValues[0] | opValues[1]
	case "LSHIFT":
		o.solution = opValues[0] << opValues[1]
	case "RSHIFT":
		o.solution = opValues[0] >> opValues[1]
	case "NOT":
		o.solution = ^opValues[0]
	case "NOOP":
		o.solution = opValues[0]
	}
	o.solved = true
	return o.solution
}

func Part1(input string) uint16 {
	ops := parseInput(input)
	return ops["a"].solve(&ops)
}

func Part2(input string, part1Result uint16) uint16 {
	ops := parseInput(input)
	ops["b"] = newOperation("NOOP", []string{strconv.Itoa(int(part1Result))})
	return ops["a"].solve(&ops)
}

func parseInput(input string) map[string]*operation {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	r := regexp.MustCompile("([a-z0-9]+ )?(\\w+) (\\w+) -> (\\w+)")
	rNoop := regexp.MustCompile("(\\w+) -> (\\w+)")

	result := make(map[string]*operation)

	for _, line := range lines {
		m := r.FindStringSubmatch(line)

		if m == nil {
			// a "no-op" ... either a literal value or just a copy of a single variable
			mNoop := rNoop.FindStringSubmatch(line)
			if mNoop == nil {
				panic("Unexpected input: " + line)
			}
			target := mNoop[2]
			operands := []string{mNoop[1]}
			result[target] = newOperation("NOOP", operands)
			continue
		}

		opcode := m[2]
		operands := make([]string, 0)
		if len(m[1]) > 0 {
			operands = append(operands, strings.TrimSpace(m[1]))
		}
		operands = append(operands, m[3])
		target := m[4]
		result[target] = newOperation(opcode, operands)
	}
	return result
}

func main() {
	fmt.Println("*** Advent of Code 2015, day 07 ***")

	start := time.Now()
	part1Result := Part1(inputDay)
	fmt.Println("part1: ", part1Result)
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay, part1Result))
	fmt.Println(time.Since(start))
}
