/**********************************************************
 * No "real" code today for me. Just eyeballed both stars
 * in a text editor.
 */

package main

import (
	_ "embed"
	"fmt"
	"time"
)

func Part1() int64 {
	return 380
}

func Part2() int64 {
	return 375
}

func main() {
	fmt.Println("*** Advent of Code 2022, day 12 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1())
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2())
	fmt.Println(time.Since(start))
}
