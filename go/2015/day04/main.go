package main

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

func Part1(input string) int {
	return doCalc(input, "00000")
}

func Part2(input string) int {
	return doCalc(input, "000000")
}

func doCalc(input string, prefix string) int {
	result := 0

	i := 0
	for true {
		hash := md(input + strconv.Itoa(i))
		if strings.HasPrefix(hash, prefix) {
			result = i
			break
		}
		i++
	}
	return result
}

func md(str string) string {
	h := md5.New()
	_, _ = io.WriteString(h, str)

	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	fmt.Println("*** Advent of Code 2015, day 04 ***")
	inputDay := "bgvyzdsv"

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
