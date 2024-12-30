package day4

import (
	"aoc2024/util"
	"fmt"
	"strings"
)

func Day4() {
	//inputLines := util.ReadInputFileToStringSlice("day4/sample")
	inputLines := util.ReadInputFileToStringSlice("day4/input")

	//fmt.Printf(">>>>TOTAL: %v", countXmasesInMatrix(inputLines))
	fmt.Printf(">>>>TOTAL: %v", countXdashMasesInMatrix(inputLines))
}

// part 2
func countXdashMasesInMatrix(matrix []string) int {
	count := 0
	masString := "MAS"

	// iterate only [1:-1] x & y ranges to avoid OOB
	for li := 1; li < len(matrix)-1; li++ {
		line := matrix[li]
		for ci := 1; ci < len(line)-1; ci++ {
			char := line[ci]
			if string(char) == "A" { // only search if we found an "A"
				negativeSlopeMatch := searchNegativeSlope(matrix, masString, li, ci) || searchNegativeSlope(matrix, reverseString(masString), li, ci)
				positiveSlopeMatch := searchPositiveSlope(matrix, masString, li, ci) || searchPositiveSlope(matrix, reverseString(masString), li, ci)
				if negativeSlopeMatch && positiveSlopeMatch {
					//fmt.Printf("Found an XMAS! A at %v, %v\n", li, ci)
					count++
				}
			}
		}
	}
	return count
}

func reverseString(s string) string {
	var reversedBuilder strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		reversedBuilder.WriteByte(s[i])
	}
	return reversedBuilder.String()
}

func searchNegativeSlope(matrix []string, matchString string, x int, y int) bool {
	// OOB already guarded at call site

	// hard-code since only checking 2 chars/runes
	if matchString[0] != matrix[x-1][y-1] {
		return false
	}
	if matchString[2] != matrix[x+1][y+1] {
		return false
	}

	return true
}

func searchPositiveSlope(matrix []string, matchString string, x int, y int) bool {
	// OOB already guarded at call site

	// hard-code since only checking 2 chars/runes
	if matchString[0] != matrix[x+1][y-1] {
		return false
	}
	if matchString[2] != matrix[x-1][y+1] {
		return false
	}

	return true
}

// part 1
func countXmasesInMatrix(matrix []string) int {
	matchString := "XMAS"
	count := 0

	for li, line := range matrix {
		for ci, char := range line {
			if string(char) == string(matchString[0]) { // only search if we found an "X"
				if rightSearchContains(matrix, matchString, li, ci) {
					fmt.Printf("found right at %v, %v\n", li, ci)
					count++
				}
				if downRightSearchContains(matrix, matchString, li, ci) {
					fmt.Printf("found downRight at %v, %v\n", li, ci)
					count++
				}
				if downSearchContains(matrix, matchString, li, ci) {
					fmt.Printf("found down at %v, %v\n", li, ci)
					count++
				}
				if downLeftSearchContains(matrix, matchString, li, ci) {
					fmt.Printf("found downLeft at %v, %v\n", li, ci)
					count++
				}
				if leftSearchContains(matrix, matchString, li, ci) {
					fmt.Printf("found left at %v, %v\n", li, ci)
					count++
				}
				if upLeftSearchContains(matrix, matchString, li, ci) {
					fmt.Printf("found upLeft at %v, %v\n", li, ci)
					count++
				}
				if upSearchContains(matrix, matchString, li, ci) {
					fmt.Printf("found up at %v, %v\n", li, ci)
					count++
				}
				if upRightSearchContains(matrix, matchString, li, ci) {
					fmt.Printf("found upRight at %v, %v\n", li, ci)
					count++
				}
			}
		}
	}
	return count
}

func rightSearchContains(matrix []string, matchString string, x int, y int) bool {
	// OOB guard
	if len(matrix[x]) < y+len(matchString) {
		return false
	}

	for mi, mrune := range matchString {
		if string(matrix[x][y+mi]) != string(mrune) {
			return false
		}
	}
	return true
}

func downRightSearchContains(matrix []string, matchString string, x int, y int) bool {
	// OOB guard
	if len(matrix[x]) < y+len(matchString) || len(matrix) < x+len(matchString) {
		return false
	}

	for mi, mrune := range matchString {
		if string(matrix[x+mi][y+mi]) != string(mrune) {
			return false
		}
	}
	return true
}

func downSearchContains(matrix []string, matchString string, x int, y int) bool {
	// OOB guard
	if len(matrix) < x+len(matchString) {
		return false
	}

	for mi, mrune := range matchString {
		if string(matrix[x+mi][y]) != string(mrune) {
			return false
		}
	}
	return true
}

func downLeftSearchContains(matrix []string, matchString string, x int, y int) bool {
	// OOB guard
	if y+1-len(matchString) < 0 || len(matrix) < x+len(matchString) {
		return false
	}

	for mi, mrune := range matchString {
		if string(matrix[x+mi][y-mi]) != string(mrune) {
			return false
		}
	}
	return true
}

func leftSearchContains(matrix []string, matchString string, x int, y int) bool {
	// OOB guard
	if y+1-len(matchString) < 0 {
		return false
	}

	for mi, mrune := range matchString {
		if string(matrix[x][y-mi]) != string(mrune) {
			return false
		}
	}
	return true
}
func upLeftSearchContains(matrix []string, matchString string, x int, y int) bool {
	// OOB guard
	if y+1-len(matchString) < 0 || x+1-len(matchString) < 0 {
		return false
	}

	for mi, mrune := range matchString {
		if string(matrix[x-mi][y-mi]) != string(mrune) {
			return false
		}
	}
	return true
}

func upSearchContains(matrix []string, matchString string, x int, y int) bool {
	// OOB guard
	if x+1-len(matchString) < 0 {
		return false
	}

	for mi, mrune := range matchString {
		if string(matrix[x-mi][y]) != string(mrune) {
			return false
		}
	}
	return true
}

func upRightSearchContains(matrix []string, matchString string, x int, y int) bool {
	// OOB guard
	if x+1-len(matchString) < 0 || len(matrix[x]) < y+len(matchString) {
		return false
	}

	for mi, mrune := range matchString {
		if string(matrix[x-mi][y+mi]) != string(mrune) {
			return false
		}
	}
	return true
}
