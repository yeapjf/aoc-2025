package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	samplePart1, sampelPart2 := solve("sample.txt")
	if samplePart1 != 3 {
		panic("Part 1 incorrect")
	}

	if sampelPart2 != 6 {
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

	position := 50
	re := regexp.MustCompile("(L|R)([0-9]+)")

	for _, line := range lines {
		match := re.FindStringSubmatch(string(line))

		rotation, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}

		loops := rotation / 100
		change := rotation % 100

		if match[1] == "L" {
			change *= -1
		}

		startingPosition := position
		position += change

		if position < 0 {
			position += 100

			if startingPosition != 0 && position != 0 {
				loops++
			}
		} else if position > 99 {
			position -= 100

			if position != 0 {
				loops++
			}
		}

		answer2 += loops

		if position == 0 {
			answer1++

			if change != 0 {
				answer2++
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
