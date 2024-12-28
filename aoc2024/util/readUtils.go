package util

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ReadInputFileToStringSlice(filepath string) []string {
	var lines []string
	inputFile, _ := os.Open(filepath)
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func SplitAtFirstEmptyLine(lines []string) ([]string, []string) {
	firstEmptyLineIndex := slices.Index(lines, "")
	return lines[:firstEmptyLineIndex], lines[firstEmptyLineIndex+1:]
}

func SplitStringToIntList(line string, separator string) []int {
	var list []int
	splitString := strings.Split(line, separator)
	for _, s := range splitString {
		i, _ := strconv.Atoi(s)
		list = append(list, i)
	}
	return list
}
