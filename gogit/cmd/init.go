package cmd

import (
	"fmt"
	"os"
)

func Init() {
	if _, err := os.Stat(".gogit"); err == nil {
		fmt.Println("Repository already initialized")
		return
	}

	// err := os.Mkdir(".gogit", os.ModePerm)

	err := os.MkdirAll(".gogit/staging", os.ModePerm)

	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.MkdirAll(".gogit/commits", os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("gogit initialized")
}
