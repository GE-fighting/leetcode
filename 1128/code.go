package main

func numEquivDominoPairs(dominoes [][]int) int {
	record := [4000]int{}
	res := []int{}

	length := len(dominoes)
	for i := 0; i < length; i++ {
		if record[i] == 0 {
			count := 0
			for j := i + 1; j < length; j++ {
				if isEquivcal(dominoes[i], dominoes[j]) {
					record[i] = 1
					record[j] = 1
					count++
				}
			}
			res = append(res, count+1)
		}

	}
	sum := 0
	for i := 0; i < len(res); i++ {
		sum += res[i] * (res[i] - 1) / 2
	}
	return sum
}

func isEquivcal(a, b []int) bool {
	if a[0] == b[0] && a[1] == b[1] {
		return true
	}
	if a[0] == b[1] && a[1] == b[0] {
		return true
	}
	return false
}
