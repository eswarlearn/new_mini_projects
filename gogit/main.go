package main

import (
	"fmt"
	"gogit-cli/cmd"
	"gogit-cli/models"
	"os"
	"strings"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("usage: gogit <init|add|commit>")
	}

	switch args[0] {
	case "init":
		return cmd.Init()
	case "add":
		if len(args) != 2 {
			return fmt.Errorf("usage: gogit add <filename>")
		}
		return cmd.Add(args[1])
	case "commit":
		message, err := parseCommitMessage(args[1:])
		if err != nil {
			return err
		}
		return models.CommitFunc(message)
	default:
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func parseCommitMessage(args []string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("usage: gogit commit -m \"message\"")
	}

	// Accept Git-style commit messages and a plain message to keep the learning CLI friendly.
	if args[0] == "-m" || args[0] == "--message" {
		if len(args) < 2 {
			return "", fmt.Errorf("commit message cannot be empty")
		}
		return strings.TrimSpace(strings.Join(args[1:], " ")), nil
	}

	return strings.TrimSpace(strings.Join(args, " ")), nil
}
