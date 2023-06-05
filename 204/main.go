package main

import "fmt"

func countPrimes(n int) int {
	record := make([]bool, n)
	for i := range record {
		record[i] = true
	}
	res := 0
	for i := 2; i < n; i++ {
		if record[i] {
			for j := 2; i*j < n; j++ {
				record[i*j] = false
			}
			res++
		}
	}
	return res
}

func main() {
	primes := countPrimes(5000000)
	fmt.Println(primes)
}
