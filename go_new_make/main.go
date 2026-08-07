package main

import (
	"fmt"
	"os"
	"strings"
)

type Items struct {
	ID       int
	Name     string
	Quantity int64
}

func main() {
	userInput := os.Args
	fmt.Println(len(userInput))
	if len(userInput) < 2 {
		fmt.Println("enter the input go run . {give your requirement and item number / name }")
	}

	caseTst := strings.ToLower(userInput[1])
	// fmt.Println(caseTst)
	switch caseTst {
	case "add":
		fmt.Println("test case")
	}
}
