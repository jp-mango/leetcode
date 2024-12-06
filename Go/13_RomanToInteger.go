package main

/*
* https://leetcode.com/problems/roman-to-integer/description/
 */
func romanToInt(s string) int {
	romanNumerals := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var result int

	for i, v := range s {
		result += romanNumerals[string(v)]
		if i != 0 {
			if romanNumerals[string(s[i-1])] < romanNumerals[string(v)] {
				result -= 2 * romanNumerals[string(s[i-1])]
			}
		}
	}

	return result
}
