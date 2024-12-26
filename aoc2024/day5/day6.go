package day5

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Day5() {
	rules := parseLinesToRules(getInputRuleLines())

	for k, v := range rules {
		fmt.Printf("%v : %v\n", k, v)
	}

	updates := getInputUpdates()
	for _, update := range updates {
		fmt.Printf("%v\n", update)
		//fmt.Print(update)
	}

	total := calculateValidUpdateMiddlePageTotal(rules, updates)
	fmt.Printf(">>>>>>>>>TOTAL: %v", total)
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

type Update []int

func (update Update) isValidUpdate(rules map[int][]int) bool {
	for thisPageIndex, thisPage := range update {
		pagesRequiredLater := rules[thisPage]
		for _, requiredLaterPage := range pagesRequiredLater {
			var requiredLaterPageIndex = slices.Index(update, requiredLaterPage) // -1 if not present
			if requiredLaterPageIndex < thisPageIndex && requiredLaterPageIndex >= 0 {
				return false
			}
		}
	}
	return true
}

func (update Update) getMiddleNumber() int {
	return update[len(update)/2]
}

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

func getInputRuleLines() []string {
	// TODO: parse input from file
	return []string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
	}
}

func getInputUpdates() []Update {
	// TODO: parse input from file
	return []Update{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}
}
