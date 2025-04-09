package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {

	sort.Ints(nums)

	result := [][]int{}
	for a := 0; a < len(nums)-3; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		for b := a + 1; b < len(nums)-2; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}
			c, d := b+1, len(nums)-1
			for c < d {
				temp := nums[a] + nums[b] + nums[c] + nums[d]
				if temp < target {
					c++
				} else if temp > target {
					d--
				} else {
					result = append(result, []int{nums[a], nums[b], nums[c], nums[d]})

					for c < d && nums[c] == nums[c+1] {
						c++
					}
					for c < d && nums[d] == nums[d-1] {
						d--
					}
					c++
					d--
				}
			}
		}
	}
	return result
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	for a := 0; a < len(nums)-3; a++ {
		// 중복된 값 건너뛰기
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}

		for b := a + 1; b < len(nums)-2; b++ {
			// 중복된 값 건너뛰기
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}

			c, d := b+1, len(nums)-1

			for c < d {
				sum := nums[a] + nums[b] + nums[c] + nums[d]

				if sum < target {
					c++
				} else if sum > target {
					d--
				} else {
					// 정답 추가
					result = append(result, []int{nums[a], nums[b], nums[c], nums[d]})

					// 중복 건너뛰기
					for c < d && nums[c] == nums[c+1] {
						c++
					}
					for c < d && nums[d] == nums[d-1] {
						d--
					}

					c++
					d--
				}
			}
		}
	}

	return result
}

func main() {
	result := fourSum([]int{1, 0, -1, 0, 2}, 0)
	fmt.Println(result)

}
