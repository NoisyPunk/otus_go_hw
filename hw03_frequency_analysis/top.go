package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	frequency, maxCounter := frequencyCalculation(text)
	if frequency == nil {
		return nil
	}

	result := lexicographicSorting(frequency, maxCounter)

	return result
}

func frequencyCalculation(text string) (frequency map[string]int, maxCounter int) {
	frequency = make(map[string]int)

	splitText := strings.Fields(text)

	if len(splitText) == 0 {
		return nil, 0
	}

	for _, comparativeValue := range splitText {
		counter := 0
		for _, j := range splitText {
			if comparativeValue == j {
				counter++
			}
		}
		frequency[comparativeValue] = counter
		if counter > maxCounter {
			maxCounter = counter
		}
	}
	return
}

func lexicographicSorting(data map[string]int, maxCounter int) (result []string) {
	var sortedFrequency []string

	for i := maxCounter; i > 0; i-- {
		for key, value := range data {
			if value == i {
				sortedFrequency = append(sortedFrequency, key)
			}
		}
		sort.Strings(sortedFrequency)
		result = append(result, sortedFrequency...)
		sortedFrequency = nil
	}
	return result[:10]
}
