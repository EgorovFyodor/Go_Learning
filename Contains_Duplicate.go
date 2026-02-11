package main

func hasDuplicate(nums []int) bool {
	for _, n1 := range nums {
		count := 0
		for _, n2 := range nums {
			if n1 == n2 {
				count++
			}
		}
		if count > 1 {
			return true
		}
	}
	return false
}
