package models

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Commit struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Files     []string  `json:"files"`
}

func CommitFunc(message string) error {
	if message == "" {
		return fmt.Errorf("commit message cannot be empty")
	}

	files, err := getFilesName()
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return fmt.Errorf("nothing to commit; add files first")
	}

	commitID, err := randomID(8)
	if err != nil {
		return err
	}

	commitDir, err := commitFileCreation()
	if err != nil {
		return err
	}

	commit := Commit{
		ID:        commitID,
		Message:   message,
		Timestamp: time.Now().UTC(),
		Files:     files,
	}
	data, err := json.MarshalIndent(commit, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode commit metadata: %w", err)
	}

	// Keep metadata with the snapshot so every commit remains inspectable later.
	if err := os.WriteFile(filepath.Join(commitDir, "commit.json"), data, 0644); err != nil {
		return fmt.Errorf("failed to write commit metadata: %w", err)
	}

	fmt.Printf("committed %d file(s) as %s\n", len(files), commitID)
	return nil
}
