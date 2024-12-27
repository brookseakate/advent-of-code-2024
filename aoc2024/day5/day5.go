package day5

import (
	"aoc2024/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Day5() error {
	//ruleLines, updateLines := util.SplitAtFirstEmptyLine(util.ReadInputFileToStringSlice("day5/sample"))
	ruleLines, updateLines := util.SplitAtFirstEmptyLine(util.ReadInputFileToStringSlice("day5/input"))

	rules := parseLinesToRules(ruleLines)
	println("Rules:")
	for k, v := range rules {
		fmt.Printf("%v : %v\n", k, v)
	}

	updates := parseLinesToUpdates(updateLines)
	println("Updates:")
	for _, update := range updates {
		fmt.Printf("%v\n", update)
	}

	total := calculateValidUpdateMiddlePageTotal(rules, updates)
	fmt.Printf(">>>>>>>>>TOTAL: %v", total)
	return nil
}

type Update []int

func parseLinesToRules(lines []string) map[int][]int {
	rules := map[int][]int{}

	for _, line := range lines {
		pages := strings.Split(line, "|")
		first, _ := strconv.Atoi(pages[0])
		second, _ := strconv.Atoi(pages[1])

		if _, ok := rules[first]; !ok {
			rules[first] = []int{second}
		} else {
			rules[first] = append(rules[first], second)
		}
	}

	return rules
}

func parseLinesToUpdates(lines []string) []Update {
	updates := make([]Update, len(lines))
	for lineIndex, line := range lines {
		stringsUpdate := strings.Split(line, ",")

		// transform []string to []int
		integers := make([]int, len(stringsUpdate))
		for i, stringPage := range stringsUpdate {
			n, err := strconv.Atoi(stringPage)
			if err != nil {
				println("OOPS err: %v", err)
			}
			integers[i] = n
		}

		updates[lineIndex] = integers
	}
	return updates
}

func calculateValidUpdateMiddlePageTotal(rules map[int][]int, updates []Update) int {
	total := 0
	for _, update := range updates {
		if update.isValidUpdate(rules) {
			total += update.getMiddleNumber()
		}
	}
	return total
}

func (update Update) isValidUpdate(rules map[int][]int) bool {
	for thisPageIndex, thisPage := range update {
		pagesRequiredLater := rules[thisPage]
		for _, requiredLaterPage := range pagesRequiredLater {
			var requiredLaterPageIndex = slices.Index(update, requiredLaterPage) // -1 if not present
			if requiredLaterPageIndex < thisPageIndex && requiredLaterPageIndex >= 0 /* ignore not present */ {
				return false
			}
		}
	}
	return true
}

func (update Update) getMiddleNumber() int {
	return update[len(update)/2]
}
