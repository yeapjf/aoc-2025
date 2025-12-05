package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	samplePart1, sampelPart2 := solve("sample.txt")
	if samplePart1 != 3 {
		panic("Part 1 incorrect")
	}

	if sampelPart2 != 14 {
		panic("Part 2 incorrect")
	}

	part1, part2 := solve("input.txt")
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

// Returns part_1_answer, part_2_answer
func solve(filename string) (int, int64) {
	lines, err := readLines(filename)
	if err != nil {
		panic(err)
	}

	answer1 := 0
	var answer2 int64 = 0

	freshFruitPairs := [][]int64{}
	freshRanges := [][]int64{}
	toSkip := map[int]bool{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.Contains(line, "-") {
			pair := strings.Split(line, "-")
			startNum, err := strconv.ParseInt(pair[0], 10, 64)
			if err != nil {
				panic(err)
			}

			endNum, err := strconv.ParseInt(pair[1], 10, 64)
			if err != nil {
				panic(err)
			}

			newRange := []int64{startNum, endNum}
			freshFruitPairs = append(freshFruitPairs, newRange)

			shouldAdd := len(freshRanges) == 0

			for i, r := range freshRanges {
				if newRange[0] >= r[0] && newRange[1] <= r[1] {
					// Subset, throw away current range
					shouldAdd = false
					break
				} else if newRange[1] < r[0] || newRange[0] > r[1] {
					// No overlap, consider next line
					shouldAdd = true
				} else if newRange[0] <= r[0] && newRange[1] >= r[1] {
					// Superset, throw away comparison range
					shouldAdd = true
					toSkip[i] = true
				} else if newRange[0] >= r[0] && newRange[1] > r[1] {
					// Last number bigger, adjust current range
					shouldAdd = true
					newRange[0] = r[1] + 1
				} else if newRange[0] < r[0] && newRange[1] <= r[1] {
					// First number smaller, adjust current range
					shouldAdd = true
					newRange[1] = r[0] - 1
				}
			}

			if shouldAdd {
				freshRanges = append(freshRanges, newRange)
			}
		} else {
			fruitNum, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				panic(err)
			}

			for _, pair := range freshFruitPairs {
				if fruitNum < pair[0] {
					continue
				}

				if fruitNum > pair[1] {
					continue
				}

				answer1++
				break
			}
		}
	}

	for i, r := range freshRanges {
		if _, ok := toSkip[i]; ok {
			continue
		}

		answer2 += r[1] - r[0] + 1
	}

	return answer1, answer2
}

// Returns all lines in the file as a string slice
func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
