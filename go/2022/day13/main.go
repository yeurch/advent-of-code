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
	nodes := parseInput(input)
	result := 0
	for i := 0; i < len(nodes)/2; i++ {
		if nodes[i*2].Compare(nodes[i*2+1]) < 1 {
			result += i + 1
		}
	}
	return result
}

func Part2(input string) int {
	nodes := parseInput(input)

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

func parseInput(input string) []*Node {
	lines := strings.Split(input, "\n")
	result := make([]*Node, 0)
	for _, line := range lines {
		if len(line) > 0 {
			result = append(result, parseNode(line))
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
		// This is a number, optionally followed by either , or ]
		endPos := strings.IndexAny(input, ",]")
		s := input
		if endPos >= 0 {
			s = s[:endPos]
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
