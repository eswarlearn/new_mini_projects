package main

import (
	"fmt"
)

func main() {
	num := 5

	recRsn(num)
}

func recRsn(n int) {
	if n <= 0 {
		return
	}
	recRsn(n - 1)
	fmt.Printf("%d", n)
	for i := 0; i < n; i++ {
		fmt.Print("#")
	}
	fmt.Println()

}
