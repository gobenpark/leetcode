package main

import (
	"fmt"
	"strings"
)

func generateParenthesis(n int) []string {
	var result []string
	var sb strings.Builder
	sb.Grow(n * 2)

	var backtrack func(open int, close int, sb *strings.Builder, maxint int) bool
	backtrack = func(open int, close int, sb *strings.Builder, maxint int) bool {
		if open == close && open == maxint {
			result = append(result, sb.String())
			return true
		}

		if open < maxint {
			copys := sb.String()
			sb.WriteRune('(')
			backtrack(open+1, close, sb, maxint)

			sb.Reset()
			sb.WriteString(copys)
		}

		if open > close {
			sb.WriteRune(')')
			backtrack(open, close+1, sb, maxint)
		}
		return false
	}

	sb.WriteRune('(')
	backtrack(1, 0, &sb, n)
	return result
}

func main() {
	result := generateParenthesis(1)
	fmt.Println(result)

}
