package main

// import "fmt"

func TryPtn(n int) string {
	var row string
	for i := 0; i < n; i++ {

		for j := n - 1; j > i; j-- {
			row += "& "
		}
		for k := 0; k < i; k++ {
			row += "1 "
		}
		row += " \n"
	}
	return row
}
