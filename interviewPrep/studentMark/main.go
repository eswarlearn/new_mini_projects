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
			firstCan, mrk := topAvg(studentMarks)
			fmt.Println("Top average : ", mrk, " by student : ", firstCan)
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

func topAvg(stdDta map[string][]int) (string, float64) {
	var toperName string
	var toperAvg float64
	for name, mark := range stdDta {
		sum := 0
		var avg float64
		for _, value := range mark {
			sum += value
		}
		avg = float64(sum) / float64(len(mark))

		if avg > toperAvg {
			toperAvg = avg
			toperName = name
		} else if avg == toperAvg {
			toperName += "or" + name
		}

	}

	return toperName, toperAvg
}
