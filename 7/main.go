package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	samplePart1, sampelPart2 := solve("sample.txt")
	if samplePart1 != 21 {
		panic("Part 1 incorrect")
	}

	if sampelPart2 != 40 {
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

	startingIndex := strings.Index(lines[0], "S")
	startingKey := getBeamKey(0, startingIndex)
	beamIndices := map[int]string{
		startingIndex: startingKey,
	}
	beamEdges := map[string][]string{}

	for i, line := range lines {
		// Skip first line
		if i == 1 {
			continue
		}

		if !strings.Contains(line, "^") {
			continue
		}

		for j, char := range line {
			if char == '^' {
				if val, ok := beamIndices[j]; ok {
					if j-1 >= 0 {
						key := getBeamKey(i, j-1)

						// Create vertical edge for overlaps
						if existingKey, ok := beamIndices[j-1]; ok && existingKey != key {
							beamEdges[existingKey] = append(beamEdges[existingKey], key)
						}

						beamIndices[j-1] = key
						beamEdges[val] = append(beamEdges[val], key)
					}

					if j+1 < len(line) {
						key := getBeamKey(i, j+1)

						// Create vertical edge for overlaps
						if existingKey, ok := beamIndices[j+1]; ok && existingKey != key {
							beamEdges[existingKey] = append(beamEdges[existingKey], key)
						}

						beamIndices[j+1] = key
						beamEdges[val] = append(beamEdges[val], key)
					}

					answer1++
					delete(beamIndices, j)
				}
			}
		}
	}

	beamPaths := map[string]int{}
	answer2 = calculateBeamPaths(startingKey, beamEdges, beamPaths)

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

func getBeamKey(lineNum, beamIndex int) string {
	return fmt.Sprintf("%d-%d", lineNum, beamIndex)
}

func calculateBeamPaths(key string, beamEdges map[string][]string, beamPaths map[string]int) int {
	if val, ok := beamPaths[key]; ok {
		return val
	}

	if _, ok := beamEdges[key]; !ok {
		beamPaths[key] = 1
		return 1
	}

	for _, childKey := range beamEdges[key] {
		beamPaths[key] += calculateBeamPaths(childKey, beamEdges, beamPaths)
	}

	return beamPaths[key]
}
