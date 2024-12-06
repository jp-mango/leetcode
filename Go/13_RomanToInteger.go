package main

func romanToInt(s string) int {
	RomanNumerals := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var result int

	for _, v := range s {
		//TODO: add logic for subtracting nums
		result += RomanNumerals[string(v)]
	}

	return result
}
