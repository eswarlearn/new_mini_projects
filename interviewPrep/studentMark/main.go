package main

import "fmt"

func main() {
	fmt.Println("Welcome to mark calculator")
	studentMarks := make(map[string][]int)
	var name string

	for {
		fmt.Println("Enter name of student or 'exit' to finish:")
		fmt.Scan(&name)
		if name == "exit" {
			fmt.Println("Exiting the program.", studentMarks)
			break
		}
		stuDetail(studentMarks, name)

	}
}

func stuDetail(adDtl map[string][]int, name string) {
	_ = adDtl
	var mark int
	fmt.Println("Enter any subject mark")
	fmt.Scan(&mark)
	if mark < 0 || mark > 100 {
		fmt.Println("enter valid mark")
	}
	if _, exist := adDtl[name]; !exist {
		adDtl[name] = []int{}
	}
	adDtl[name] = append(adDtl[name], mark)
	fmt.Println("output test", adDtl)

}
