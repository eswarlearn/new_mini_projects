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

	info, err := os.Stat(fileName)
	if err != nil {
		return fmt.Errorf("failed to read source file: %w", err)
	}
	if info.IsDir() {
		return fmt.Errorf("cannot add directory %q", fileName)
	}

	src, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer src.Close()

	stagingDir := filepath.Join(".gogit", "staging")
	if err := os.MkdirAll(stagingDir, 0755); err != nil {
		return fmt.Errorf("failed to create staging directory: %w", err)
	}

	stagedName := filepath.Base(fileName)
	destPath := filepath.Join(stagingDir, stagedName)
	dest, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dest.Close()

	if _, err := io.Copy(dest, src); err != nil {
		return fmt.Errorf("failed to copy data: %w", err)
	}

	fmt.Printf("added %s to staging\n", stagedName)
	return nil
}
