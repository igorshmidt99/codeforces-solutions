package prologue

import "fmt"

func Fib2(n int) int {
	nums := []int{}
	nums[0] = 0
	nums[1] = 1

	for i := 2; i <= n; i++ {
		nums[i] = nums[i - 1] + nums[i - 2]	
	}
	return nums[n]
}

func Fib3(n int) int {
	nums := [][2]int{}
	nums[0] = [2]int{0, 1} // f2
	nums[1] = [2]int{nums[0][1], nums[0][0] + nums[0][1]} // f3
	fmt.Printf("%T", nums)
	
	return 0
}
