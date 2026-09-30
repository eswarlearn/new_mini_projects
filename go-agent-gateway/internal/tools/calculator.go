package tools

import "strconv"

type CalculatorTool struct {
	num int
}

func (c *CalculatorTool) Name() string {
	return "calculator"
}

func (c *CalculatorTool) Execute(input string) (string, error) {
	c.num = 5

	return "Calculator executed with input: " +
		input +
		" struct " +
		strconv.Itoa(c.num), nil
}
