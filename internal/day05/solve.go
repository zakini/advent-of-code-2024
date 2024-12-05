package day05

import (
	"fmt"
	"slices"
	"strings"
	"zakini/advent-of-code-2024/internal/utils"
)

type pageOrderingRule struct {
	before int
	after  int
}

func SolvePart1(input string) int {
	rules, pageLists := parseInput(input)

	result := 0
	for _, pageList := range pageLists {
		if pageListValid(rules, pageList) {
			result += pageList[(len(pageList)-1)/2]
		}
	}

	return result
}

func parseInput(input string) ([]pageOrderingRule, [][]int) {
	lines := strings.Split(input, "\n")

	rules := make([]pageOrderingRule, 0, len(lines))
	pageLists := make([][]int, 0, len(lines))
	foundBlankLine := false
	for _, line := range lines {
		if !foundBlankLine {
			if line == "" {
				foundBlankLine = true
				continue
			}

			nums := strings.Split(line, "|")
			utils.Assert(len(nums) == 2, fmt.Sprintf("Invalid page ordering rule line: %v", line))
			before := utils.ParseIntAndAssert(nums[0])
			after := utils.ParseIntAndAssert(nums[1])
			// TODO do we need to be able to handle rules like these?
			utils.Assert(before != after, fmt.Sprintf("Found page ordering rule with duplicate page numbers: %v", line))
			rules = append(rules, pageOrderingRule{before, after})
		} else {
			nums := strings.Split(line, ",")
			pageList := make([]int, len(nums))
			for i, num := range nums {
				pageList[i] = utils.ParseIntAndAssert(num)
			}
			// TODO do we need to be able to handle page lists like these?
			utils.Assert(len(pageList) == len(unique(pageList)), fmt.Sprintf("Found page list with duplicate page numbers: %v", pageList))
			pageLists = append(pageLists, pageList)
		}
	}

	return rules, pageLists
}

func pageListValid(rules []pageOrderingRule, pageList []int) bool {
	for _, rule := range rules {
		if !slices.Contains(pageList, rule.before) || !slices.Contains(pageList, rule.after) {
			continue
		}

		if slices.Index(pageList, rule.after) < slices.Index(pageList, rule.before) {
			return false
		}
	}

	return true
}

func unique(s []int) []int {
	found := make(map[int]struct{})
	u := make([]int, 0, len(s))
	for _, val := range s {
		if _, valAlreadyFound := found[val]; !valAlreadyFound {
			u = append(u, val)
			found[val] = struct{}{}
		}
	}

	return u
}
