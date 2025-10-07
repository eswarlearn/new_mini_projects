package out

import "fmt"

func test3(myFun func(int) int) {
	fmt.Println("test3:", myFun(5))
}

func returnFun(x string) func() {
	return func() { fmt.Println("returnFun:", x) }
}

func TryFun(v int) func(string) func(int) {
	fmt.Printf("tryFun: %d\n", v)

	test2 := func(num2 int) int {
		return num2 + v
	}

	test3(test2)

	returnFun("function returned")()

	Test := func(nm string) func(int) {
		fmt.Println("nested function")

		doublrNext := func(num int) {
			fmt.Println("doublrNext:", num*2)
		}
		return doublrNext
	}

	return Test

}
