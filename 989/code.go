package main

import "strconv"

func addToArrayForm(num []int, k int) []int {
	kArray := []int{}
	kStr := strconv.Itoa(k)
	for _, v := range kStr {
		num := int(v - '0')
		kArray = append(kArray, num)
	}
	num1 := sliceOverTrun(num)
	num2 := sliceOverTrun(kArray)
	len1 := len(num1)
	len2 := len(num2)
	res := []int{}
	count := 1
	flag := 0
	for {
		if count > len1 && count > len2 && flag == 0 {
			break
		}
		a := 0
		b := 0
		if count <= len1 {
			a = num1[count-1]
		}
		if count <= len2 {
			b = num2[count-1]
		}
		x := a + b + flag
		if x >= 10 {
			x = x - 10
			flag = 1
		} else {
			flag = 0
		}
		res = append(res, x)
		count++
	}
	return sliceOverTrun(res)

}

func sliceOverTrun(a []int) []int {
	res := []int{}
	for i := len(a) - 1; i >= 0; i-- {
		res = append(res, a[i])
	}
	return res
}
