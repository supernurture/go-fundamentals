package exercise

func EvenNumberAndSort(nums []int) []int {
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
