package main

import (
	_ "embed"
	"fmt"
	"github.com/yeurch/advent-of-code/go/ysl"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type monkey struct {
	name    string
	args    []string
	bSolved bool
	result  int64
}

func Part1(input string) int64 {
	monkeys := parseInput(input)
	for !monkeys["root"].bSolved {
		for _, monkey := range monkeys {
			if !monkey.bSolved {
				if monkeys[monkey.args[0]].bSolved && monkeys[monkey.args[2]].bSolved {
					op1 := monkeys[monkey.args[0]].result
					op2 := monkeys[monkey.args[2]].result
					switch monkey.args[1] {
					case "*":
						monkey.result = op1 * op2
					case "+":
						monkey.result = op1 + op2
					case "/":
						monkey.result = op1 / op2
					case "-":
						monkey.result = op1 - op2
					}
					monkey.bSolved = true
				}
			}
		}
	}
	return monkeys["root"].result
}

func Part2(input string) int {
	monkeys := parseInput(input)
	monkeys["humn"].bSolved = false
	bFound := true
	for bFound {
		bFound = false
		for _, monkey := range monkeys {
			if !monkey.bSolved && monkey.name != "root" && monkey.name != "humn" {
				if monkeys[monkey.args[0]].bSolved && monkeys[monkey.args[2]].bSolved {
					op1 := monkeys[monkey.args[0]].result
					op2 := monkeys[monkey.args[2]].result
					switch monkey.args[1] {
					case "*":
						monkey.result = op1 * op2
					case "+":
						monkey.result = op1 + op2
					case "/":
						monkey.result = op1 / op2
					case "-":
						monkey.result = op1 - op2
					}
					monkey.bSolved = true
					bFound = true
				}
			}
		}
	}

	unsolvedCount := 0
	for _, v := range monkeys {
		if !v.bSolved {
			unsolvedCount++
		}
	}
	fmt.Printf("Unsolved: %d/%d\n", unsolvedCount, len(monkeys))
	fmt.Println(monkeys["root"])
	fmt.Println(monkeys["humn"])

	// Let's build up a list of monkeys we still care about and purge the rest
	monkeysWeCareAbout := ysl.NewSet[string]()
	monkeysWeCareAbout.Add("humn")
	for _, m := range monkeys {
		if !m.bSolved && m.name != "humn" {
			monkeysWeCareAbout.Add(m.args[0])
			monkeysWeCareAbout.Add(m.args[2])
			monkeysWeCareAbout.Add(m.name)
		}
	}
	fmt.Printf("We still care about %d monkeys.\n", monkeysWeCareAbout.Cardinality())

	monkeys["root"].args[1] = "="

	for i := 0; i < math.MaxInt; i++ {
		if i%1000 == 0 {
			fmt.Printf("Attempt %d\n", i)
		}

		// making a copy of our state
		monkeyClones := make(map[string]*monkey)

		for k := range monkeys {
			if monkeysWeCareAbout.Contains(k) {
				monkeyClone := *(monkeys[k])
				monkeyClones[k] = &monkeyClone
			}
		}

		human := monkeyClones["humn"]
		human.bSolved = true
		human.result = int64(i)

		bFound := true
		for bFound {
			bFound = false
			for _, monkey := range monkeyClones {
				if !monkey.bSolved {
					if monkeyClones[monkey.args[0]].bSolved && monkeyClones[monkey.args[2]].bSolved {
						op1 := monkeyClones[monkey.args[0]].result
						op2 := monkeyClones[monkey.args[2]].result
						bExit := false
						switch monkey.args[1] {
						case "*":
							monkey.result = op1 * op2
							break
						case "+":
							monkey.result = op1 + op2
							break
						case "/":
							monkey.result = op1 / op2
							break
						case "-":
							monkey.result = op1 - op2
							break
						case "=":
							if op1 == op2 {
								return i
							} else {
								bFound = false
								bExit = true
								break
							}
						}
						if bExit {
							break
						}
						monkey.bSolved = true
						bFound = true
					}
				}
			}
		}
	}

	panic("We should not reach here")
}

func parseInput(input string) map[string]*monkey {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	result := make(map[string]*monkey)
	for _, line := range lines {
		name := line[0:4]
		args := strings.Split(line[6:], " ")
		if len(args) == 1 {
			val, _ := strconv.Atoi(args[0])
			result[name] = &monkey{name, args, true, int64(val)}
		} else {
			result[name] = &monkey{name, args, false, 0}
		}
	}
	return result
}

func main() {
	fmt.Println("*** Advent of Code 2022, day 21 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
