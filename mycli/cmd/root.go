package cmd
import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use : "myCli",
	Short: "My first Cobra CLI",
} 

func Trial(){
	err := rootCmd.Execute()
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}