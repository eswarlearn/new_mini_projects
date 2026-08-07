package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var (
	asScript    bool   // run the whole file at once via an interpreter
	interpreter string // interpreter to use when --script is set
	dryRun      bool   // print commands instead of running them
	stopOnError bool   // stop line-by-line execution on first failure
)

// runCmd implements: filerunner run <path>
var runCmd = &cobra.Command{
	Use:   "run [file]",
	Short: "Read a file and execute it in the terminal",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := args[0]

		if _, err := os.Stat(path); err != nil {
			return fmt.Errorf("cannot access %q: %w", path, err)
		}

		if asScript {
			return runAsScript(path)
		}
		return runLineByLine(path)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().BoolVarP(&asScript, "script", "s", false,
		"treat the file as a single script and hand it to an interpreter, instead of running it line by line")
	runCmd.Flags().StringVarP(&interpreter, "interpreter", "i", "bash",
		"interpreter used with --script (e.g. bash, sh, python3)")
	runCmd.Flags().BoolVarP(&dryRun, "dry-run", "n", false,
		"print what would run without executing anything")
	runCmd.Flags().BoolVar(&stopOnError, "stop-on-error", true,
		"stop line-by-line execution as soon as a command fails")
}

// runAsScript hands the whole file to an interpreter (e.g. `bash file.sh`)
// and streams its output live to the current terminal.
func runAsScript(path string) error {
	if dryRun {
		fmt.Printf("[dry-run] %s %s\n", interpreter, path)
		return nil
	}

	c := exec.Command(interpreter, path)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	fmt.Printf("==> running %s with %s\n", path, interpreter)
	return c.Run()
}

// runLineByLine reads the file and runs each non-empty, non-comment line
// as its own shell command, streaming output as it goes.
func runLineByLine(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue // skip blank lines and comments
		}

		if dryRun {
			fmt.Printf("[dry-run] line %d: %s\n", lineNum, line)
			continue
		}

		fmt.Printf("==> [%d] %s\n", lineNum, line)

		// Run through the shell so pipes, redirects, env vars, etc. work.
		c := exec.Command("sh", "-c", line)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Stdin = os.Stdin

		if err := c.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "line %d failed: %v\n", lineNum, err)
			if stopOnError {
				return fmt.Errorf("stopped at line %d", lineNum)
			}
		}
	}

	return scanner.Err()
}
