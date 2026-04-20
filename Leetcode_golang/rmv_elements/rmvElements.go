package rmvelements

import "fmt"

func RmvEle(nums []int, val int) int {
	arryLen := len(nums)
	if arryLen == 0 {
		return -1
	}

	count := 0
	fmt.Println(count)
	//0,1,2,2,3,0,4,2
	for i := 0; i < arryLen; i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		} 
	}
	fmt.Println(nums)

	return count
}
