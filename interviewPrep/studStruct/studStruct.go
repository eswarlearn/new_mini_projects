package main

import (
	"fmt"
	"strings"
)

type Prods struct {
	Prod     string
	Quantity int64
}

type StdDtl struct {
	Name  string
	Score []int
}

func main() {
	stDt := []StdDtl{}
	var StdNm string
	for {
		fmt.Println("enter your name")
		fmt.Scan(&StdNm)
		if strings.ToLower(StdNm) != "exit" {
			getMark(&stDt, StdNm)
		} else {
			fndAvg(&stDt)
		}
	}
}

func getMark(stnDt *[]StdDtl, nm string) {
	var mark int
	fmt.Println("enter your marl")
	fmt.Scan(&mark)

	// Check if student already exists
	for i := range *stnDt {
		if (*stnDt)[i].Name == nm {
			(*stnDt)[i].Score = append((*stnDt)[i].Score, mark) // ✅ append mark to existing student
			fmt.Println("Updated:", (*stnDt)[i])
			return
		}
	}

	// If new student → add to slice
	*stnDt = append(*stnDt, StdDtl{Name: nm, Score: []int{mark}}) // ✅ fix slice creation
	fmt.Println("Added:", nm, "->", mark)
}

func fndAvg(flDet *[]StdDtl) {

}
