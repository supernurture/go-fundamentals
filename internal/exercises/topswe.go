package exercises

import (
	"fmt"
	"math/bits"
	"regexp"
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
	for x := range nums {
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
	if len([]rune(a)) != len([]rune(b)) {
		return false
	}

	sortedA := []rune(a)
	sortedB := []rune(b)
	slices.Sort(sortedA)
	slices.Sort(sortedB)

	return string(sortedA) == string(sortedB)
}

func ValidPalindrome(s string) bool {
	if strings.TrimSpace(s) == "" {
		return true
	}
	sanitator := regexp.MustCompile("[^a-zA-Z0-9]")
	// Convert the characters to lowercase and then sanitize them.
	sanitizedS := sanitator.ReplaceAllString(strings.ToLower(s), "")

	for x, y := 0, len(sanitizedS)-1; x < y; x, y = x+1, y-1 {
		if sanitizedS[x] != sanitizedS[y] {
			return false
		}
	}
	return true
}

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	for k, v := range strs[0] {
		for _, v2 := range strs[1:] {
			if len(v2)-1 < k || v != rune(v2[k]) {
				return strs[0][:k]
			}
		}
		if len(strs[0])-1 == k {
			return strs[0]
		}
	}
	return ""
}

func PowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}

func NumberOfSetBits(n int) int {
	return bits.OnesCount(uint(n))
}

func RangeSumQuery(nums []int, queries [][]int) []int {
	prefixs := make([]int, len(nums)+1)
	for k, v := range nums {
		prefixs[k+1] = prefixs[k] + v
	}

	result := []int{}
	for _, v := range queries {
		add := prefixs[v[1]+1] - prefixs[v[0]]
		result = append(result, add)
	}
	return result
}

func FibonacciNumber(n int) int {
	if n < 0 {
		return -1
	}
	fibs := []int{0, 1}
	if n == 0 || n == 1 {
		return fibs[n]
	}
	for x := 2; x <= n; x++ {
		fibs = append(fibs, fibs[x-1]+fibs[x-2])
	}
	return fibs[n]
}

func ClimbingStairs(n int) int {
	if n < 2 {
		return 1
	}

	prev1 := 1
	prev2 := 1
	for x := 2; x <= n; x++ {
		prev2 += prev1
		prev1 = prev2 - prev1
	}

	return prev2
}
