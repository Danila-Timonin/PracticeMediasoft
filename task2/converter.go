package main

import (
	"fmt"
	"strings"
)

type RomanConverter struct {
	romanMap map[byte]int
}

func NewRomanConverter() *RomanConverter {
	return &RomanConverter{
		romanMap: map[byte]int{
			'I': 1,
			'V': 5,
			'X': 10,
			'L': 50,
			'C': 100,
			'D': 500,
			'M': 1000,
		},
	}
}

func (rc *RomanConverter) RomanToArabic(roman string) (int, error) {
	if roman == "" {
		return 0, fmt.Errorf("пустая строка")
	}

	roman = strings.ToUpper(roman)
	total := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		currentValue, exists := rc.romanMap[roman[i]]
		if !exists {
			return 0, fmt.Errorf("недопустимый символ: %c", roman[i])
		}

		if currentValue < prevValue {
			total -= currentValue
		} else {
			total += currentValue
		}
		prevValue = currentValue
	}

	return total, nil
}
