package meatcounter

import (
	"regexp"
	"strings"
)

func CountMeatNames(text string) map[string]int {
	re := regexp.MustCompile(`\b(beef|t-bone|fatback|pastrami|pork|meatloaf|jowl|enim|bresaola)\b`)
	// Convert the input text to lowercase once
	lowercaseText := strings.ToLower(text)

	// Find all matches in the text
	matches := re.FindAllString(lowercaseText, -1)

	// Count the occurrences of each meat name
	meatCounts := make(map[string]int)
	for _, match := range matches {
		meatCounts[match]++
	}

	return meatCounts
}
