package exercises

import (
	"fmt"
	"slices"
	"strings"
)

func FizzBuzz(n int) []string {
	if n < 1 {
		return []string{}
	}

	var result []string
	for x := 1; x <= n; x++ {
		if x%5 == 0 && x%3 == 0 {
			result = append(result, "FizzBuzz")
		} else if x%5 == 0 {
			result = append(result, "Buzz")
		} else if x%3 == 0 {
			result = append(result, "Fizz")
		} else {
			result = append(result, fmt.Sprintf("%d", x))
		}
	}
	return result
}

func TwoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}

	seen := make(map[int]int, len(nums))
	for x := 0; x < len(nums); x++ {
		if v, ok := seen[target-nums[x]]; ok {
			return []int{v, x}
		}
		seen[nums[x]] = x
	}
	return []int{}
}

func ReverseString(v []string) {
	if len(v) == 0 {
		return
	}
	for x, y := 0, len(v)-1; x < y; x, y = x+1, y-1 {
		v[x], v[y] = v[y], v[x]
	}
}

func ValidAnagram(a, b string) bool {
	if len(strings.Split(a, "")) != len(strings.Split(b, "")) {
		return false
	}

	sortedA := []rune(a)
	sortedB := []rune(b)
	slices.Sort(sortedA)
	slices.Sort(sortedB)

	return string(sortedA) == string(sortedB)
}
