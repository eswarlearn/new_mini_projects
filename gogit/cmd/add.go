package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Add(fileName string) error {
	if _, err := os.Stat(".gogit"); os.IsNotExist(err) {
		return fmt.Errorf("repository not initialized")
	}
	src, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer src.Close()

	fileName = filepath.Base(fileName)

	destPath := filepath.Join(
		".gogit",
		"staging",
		fileName,
	)
	dest, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dest.Close()
	_, err = io.Copy(dest, src)
	if err != nil {
		return fmt.Errorf("failed to copy data: %w", err)
	}
	fmt.Printf("added %s to staging\n", fileName)
	return nil
}
