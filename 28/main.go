package main

import "fmt"

func strStr(haystack string, needle string) int {

	parentLen := len(haystack)
	patternLen := len(needle)

	parentHash := 0
	patternHash := 0
	power := 1

	for i := 0; i < parentLen-patternLen+1; i++ {
		if i == 0 {
			for j := 0; j < patternLen; j++ {
				parentHash += int(haystack[patternLen-1-j]) * power
				patternHash += int(needle[patternLen-1-j]) * power

				if j < patternLen-1 {
					power *= 403
				}
			}
		} else {
			parentHash = (403 * (parentHash - int(haystack[i-1])*power)) + int(haystack[patternLen+i-1])
		}

		if parentHash == patternHash {
			found := true
			for j := 0; j < patternLen; j++ {
				if haystack[i+j] != needle[j] {
					found = false
					break
				}
			}
			if found {
				return i
			}

		}
	}

	return -1
}

func main() {
	result := strStr("sadbutsad", "sad")
	fmt.Println(result)
	result = strStr("leetcode", "leeto")
	fmt.Println(result)

}
