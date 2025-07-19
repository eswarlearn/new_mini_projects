package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("please enter year of birth ")
	scanner.Scan()
	input, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		fmt.Println("Invalid input, please enter a valid year.")
		return
	}

	fmt.Printf("Youtype: %d\n", 2024-input)
	Array()

}
