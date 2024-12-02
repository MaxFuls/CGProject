package parsing

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ParseFormula(formula string) map[string]int {
	elementCounts := make(map[string]int)

	elementPattern := regexp.MustCompile(`([A-Z][a-z]*)(\d*)`)
	groupPattern := regexp.MustCompile(`\(([^()]+)\)(\d*)`)

	for {
		matches := groupPattern.FindStringSubmatch(formula)
		if matches == nil {
			break
		}

		group := matches[1]
		multiplier := 1
		if matches[2] != "" {
			multiplier, _ = strconv.Atoi(matches[2])
		}

		expanded := ""
		for _, match := range elementPattern.FindAllStringSubmatch(group, -1) {
			element := match[1]
			count := 1
			if match[2] != "" {
				count, _ = strconv.Atoi(match[2])
			}
			expanded += fmt.Sprintf("%s%d", element, count*multiplier)
		}

		formula = strings.Replace(formula, matches[0], expanded, 1)
	}

	for _, match := range elementPattern.FindAllStringSubmatch(formula, -1) {
		element := match[1]
		count := 1
		if match[2] != "" {
			count, _ = strconv.Atoi(match[2])
		}
		elementCounts[element] += count
	}

	return elementCounts
}
