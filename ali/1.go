package main

//题目描述
//定义f(0) = 1，f(2n) = f(n), f(2n+1) = f(n)*2;现在小红有一个数组a，她想知道f(a_i) 共有多少个不同的数;
//输入描述
// 第一行输入一个整数n,表示数组的长度;第二行n个整数a_i，表示数组的元素
//输入样例
// 5
//1 2 3 4 5
// 输出
// 2

func f(n int) int {
	res := 1
	for n > 0 {
		if n%2 != 0 {
			res *= 2
		}
		n /= 2
	}
	return res
}

//func main() {
//	var n int
//	var a int
//	fmt.Scan(&n)
//	res := map[int]int{}
//	for n > 0 {
//		fmt.Scan(&a)
//		res[f(a)] = 1
//		n--
//	}
//	fmt.Println(len(res))
//}
