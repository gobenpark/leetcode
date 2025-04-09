package main

import "fmt"

var runemap = map[int][]rune{
	50: {'a', 'b', 'c'},
	51: {'d', 'e', 'f'},
	52: {'g', 'h', 'i'},
	53: {'j', 'k', 'l'},
	54: {'m', 'n', 'o'},
	55: {'p', 'q', 'r', 's'},
	56: {'t', 'u', 'v'},
	57: {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	return Combinations([]rune{}, digits)
}

func Combinations(data []rune, digits string) []string {
	if len(digits) == 0 {
		return []string{string(data)}
	}

	result := []string{}
	first := int(digits[0])
	for _, v := range runemap[first] {
		fmt.Println(string(v))
		result = append(result, Combinations(append(data, v), digits[1:])...)
	}

	return result
}

func main() {
	result := letterCombinations("7")
	fmt.Println(result)
}
