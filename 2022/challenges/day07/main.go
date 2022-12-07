package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

func Part1(input string) int {
	totalSizes := getTotalSizes(input)

	result := 0
	for _, v := range totalSizes {
		if v <= 100000 {
			result += v
		}
	}
	return result
}

func Part2(input string) int {
	totalSizes := getTotalSizes(input)

	usedSpace := totalSizes[""]
	maxUsedSpace := 40000000
	requiredDeleteSize := usedSpace - maxUsedSpace

	result := usedSpace
	for _, v := range totalSizes {
		if v >= requiredDeleteSize && v < result {
			result = v
		}
	}
	return result
}

func getTotalSizes(input string) map[string]int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	currentDir := ""
	directorySizes := make(map[string]int)

	for _, line := range lines {
		if line[0] == '$' {
			// command
			if !strings.HasPrefix(line, "$ cd ") {
				continue // we only care about cd commands
			}
			currentDir = cd(currentDir, line[5:])
			_, ok := directorySizes[currentDir]
			if !ok {
				directorySizes[currentDir] = 0
			}
		} else {
			// data
			if strings.HasPrefix(line, "dir ") {
				continue // we only care about files, not subdirs
			}
			var fileSize int
			_, _ = fmt.Sscan(line, &fileSize)

			directorySizes[currentDir] += fileSize
		}
	}

	totalSizes := make(map[string]int)
	for k, v := range directorySizes {
		totalSize := v
		for k2, v2 := range directorySizes {
			if strings.HasPrefix(k2, k+"/") {
				totalSize += v2
			}
		}
		totalSizes[k] = totalSize
	}
	return totalSizes
}

func cd(current, new string) string {
	if new == "/" {
		return ""
	}
	if new == ".." {
		return current[0:strings.LastIndex(current, "/")]
	} else {
		return current + "/" + new
	}
}

func main() {
	fmt.Println("*** Advent of Code 2022, day 07 ***")

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
