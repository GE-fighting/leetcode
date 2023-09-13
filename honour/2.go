package main

func calculateRes(nums []int, target int) (int, int, int) {
	left := 0
	right := len(nums) - 1
	for right > left {
		if nums[left]+nums[right] > target {
			right--
		} else if nums[left]+nums[right] < target {
			left++
		} else {
			return 1, nums[left], nums[right]
		}
	}
	return 0, 0, 0
}

//func main() {
//	reader := bufio.NewReader(os.Stdin)
//	text, _ := reader.ReadString('\n')
//	numStr := strings.Split(text, " ")
//	nums := []int{}
//	for _, val := range numStr {
//		var n int
//		fmt.Sscanf(val, "%d", &n)
//		nums = append(nums, n)
//	}
//	fmt.Println(nums)
//	var target int
//	fmt.Scan(&target)
//	fmt.Println(calculateRes(nums, target))
//}
