package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	folderName := "mainFol/subFol"
	fileNmae := "test_file.txt"
	content := "second input is created\n"

	filePath := filepath.Join(folderName, fileNmae)

	if err := os.MkdirAll(folderName, 0755); err != nil {
		log.Fatalf("Error creating directory %s: %v\n", folderName, err)
	}
	fmt.Printf("Ensured directory '%s' exists.\n", folderName)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		log.Fatalf("Error writing to file %s: %v\n", filePath, err)
	}

	defer file.Close()

	if _, err := file.Write([]byte(content)); err != nil {
		log.Fatalf("Error appending data to file: %v\n", err)
	}

	fmt.Println("file created successfully")
}
