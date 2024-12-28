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
	safeWithToleranceReportTotal := countSafeReportsWithTolerance(getInput("day2/input"), 1)
	//safeWithToleranceReportTotal := countSafeReportsWithTolerance(getInput("day2/edgeCasesSample"), 1) // input from: https://www.reddit.com/r/adventofcode/comments/1h4shdu/2024_day_2_part2_edge_case_finder/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button

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
		if safe, _ := isSafe(report); safe {
			//fmt.Printf("found safe reportIndex: %v, report values: %v\n", reportIndex, report)
			count++
		} else {
			//fmt.Printf("found unsafe reportIndex: %v, report values: %v\n", reportIndex, report)
			if toleranceCount > 0 {
				toleranceCount--

				for i := range report {
					retriedReport := slices.Delete(slices.Clone(report), i, i+1)
					if retriedSafe, _ := isSafe(retriedReport); retriedSafe {
						fmt.Printf("found safe on retry reportIndex: %v, report values: %v\n", reportIndex, retriedReport)
						count++
						break
					}
				}

				toleranceCount = originalTolerance
			}
		}
	}
	return count
}

func isSafe(report []int) (bool, int) {
	reportToEvaluate := slices.Clone(report)
	if report[1] < report[0] {
		// decreasing. reverse, to pass into below block
		slices.Reverse(reportToEvaluate)
	}

	for i := 1; i < len(reportToEvaluate); i++ {
		diff := reportToEvaluate[i] - reportToEvaluate[i-1]
		if diff < 1 || diff > 3 {
			return false, i
		}
	}
	return true, -1
}
