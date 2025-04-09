package main

import (
	"fmt"
	"sort"
)

func long(strs []string) string {

	fl := ""
	for i := 0; i < len(strs); i++ {
		if fl == "" && i == 0 {
			fl = strs[i]
			continue
		}

		fl = CheckPrefix(fl, strs[i])
		if fl == "" {
			return ""
		}
	}
	return fl
}

func CheckPrefix(prefix string, target string) string {
	maxIndex := 0
	if len(target) < len(prefix) {
		maxIndex = len(target)
	} else {
		maxIndex = len(prefix)
	}

	var i int = 0
	for i < maxIndex {
		if prefix[i] != target[i] {
			break
		}
		i++
	}
	return prefix[:i]
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	sort.Strings(strs)
	first, last := strs[0], strs[len(strs)-1]
	result := ""
	for i := 0; i < len(first); i++ {
		if i < len(last) && first[i] == last[i] {
			result += string(first[i])
		} else {
			break
		}
	}
	return result
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	//fmt.Println(long([]string{"dog", "racecar", "car"}))
	//fmt.Println(long([]string{"", "b"}))

}
