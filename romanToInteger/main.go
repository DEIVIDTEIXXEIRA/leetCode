package main

import "fmt"

func main() {
	roman := "XIV"
	result := RomanToInt(roman)

	fmt.Printf("result: %v\n", result)
}

func RomanToInt(s string) int {
	romanNumerals := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	total := 0
	for i := 0; i < len(s); i++ {
		currentChar := string(s[i])
		value := romanNumerals[currentChar]

		// Verifica se o próximo numeral é maior que o atual
		if i+1 < len(s) && romanNumerals[string(s[i+1])] > value {
			total -= value
		} else {
			total += value
		}
	}

	return total
}

func RomanToInt2(s string) int {
	sum := 0

	rim := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i, v := range s {
		sum += rim[string(v)]
		if i != 0 {
			if rim[string(s[i-1])] < rim[string(v)] {
				sum -= 2 * rim[string(s[i-1])]
			}
		}
	}

	return sum
}
