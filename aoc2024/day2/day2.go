package day2

import (
	"aoc2024/util"
	"fmt"
	"slices"
)

func Day2() {
	//safeReportTotal := countSafeReports(getInput("day2/sample"))
	safeReportTotal := countSafeReports(getInput("day2/input"))

	//safeWithToleranceReportTotal := countSafeReportsWithTolerance(getInput("day2/sample"), 1)
	//safeWithToleranceReportTotal := countSafeReportsWithTolerance(getInput("day2/input"), 1)
	safeWithToleranceReportTotal := countSafeReportsWithTolerance(getInput("day2/edgeCasesSample"), 1) // input from: https://www.reddit.com/r/adventofcode/comments/1h4shdu/2024_day_2_part2_edge_case_finder/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button

	println(">>>>> TOTAL SAFE WITHOUT TOLERANCE: %v", safeReportTotal)
	println(">>>>> TOTAL SAFE WITH TOLERANCE: %v", safeWithToleranceReportTotal)
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
		if safe, _ := isSafe(report); safe {
			count++
		}
	}
	return count
}

func countSafeReportsWithTolerance(reports [][]int, toleranceCount int) int {
	count := 0
	originalTolerance := toleranceCount

	for reportIndex, report := range reports {
		if report[1] < report[0] {
			// decreasing. reverse, to pass into below block
			slices.Reverse(report)
		}

		if safe, indexFoundUnsafe := isSafe(report); safe {
			fmt.Printf("found safe reportIndex: %v, report values: %v\n", reportIndex, report)
			count++
		} else {
			fmt.Printf("found unsafe reportIndex: %v, report values: %v\n", reportIndex, report)
			if toleranceCount > 0 {
				toleranceCount--
				reportWithHigherIndexRemoved := slices.Concat(report[:indexFoundUnsafe], report[indexFoundUnsafe+1:])
				reportWithLowerIndexRemoved := slices.Concat(report[:indexFoundUnsafe-1], report[indexFoundUnsafe:])
				countMaybeSaferHigherIndexRemoved := countSafeReportsWithTolerance([][]int{reportWithHigherIndexRemoved}, toleranceCount)
				countMaybeSaferLowerIndexRemoved := countSafeReportsWithTolerance([][]int{reportWithLowerIndexRemoved}, toleranceCount)
				if countMaybeSaferHigherIndexRemoved > 0 || countMaybeSaferLowerIndexRemoved > 0 {
					count += max(countMaybeSaferHigherIndexRemoved, countMaybeSaferLowerIndexRemoved)
				}
			}
		}
		toleranceCount = originalTolerance
	}
	return count
}

func isSafe(report []int) (bool, int) {
	if report[1] < report[0] {
		// decreasing. reverse, to pass into below block
		slices.Reverse(report)
	}

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if diff < 1 || diff > 3 {
			return false, i
		}
	}
	return true, -1
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
