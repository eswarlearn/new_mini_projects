package cmd

import (
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Say hello",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Hello from Cobra!")
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
