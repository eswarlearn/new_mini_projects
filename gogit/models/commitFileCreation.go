package models

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func commitFileCreation() (string, error) {
	commitsDir := filepath.Join(".gogit", "commits")
	entries, err := os.ReadDir(commitsDir)
	if err != nil {
		return "", fmt.Errorf("commit directory does not exist, run 'gogit init': %w", err)
	}

	lastFolder := 0
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		n, err := strconv.Atoi(entry.Name())
		if err == nil && n > lastFolder {
			lastFolder = n
		}
	}

	commitDir := filepath.Join(commitsDir, strconv.Itoa(lastFolder+1))
	if err := os.MkdirAll(commitDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create commit folder: %w", err)
	}

	stagingDir := filepath.Join(".gogit", "staging")
	stagedEntries, err := os.ReadDir(stagingDir)
	if err != nil {
		return "", fmt.Errorf("failed to read staging directory: %w", err)
	}

	for _, entry := range stagedEntries {
		if entry.IsDir() {
			continue
		}

		source := filepath.Join(stagingDir, entry.Name())
		destination := filepath.Join(commitDir, entry.Name())
		if err := os.Rename(source, destination); err != nil {
			return "", fmt.Errorf("failed to move staged file %q: %w", entry.Name(), err)
		}
	}

	return commitDir, nil
}
