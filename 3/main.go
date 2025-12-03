package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	samplePart1, sampelPart2 := solve("sample.txt")
	if samplePart1 != 357 {
		panic("Part 1 incorrect")
	}

	if sampelPart2 != 3121910778619 {
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
		lineLen := len(line)
		firstIndex, firstDigit := largestDigitInRange(line, 0, lineLen-1)
		_, secondDigit := largestDigitInRange(line, firstIndex+1, lineLen)
		answer1 += firstDigit*10 + secondDigit

		digitIndices := []int{0}
		for i := range 12 {
			digitsLeft := 12 - i
			startingIndex := digitIndices[i]

			if startingIndex+digitsLeft == lineLen {
				digitStr := line[startingIndex]
				digit, err := strconv.Atoi(string(digitStr))
				if err != nil {
					panic(err)
				}

				digitIndices = append(digitIndices, startingIndex+1)
				answer2 += digit * int(math.Pow(10, float64(11-i)))
				continue
			}

			digitIndex, digit := largestDigitInRange(line, startingIndex, lineLen-11+i)
			digitIndices = append(digitIndices, startingIndex+digitIndex+1)
			answer2 += digit * int(math.Pow(10, float64(11-i)))
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

func largestDigitInRange(numStr string, startIndex, endIndex int) (int, int) {
	largestIndex := 0
	largestDigit := 0

	for i, digitStr := range numStr[startIndex:endIndex] {
		digit, err := strconv.Atoi(string(digitStr))
		if err != nil {
			panic(err)
		}

		if digit > largestDigit {
			largestIndex = i
			largestDigit = digit
		}
	}

	return largestIndex, largestDigit
}
