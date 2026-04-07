package main

import (
	"fmt"

	rmvdup "github.com/eswar/project/rmvdupLd"
)

func main() {
	prob26 := []int{0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 4, 4, 8}
	re := rmvdup.Sum26(prob26)

	fmt.Println(re)
}
