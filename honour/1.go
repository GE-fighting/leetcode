package main

func calculate(a, b int, op byte) int {
	if op == '+' {
		return a + b
	}
	return a - b
}

//func main() {
//	nums := []int{}
//	ops := []byte{}
//	var op string
//	_, err := fmt.Scan(&op)
//	if err != nil {
//		return
//	}
//	for i := 0; i < len(op); i++ {
//		if op[i] >= '0' && op[i] <= '9' {
//			num, _ := strconv.Atoi(string(op[i]))
//			if i < len(op)-1 {
//				if op[i+1] >= '0' && op[i+1] <= '9' {
//					numB, _ := strconv.Atoi(string(op[i+1]))
//					num = 10*num + numB
//					i++
//				}
//			}
//			nums = append(nums, num)
//		} else {
//			if op[i] == '=' {
//				break
//			}
//			ops = append(ops, op[i])
//		}
//		if len(nums) == 2 && len(ops) == 1 {
//			returnVal := calculate(nums[0], nums[1], ops[0])
//			nums = []int{returnVal}
//			ops = ops[1:]
//		}
//	}
//	fmt.Println(nums[0])
//}
