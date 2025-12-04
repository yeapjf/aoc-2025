package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	samplePart1, sampelPart2 := solve("sample.txt")
	if samplePart1 != 13 {
		panic("Part 1 incorrect")
	}

	if sampelPart2 != 43 {
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
	loops := 0

	for {
		newLines := []string{}
		changed := false

		for i, line := range lines {
			newLine := line

			for j, char := range line {
				if char != '@' {
					continue
				}

				neighbours := countNeighbours(lines, i, j)
				if neighbours < 4 {
					if loops == 0 {
						answer1++
					}

					answer2++
					newLine = replaceChar(newLine, j)
					changed = true
				}
			}

			newLines = append(newLines, newLine)
		}

		if !changed {
			break
		}

		lines = newLines
		loops++
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

func countNeighbours(lines []string, i, j int) int {
	total := 0
	maxRow := len(lines) - 1
	maxCol := len(lines[0]) - 1

	for row := i - 1; row < i+2; row++ {
		if row < 0 || row > maxRow {
			continue
		}

		for col := j - 1; col < j+2; col++ {
			if col < 0 || col > maxCol {
				continue
			}

			if row == i && col == j {
				continue
			}

			if lines[row][col] == '@' {
				total++
			}
		}
	}

	return total
}

func replaceChar(line string, position int) string {
	out := []rune(line)
	out[position] = '.'
	return string(out)
}
