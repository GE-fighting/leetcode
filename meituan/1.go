package main

func isOk(nums []int) string {
	for i := 0; i < len(nums)-1; i++ {
		origin := nums[1] - nums[0]
		if nums[i] >= nums[i+1] {
			return "No"
		}
		if i > 0 {
			temp := nums[i+1] - nums[i]
			if temp >= origin {
				return "No"
			}
		}
	}
	return "Yes"
}

//func main() {
//	var n int
//	nums := []int{}
//	fmt.Scan(&n)
//	for n > 0 {
//		var num int
//		fmt.Scan(&num)
//		nums = append(nums, num)
//		n--
//	}
//	fmt.Println(isOk(nums))
//}
