package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is the base command when filerunner is called with no subcommands.
var rootCmd = &cobra.Command{
	Use:   "filerunner",
	Short: "filerunner reads a file and executes its contents in the terminal",
	Long: `filerunner is a small CLI tool that reads a file (a shell script or
a plain list of shell commands, one per line) and runs it/them in your terminal,
streaming stdout and stderr live as the command runs.`,
}

// Execute adds all child commands to the root command and sets flags
// appropriately. This is called by main.main(). It only needs to happen
// once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Global flags could go here, e.g. rootCmd.PersistentFlags().Bool(...)
}
