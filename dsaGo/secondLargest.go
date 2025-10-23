package main

// import "fmt"

func SecLrg(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	fstLrg := arr[0]
	secLrg := arr[1]

	for _, num := range arr {
		if num > fstLrg {
			secLrg = fstLrg
			fstLrg = num
		} else if num < fstLrg && num > secLrg {
			secLrg = num
		}

	}

	return secLrg
}
