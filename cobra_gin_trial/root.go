package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bashenv-server",
	Short: "A Gin server configured entirely from a bash rc file",
	Long: `bashenv-server is a small demo CLI (built with Cobra) that reads
configuration by sourcing a bash file such as ~/.bashrc, then starts a
Gin HTTP server using whatever values it finds there (PORT, HOST,
GIN_MODE, APP_NAME).`,
}

// Execute is called by main.main(). It only needs to happen once.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
