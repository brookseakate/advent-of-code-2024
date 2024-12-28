package day3

import (
	"aoc2024/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day3() {
	//fmt.Printf(">>>>>TOTAL: %v", parseInputToMulListAndCalculateSum("day3/sample"))
	//fmt.Printf(">>>>>TOTAL: %v", parseInputToMulListAndCalculateSum("day3/sampleIgnorableCases"))
	fmt.Printf(">>>>>TOTAL: %v", parseInputToMulListAndCalculateSum("day3/input"))
}

func parseInputToMulListAndCalculateSum(filepath string) int {
	inputLines := util.ReadInputFileToStringSlice(filepath)

	inputString := strings.Join(inputLines, "")
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	mulStringSubmatches := re.FindAllStringSubmatch(inputString, -1)
	//fmt.Println("%v\n", mulStringSubmatches)

	fmt.Printf("mul count: %v\n", len(mulStringSubmatches))
	sum := 0
	for _, mul := range mulStringSubmatches {
		fmt.Printf("%v\n", mul)
		left, _ := strconv.Atoi(mul[1])
		right, _ := strconv.Atoi(mul[2])
		sum += left * right
	}
	return sum
}
