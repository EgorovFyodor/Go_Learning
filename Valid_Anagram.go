package main

func isAnagram(s string, t string) bool {
	arr1 := createMap(s)
	arr2 := createMap(t)

	if len(arr1) != len(arr2) {
		return false
	}
	for k, v1 := range arr1 {
		if v2, ok := arr2[k]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}

func createMap(s string) map[string]int {
	arr := make(map[string]int)
	for _, ch := range []rune(s) {
		count := 0
		for _, t := range []rune(s) {
			if ch == t {
				count++
			}
		}
		if _, ok := arr[string(ch)]; !ok {
			arr[string(ch)] = count
		}
	}
	return arr
}
