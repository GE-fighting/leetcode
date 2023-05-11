package main

func canArrange(arr []int, k int) bool {
	record := map[int]int{}
	for _, v := range arr {
		remainder := v % k
		if remainder < 0 {
			remainder = remainder + k
		}
		if val, ok := record[remainder]; ok {
			record[remainder] = val + 1
		} else {
			record[remainder] = 1
		}
	}
	for key, v := range record {
		if key == 0 {
			if v%2 != 0 {
				return false
			}
		} else if key*2 == k {
			if v%2 != 0 {
				return false
			}
		} else {
			if record[k-key] != v {
				return false
			}
		}
	}
	return true
}
