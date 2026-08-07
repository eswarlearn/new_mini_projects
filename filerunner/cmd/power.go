package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var deleteAfter bool

// powerCmd implements: filerunner power <file>
//
// The file is a simple key=value format, e.g.:
//
//	action=shutdown
//	password=yourpassword
//
// action must be "shutdown" or "restart".
var powerCmd = &cobra.Command{
	Use:   "power [file]",
	Short: "Read an action + password from a file and shut down or restart the machine",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]

		action, password, err := parsePowerFile(path)
		if err != nil {
			return err
		}

		if deleteAfter {
			defer func() {
				if rmErr := os.Remove(path); rmErr != nil {
					fmt.Fprintf(os.Stderr, "warning: could not delete %s: %v\n", path, rmErr)
				} else {
					fmt.Printf("deleted %s\n", path)
				}
			}()
		}

		var shutdownArgs []string
		switch action {
		case "shutdown":
			shutdownArgs = []string{"shutdown", "-h", "now"}
		case "restart", "reboot":
			shutdownArgs = []string{"shutdown", "-r", "now"}
		default:
			return fmt.Errorf("unknown action %q (expected \"shutdown\" or \"restart\")", action)
		}

		if dryRun {
			fmt.Printf("[dry-run] sudo -S %s   (password read from %s, not printed)\n",
				strings.Join(shutdownArgs, " "), path)
			return nil
		}

		return runWithSudoPassword(password, shutdownArgs)
	},
}

func init() {
	rootCmd.AddCommand(powerCmd)
	powerCmd.Flags().BoolVarP(&dryRun, "dry-run", "n", false,
		"print what would run without executing anything")
	powerCmd.Flags().BoolVar(&deleteAfter, "delete-after", true,
		"delete the password file after reading it, whether or not the command succeeds")
}

// parsePowerFile reads simple key=value lines and returns action + password.
func parsePowerFile(path string) (action, password string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", "", fmt.Errorf("cannot open %q: %w", path, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		switch key {
		case "action":
			action = val
		case "password":
			password = val
		}
	}
	if err := scanner.Err(); err != nil {
		return "", "", err
	}
	if action == "" {
		return "", "", fmt.Errorf("file is missing an \"action=\" line")
	}
	if password == "" {
		return "", "", fmt.Errorf("file is missing a \"password=\" line")
	}
	return action, password, nil
}

// runWithSudoPassword runs `sudo -S <args...>`, feeding the password to
// sudo's stdin (the -S flag tells sudo to read the password from stdin
// instead of prompting the terminal directly).
func runWithSudoPassword(password string, args []string) error {
	full := append([]string{"-S"}, args...)
	c := exec.Command("sudo", full...)
	c.Stdin = strings.NewReader(password + "\n")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	fmt.Printf("==> running: sudo %s\n", strings.Join(args, " "))
	return c.Run()
}
