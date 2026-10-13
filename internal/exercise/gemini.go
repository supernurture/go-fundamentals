package exercise

func EvenAndSortAsc(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	result := []int{}
	for _, v := range nums {
		if v%2 == 0 {
			result = append(result, v)
		}
	}

	for x := 0; x < len(result); x++ {
		if x == len(result)-1 {
			break
		}
		for y := x + 1; y < len(result); y++ {
			if result[y] < result[x] {
				result[x], result[y] = result[y], result[x]
			}
		}
	}

	return result
}

func OddAndSortDesc(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	result := []int{}
	for _, v := range nums {
		if v%2 != 0 {
			result = append(result, v)
		}
	}

	for x := 0; x < len(result); x++ {
		if x == len(result)-1 {
			break
		}
		for y := x + 1; y < len(result); y++ {
			if result[y] > result[x] {
				result[y], result[x] = result[x], result[y]
			}
		}
	}

	return result
}

func SumOfTwoNums(a, b int) int {
	return a + b
}
