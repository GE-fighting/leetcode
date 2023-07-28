package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	flag := 1
	if (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0) {
		if dividend < 0 {
			if dividend == math.MinInt32 {
				dividend = math.MaxInt32
			} else {
				dividend = -dividend
			}
			dividend = -dividend
			divisor = -divisor
		}
	} else {
		if dividend < 0 {
			if dividend == math.MinInt32 {
				dividend = math.MaxInt32
			} else {
				dividend = -dividend
			}

		}
		if divisor < 0 {
			divisor = -divisor
		}
		flag = -1

	}
	if dividend < divisor {
		return 0
	}
	if dividend == divisor {
		return 1 * flag
	}
	low := 1
	high := dividend
	for high > low+1 {
		if divisor*((high+low)/2) > dividend {
			high = (high + low) / 2
		} else if divisor*((high+low)/2) < dividend {
			low = (high + low) / 2
		} else {
			return (high + low) / 2
		}
	}
	if divisor*high == dividend {
		return high * flag
	}
	return low * flag

}
func main() {
	fmt.Println(divide(2147483647, 1))
}
