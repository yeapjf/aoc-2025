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
	if samplePart1 != 1227775554 {
		panic("Part 1 incorrect")
	}

	if sampelPart2 != 4174379265 {
		panic("Part 2 incorrect")
	}

	part1, part2 := solve("input.txt")
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

// Returns part_1_answer, part_2_answer
func solve(filename string) (int, int) {
	lines, err := readLines(filename)
	if err != nil {
		panic(err)
	}

	answer1 := 0
	answer2 := 0

	for _, line := range lines {
		for _, pair := range strings.Split(line, ",") {
			pairStr := strings.Split(pair, "-")

			start, err := strconv.Atoi(string(pairStr[0]))
			if err != nil {
				panic(err)
			}

			end, err := strconv.Atoi(string(pairStr[1]))
			if err != nil {
				panic(err)
			}

			for num := start; num <= end; num++ {
				if isInvalid(num) {
					answer1 += num
					answer2 += num
				} else if isRepeating(num) {
					answer2 += num
				}
			}
		}
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

func isInvalid(num int) bool {
	numStr := strconv.Itoa(num)
	numLen := len(numStr)

	if numLen%2 == 1 {
		return false
	}

	midIndex := numLen / 2
	left := numStr[:midIndex]
	right := numStr[midIndex:]

	if left == right {
		return true
	}

	return false
}

func isRepeating(num int) bool {
	numStr := strconv.Itoa(num)
	numLen := len(numStr)

	for i := 1; i <= numLen/2; i++ {
		if numLen%i != 0 {
			continue
		}

		prefix := numStr[:i]
		isRepeating := true

		for startIndex := i; startIndex < numLen; startIndex += i {
			chunk := numStr[startIndex : startIndex+i]
			if chunk != prefix {
				isRepeating = false
				break
			}
		}

		if isRepeating {
			return true
		}
	}

	return false
}
