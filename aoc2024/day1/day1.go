package day1

import (
	"aoc2024/util"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func Day1() {
	//inputLines := util.ReadInputFileToStringSlice("day1/sample")
	inputLines := util.ReadInputFileToStringSlice("day1/input")

	leftList, rightList := splitInputLinesToIntLists(inputLines)

	slices.Sort(leftList)
	slices.Sort(rightList)

	totalDistance, _ := calculateTotalDistance(leftList, rightList)

	totalSimilarity := calculateTotalSimilarityScore(leftList, rightList)

	println(">>>>>TOTAL DISTANCE: %v", totalDistance)
	println(">>>>>TOTAL SIMILARITY: %v", totalSimilarity)
}

func splitInputLinesToIntLists(lines []string) ([]int, []int) {
	var left, right []int
	for _, line := range lines {
		splitString := strings.Split(line, "   ")
		leftInt, _ := strconv.Atoi(splitString[0])
		rightInt, _ := strconv.Atoi(splitString[1])
		left = append(left, leftInt)
		right = append(right, rightInt)
	}
	return left, right
}

// Part 1
func calculateTotalDistance(leftList []int, rightList []int) (int, error) {
	if len(leftList) != len(rightList) {
		return -1, fmt.Errorf("something went wrong, leftList len != rightList len")
	}

	total := float64(0)
	for i := range leftList {
		total += math.Abs(float64(rightList[i]) - float64(leftList[i]))
	}
	return int(total), nil
}

// Part 2
func calculateTotalSimilarityScore(leftList []int, rightList []int) int {
	total := 0
	for _, leftNum := range leftList {
		rightCount := util.CountInSlice(
			rightList,
			func(x int) bool {
				return x == leftNum
			},
		)
		total += leftNum * rightCount
	}
	return total
}
