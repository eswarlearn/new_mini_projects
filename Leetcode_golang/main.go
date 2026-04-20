package main

import (
	"fmt"
	"os"
	"strconv"

	arysli "github.com/eswar/project/AarySli"
	mxadis "github.com/eswar/project/mxaDis"
	// rmvdup "github.com/eswar/project/rmvdupLd"
	// rmuEle27 "github.com/eswar/project/rmv_elements"
)

func main() {
	prob26 := []int{1, 1, 1, 6, 1, 1, 1}
	// val := 2
	// re := rmvdup.Sum26(prob26)
	re := mxadis.MaxDistance(prob26)

	// re := rmuEle27.RmvEle(prob26, val)

	fmt.Println(re)

	nbr := 6

	arysli.WhatIsAry(nbr)
	fmt.Println()
	if len(os.Args) < 3 {
		fmt.Println("Provide two numbers")
		return
	}

	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])

	fmt.Println("Sum:", a+b)
}
