package main

import (
	"fmt"

	arysli "github.com/eswar/project/AarySli"
	// rmvdup "github.com/eswar/project/rmvdupLd"
	rmuEle27 "github.com/eswar/project/rmv_elements"
)

func main() {
	prob26 := []int{0,1,2,2,3,0,4,2}
	val := 2
	// re := rmvdup.Sum26(prob26)

	re := rmuEle27.RmvEle(prob26, val)

	fmt.Println(re)

	nbr := 6

	arysli.WhatIsAry(nbr)
	fmt.Println()
}
