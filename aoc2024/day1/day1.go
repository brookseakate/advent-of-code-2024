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

	total, _ := calculateTotalDistance(leftList, rightList)

	println(">>>>>TOTAL: %v", total)
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
