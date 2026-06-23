package main

import (
	"fmt"
	"gogit-cli/cmd"
	"gogit-cli/models"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("provide any input")
	}
	command := os.Args[1]
	switch command {
	case "init":
		cmd.Init()
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("usage: gogit add <filename>")
			return
		}
		fileName := os.Args[2]
		err := cmd.Add(fileName)
		if err != nil {
			fmt.Println(err)
		}
	case "commit":
		if len(os.Args) < 5 {
			fmt.Println("enter proper commit message")
		}
		models.CommitFunc()

	default:
		fmt.Println("unknown commands")
	}

}
