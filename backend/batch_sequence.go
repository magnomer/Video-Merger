package backend

import (
	"fmt"
	"sort"
)

func LSequenceFind(files []LClip) []string {
	if len(files) == 0 {
		return nil
	}

	var warnings []string

	// Parts are sorted ascending. A leading 0 is accepted as the first part:
	// it covers a real "(0)" file and the unnumbered fallback, both of which
	// carry LClipNumber 0 and are treated as the first part.
	first := files[0].LClipNumber
	if first != 0 && first != 1 {
		warnings = append(warnings, fmt.Sprintf("First part number is %d, not 1", first))
	}

	for i := 1; i < len(files); i++ {
		prev := files[i-1].LClipNumber
		curr := files[i].LClipNumber

		if curr == prev {
			warnings = append(warnings, fmt.Sprintf("Duplicate part number %d", curr))
			continue
		}

		if curr > prev+1 {
			warnings = append(
				warnings,
				fmt.Sprintf("Missing part numbers between %d and %d", prev, curr),
			)
		}
	}

	return warnings
}

func LSequenceDuplicateFind(files []LClip) []string {
	if len(files) < 2 {
		return nil
	}

	counts := map[int]int{}
	for _, file := range files {
		counts[file.LClipNumber]++
	}

	numbers := []int{}
	for number, count := range counts {
		if count > 1 {
			numbers = append(numbers, number)
		}
	}
	sort.Ints(numbers)

	warnings := []string{}
	for _, number := range numbers {
		warnings = append(warnings, fmt.Sprintf("Duplicate part number %d creates an ambiguous merge order.", number))
	}

	return warnings
}
