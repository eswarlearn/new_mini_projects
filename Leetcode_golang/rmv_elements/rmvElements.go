package rmvelements

import "fmt"

func RmvEle(nums []int, val int) int {
	arryLen := len(nums)
	if arryLen == 0 {
		return -1
	}

	count := arryLen
	fmt.Println(count)
	//0,1,2,2,3,0,4,2
	for i := 0; i < arryLen; i++ {
		if nums[i] == val && (1+i) < arryLen {
			nums[i] = nums[1+i]
			nums[1+i] = val
			count--
		} else if nums[i] == val {
			count--
		}
	}
	fmt.Println(nums)

	return count
}
