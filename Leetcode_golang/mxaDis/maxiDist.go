package mxadis

import "fmt"

func MaxDistance(colors []int) int {
	maxHs := len(colors)

	fmt.Println(maxHs)
	longDis := 0
	for i := 0; i < maxHs; i++ {
		for j := maxHs - 1; j > i; j-- {
			fmt.Println("index: ", maxHs-1, " value: ", colors[maxHs-1])
			if colors[i] != colors[j] && longDis < (j-i) {
				fmt.Println("breakkk index: ", maxHs-1, " value: ", colors[maxHs-1])
				longDis = j - i
				break
			}
		}
	}
	return longDis
}
