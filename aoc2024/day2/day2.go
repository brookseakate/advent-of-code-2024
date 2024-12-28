package day2

import (
	"aoc2024/util"
	"slices"
)

func Day2() {
	//safeReportTotal := countSafeReports(getInput("day2/sample"))
	safeReportTotal := countSafeReports(getInput("day2/input"))

	println(">>>>> TOTAL: %v", safeReportTotal)
}

func getInput(filepath string) [][]int {
	lines := util.ReadInputFileToStringSlice(filepath)
	var output [][]int
	for _, line := range lines {
		output = append(output, util.SplitStringToIntList(line, " "))
	}
	return output
}

func countSafeReports(reports [][]int) int {
	count := 0
	for _, report := range reports {
		if isSafe(report) {
			count++
		}
	}
	return count
}

func isSafe(report []int) bool {
	if report[1] < report[0] {
		// decreasing. reverse, to pass into below block
		slices.Reverse(report)
	}

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func getSampleInputReports() [][]int {
	return [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
}
