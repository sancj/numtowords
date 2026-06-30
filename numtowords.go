package numtowords

import (
	"fmt"
	"strings"
)

// MaxNumber is the maximum number that can be converted to words.
const MaxNumber = 9999

// Convert converts a number to its word representation.
func Convert(number int) string {
	if number < 0 || number > MaxNumber {
		return "Number out of range"
	}
	if number == 0 {
		return "zero"
	}

	units := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	unitsTeens := []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	unitsTens := []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

	var parts []string
	if result := procNumber(&number, 1000, 1000, 0, units); result != "" {
		parts = append(parts, fmt.Sprintf("%s thousand", result))
	}
	if result := procNumber(&number, 100, 100, 0, units); result != "" {
		parts = append(parts, fmt.Sprintf("%s hundred", result))
	}
	if result := procNumber(&number, 20, 10, 0, unitsTens); result != "" {
		parts = append(parts, result)
	}
	if result := procNumber(&number, 10, 1, -10, unitsTeens); result != "" {
		parts = append(parts, result)
	}
	if result := procNumber(&number, 1, 1, 0, units); result != "" {
		parts = append(parts, result)
	}
	return strings.Join(parts, " ")
}

// procNumber extracts a word from words for the current place value.
// threshold: minimum value of *number to process this place.
// div: divisor used to compute the index into words.
// offset: added to the raw index (use -10 for teens).
func procNumber(number *int, threshold, div, offset int, words []string) string {
	if *number < threshold {
		return ""
	}
	idx := *number/div + offset
	*number = *number - (idx-offset)*div
	return words[idx]
}
