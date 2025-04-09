package main

import "fmt"

func isValid(s string) bool {
	st := []rune{}

	for _, r := range s {
		if r == '(' || r == '{' || r == '[' {
			st = append(st, r)
			continue
		}

		if r == ')' {
			temp := st[len(st)-1]
			st = st[:len(st)-1]

			if temp == '(' {
				continue
			} else {
				return false
			}
		}

		if r == '}' {
			temp := st[len(st)-1]
			st = st[:len(st)-1]

			if temp == '{' {
				continue
			} else {
				return false
			}
		}

		if r == ']' {
			temp := st[len(st)-1]
			st = st[:len(st)-1]

			if temp == '[' {
				continue
			} else {
				return false
			}
		}
	}

	if len(st) > 0 {
		return false
	}
	return true
}

func main() {
	result := isValid("(]")
	fmt.Println(result)

}
