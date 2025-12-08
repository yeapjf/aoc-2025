package main

import (
	"bufio"
	"cmp"
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	samplePart1, sampelPart2 := solve("sample.txt")
	if samplePart1 != 40 {
		panic("Part 1 incorrect")
	}

	if sampelPart2 != 25272 {
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

	answer1 := 1
	answer2 := 0

	distanceNodes := map[int][]int{}
	for i := range len(lines) {
		for j := i + 1; j < len(lines); j++ {
			distance := getDistance(lines[i], lines[j])
			if _, ok := distanceNodes[distance]; ok {
				panic("Duplicate distances")
			}

			distanceNodes[distance] = append(distanceNodes[distance], i, j)
		}
	}

	keys := slices.Collect(maps.Keys(distanceNodes))
	slices.Sort(keys)

	iterations := 10
	if filename == "input.txt" {
		iterations = 1000
	}

	nodeClusters := map[int]int{}
	clusterCounts := map[int]int{}
	nextCluster := 1

	for i, key := range keys {
		nodes := distanceNodes[key]

		existingCluster := 0
		mergeClusters := []int{}

		for _, node := range nodes {
			if val, ok := nodeClusters[node]; ok && existingCluster == 0 {
				existingCluster = val
			} else if ok && existingCluster != 0 && existingCluster != val {
				mergeClusters = []int{val, existingCluster}
			}
		}

		for _, node := range nodes {
			if existingCluster == 0 {
				nodeClusters[node] = nextCluster
				clusterCounts[nextCluster]++
			} else {
				if _, ok := nodeClusters[node]; !ok {
					nodeClusters[node] = existingCluster
					clusterCounts[existingCluster]++
				}
			}
		}

		if len(mergeClusters) > 0 {
			for key, val := range nodeClusters {
				if val == mergeClusters[1] {
					nodeClusters[key] = mergeClusters[0]
					clusterCounts[mergeClusters[0]]++
				}
			}

			delete(clusterCounts, mergeClusters[1])
		}

		if existingCluster == 0 {
			nextCluster++
		}

		if i+1 == iterations {
			counts := slices.Collect(maps.Values(clusterCounts))
			slices.SortFunc(counts, func(a, b int) int { return cmp.Compare(b, a) })

			for j := range 3 {
				answer1 *= counts[j]
			}
		}

		if len(clusterCounts) == 1 {
			values := slices.Collect(maps.Values(clusterCounts))
			if values[0] == len(lines) {
				cood1 := lines[nodes[0]]
				cood2 := lines[nodes[1]]

				answer2 = getXProduct(cood1, cood2)
				break
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

func getDistance(cood1, cood2 string) int {
	coodArray1 := strings.Split(cood1, ",")
	coodArray2 := strings.Split(cood2, ",")

	distance := 0
	for i := range 3 {
		num1, err := strconv.Atoi(coodArray1[i])
		if err != nil {
			panic(err)
		}

		num2, err := strconv.Atoi(coodArray2[i])
		if err != nil {
			panic(err)
		}

		distance += (num1 - num2) * (num1 - num2)
	}

	return distance
}

func getXProduct(cood1, cood2 string) int {
	coodArray1 := strings.Split(cood1, ",")
	coodArray2 := strings.Split(cood2, ",")

	num1, err := strconv.Atoi(coodArray1[0])
	if err != nil {
		panic(err)
	}

	num2, err := strconv.Atoi(coodArray2[0])
	if err != nil {
		panic(err)
	}

	return num1 * num2
}
