package main

import (
	_ "embed"
	"fmt"
	"github.com/yeurch/advent-of-code/go/ysl"
	"sort"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type NullableInt struct {
	Value int
	Valid bool
}

func (i *NullableInt) Set(v int) {
	i.Valid = true
	i.Value = v
}

type Node struct {
	Value    NullableInt
	Children []*Node
}

func (node *Node) Compare(other *Node) int {
	if node == nil || other == nil {
		panic("Can't compare nil nodes")
	}
	if node.Value.Valid && other.Value.Valid {
		return ysl.Compare(node.Value.Value, other.Value.Value)
	} else if node.Value.Valid && !other.Value.Valid {
		// We have one number and one list of children
		var valAsList Node
		valAsList.Children = []*Node{node}
		return valAsList.Compare(other)
	} else if !node.Value.Valid && other.Value.Valid {
		var valAsList Node
		valAsList.Children = []*Node{other}
		return node.Compare(&valAsList)
	}

	// Both have children
	minLen := len(node.Children)
	if len(other.Children) < minLen {
		minLen = len(other.Children)
	}
	for i := 0; i < minLen; i++ {
		c := node.Children[i].Compare(other.Children[i])
		if c != 0 {
			return c
		}
	}
	return ysl.Sgn(len(node.Children) - len(other.Children))
}

func (node *Node) toString() string {
	if node.Value.Valid {
		return fmt.Sprintf("%d", node.Value.Value)
	}
	children := make([]string, len(node.Children))
	for i, c := range node.Children {
		children[i] = c.toString()
	}
	return fmt.Sprintf("[%s]", strings.Join(children, ", "))
}

type pair struct {
	first  *Node
	second *Node
}

// The following type and associated three functions implement the sort.Interface so we can sort our nodes in part 2
type nodeCollection []*Node

func (nc nodeCollection) Len() int {
	return len(nc)
}

func (nc nodeCollection) Less(i, j int) bool {
	return nc[i].Compare(nc[j]) == -1
}

func (nc nodeCollection) Swap(i, j int) {
	nc[i], nc[j] = nc[j], nc[i]
}

func Part1(input string) int {
	pairs := parseInput(input)
	result := 0
	for i, v := range pairs {
		if v.first.Compare(v.second) < 1 {
			result += i + 1
		}
	}
	return result
}

func Part2(input string) int {
	pairs := parseInput(input)
	// unwrap the pairs0
	nodes := make([]*Node, len(pairs)*2)
	for i, v := range pairs {
		nodes[i*2] = v.first
		nodes[i*2+1] = v.second
	}
	// append our divider packets
	divider1 := parseNode("[[2]]")
	divider2 := parseNode("[[6]]")
	nodes = append(nodes, divider1)
	nodes = append(nodes, divider2)

	sort.Sort(nodeCollection(nodes))

	result := 1
	for i, n := range nodes {
		if divider1.Compare(n) == 0 || divider2.Compare(n) == 0 {
			result *= i + 1
		}
	}

	return result
}

func parseInput(input string) []pair {
	lines := strings.Split(input, "\n")
	numPairs := len(lines) / 3
	result := make([]pair, numPairs)
	for i := 0; i < numPairs; i++ {
		firstLine := i * 3
		result[i] = pair{
			parseNode(lines[firstLine]),
			parseNode(lines[firstLine+1]),
		}
	}
	return result
}

func parseNode(input string) *Node {
	var result Node
	if input == "[]" {
		return &result // empty node
	}
	if v, err := strconv.Atoi(input); err == nil {
		// This is just a numeric node
		result.Value.Set(v)
	} else if input[0] != '[' || input[len(input)-1] != ']' {
		panic("Input can't be parsed: " + input)
	} else {
		s := input[1 : len(input)-1] // strip wrapping pair of []
		children := make([]*Node, 0)
		for len(s) > 0 {
			l := findNodeLen(s)
			children = append(children, parseNode(s[:l]))
			if l == len(s) {
				break
			}
			s = s[l+1:]
		}
		result.Children = children
	}
	return &result
}

func findNodeLen(input string) int {
	if input[0] != '[' {
		// Could be a number, followed by either , or ]
		endPos := strings.IndexAny(input, ",]")
		s := input
		if endPos >= 0 {
			s = s[:endPos]
		}
		_, err := strconv.Atoi(s)
		if err != nil {
			panic("Input doesn't start with [ and isn't a number: " + input)
		}
		return len(s)
	}

	bracketCount := 0
	for i, c := range input {
		if c == '[' {
			bracketCount += 1
		} else if c == ']' {
			bracketCount -= 1
		}
		if bracketCount == 0 {
			return i + 1
		}
	}
	panic("Unterminated sequence: " + input)
}

func main() {
	fmt.Println("*** Advent of Code 2022, day 13 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
