// package main

// import (
// 	"mycli/cmd"
// )

// func main() {
// 	cmd.Trial()
// }

// package main

// import (
// 	"flag"
// 	"fmt"
// )

// func main() {
// 	name := flag.String("name", "world", "your name")
// 	age := flag.Int("age", 0, "your age")

// 	flag.Parse()

//		fmt.Println("Hello", *name)
//		fmt.Println("Age:", *age)
//	}
package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mycli",
	Short: "My first CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rootCmd Hello!")
	},
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Print hello",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello!")
	},
}

func main() {
	rootCmd.AddCommand(helloCmd)

	rootCmd.Execute()
}
