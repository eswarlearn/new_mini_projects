package main

import (
	"fmt"
	outTst "go_basics/outAs"
)

func Array() {

	a := []int{1, 2, 3, 4, 5, 5}

	for _, v := range a {
		if v < 3 {
			outTst.TryFun(v)
		}
		fmt.Println(v)
	}

	fmt.Println(len(a))
	outTst.TryFun(10)("example")(3)

}
