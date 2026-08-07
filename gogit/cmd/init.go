package cmd

import (
	"fmt"
	"os"
)

func Init() error {
	if _, err := os.Stat(".gogit"); err == nil {
		fmt.Println("Repository already initialized")
		return nil
	}

	if err := os.MkdirAll(".gogit/staging", 0755); err != nil {
		return fmt.Errorf("failed to create staging directory: %w", err)
	}
	if err := os.MkdirAll(".gogit/commits", 0755); err != nil {
		return fmt.Errorf("failed to create commits directory: %w", err)
	}

	fmt.Println("gogit initialized")
	return nil
}
