package main

import (
	"fmt"
	"go_basics/out"
)

func Array() {

	a := []int{1, 2, 3, 4, 5}

	for _, v := range a {
		if v < 3 {
			out.TryFun(v)
		}
		fmt.Println(v)
	}

	fmt.Println(len(a))

}
