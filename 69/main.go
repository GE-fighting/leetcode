package main

func mySqrt(x int) int {
	low := 1
	high := x/2 + 1
	for high > low+1 {
		if ((high+low)/2)*((high+low)/2) > x {
			high = (high + low) / 2
		} else if ((high+low)/2)*((high+low)/2) < x {
			low = (low + high) / 2
		} else {
			return (low + high) / 2
		}
	}
	return low
}

func main() {
	println(mySqrt(4))
}
