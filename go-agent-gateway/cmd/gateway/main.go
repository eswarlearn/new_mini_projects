package main

import (
	"fmt"

	"go-agent-gateway/internal/tools"
)

func main() {
	registry := tools.NewRegistry()

	fmt.Println("first test es: ", registry)

	calculator := &tools.CalculatorTool{} //am i calling the struct , if so what is the use of this?

	fmt.Println("second test es: ", calculator)

	registry.Register(calculator)

	tool, exists := registry.Get("calculator")

	if !exists {
		fmt.Println("Tool not found")
		return
	}

	result, err := tool.Execute("10 + 20")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(tool.Name())
	fmt.Println(result)
}
