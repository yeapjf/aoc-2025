package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	samplePart1, sampelPart2 := solve("sample.txt")
	if samplePart1 != 4277556 {
		panic("Part 1 incorrect")
	}

	if sampelPart2 != 3263827 {
		panic("Part 2 incorrect")
	}

	part1, part2 := solve("input.txt")
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

// Returns part_1_answer, part_2_answer
func solve(filename string) (int64, int64) {
	lines, err := readLines(filename)
	if err != nil {
		panic(err)
	}

	var answer1 int64 = 0
	var answer2 int64 = 0

	lastIndex := len(lines) - 1
	operators := strings.ReplaceAll(lines[lastIndex], " ", "")
	totals := make([]int64, len(operators))
	re := regexp.MustCompile(" +")

	lines2 := [][]rune{}

	for _, line := range lines[:lastIndex] {
		line2 := []rune(line)
		lines2 = append(lines2, line2)

		line = strings.TrimSpace(line)
		line = re.ReplaceAllString(line, " ")

		numbers := strings.Split(line, " ")
		for i, numStr := range numbers {
			num, err := strconv.ParseInt(numStr, 10, 64)
			if err != nil {
				panic(err)
			}

			operator := operators[i]
			if operator == '+' {
				totals[i] += num
			} else if operator == '*' && totals[i] == 0 {
				totals[i] = num
			} else if operator == '*' {
				totals[i] *= num
			}
		}
	}

	for _, num := range totals {
		answer1 += num
	}

	lines2 = transpose(lines2)
	totals2 := make([]int64, len(operators))
	currentSet := 0

	for _, line := range lines2 {
		numStr := strings.TrimSpace(string(line))
		if numStr == "" {
			currentSet++
			continue
		}

		num, err := strconv.ParseInt(numStr, 10, 64)
		if err != nil {
			panic(err)
		}

		operator := operators[currentSet]
		if operator == '+' {
			totals2[currentSet] += num
		} else if operator == '*' && totals2[currentSet] == 0 {
			totals2[currentSet] = num
		} else if operator == '*' {
			totals2[currentSet] *= num
		}
	}

	for _, num := range totals2 {
		answer2 += num
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

func transpose(matrix [][]rune) [][]rune {
	if len(matrix) == 0 {
		return [][]rune{}
	}

	maxCols := 0
	for _, row := range matrix {
		if len(row) > maxCols {
			maxCols = len(row)
		}
	}

	rows := len(matrix)

	result := make([][]rune, maxCols)
	for i := range result {
		result[i] = make([]rune, 0, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			result[j] = append(result[j], matrix[i][j])
		}
	}

	return result
}
