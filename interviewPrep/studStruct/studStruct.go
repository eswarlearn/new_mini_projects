// package main

// import (
// 	"fmt"
// 	"strings"
// )

// type Prods struct {
// 	Prod     string
// 	Quantity int64
// }

// type StdDtl struct {
// 	Name  string
// 	Score []int
// }

// func main() {
// 	stDt := []StdDtl{}
// 	var StdNm string
// 	for {
// 		fmt.Println("enter your name")
// 		fmt.Scan(&StdNm)
// 		if strings.ToLower(StdNm) != "exit" {
// 			getMark(&stDt, StdNm)
// 		} else {
// 			tpNM, top := fndAvg(&stDt)
// 			fmt.Printf("topAvg is %v by %v\n", tpNM, top)
// 			break
// 		}
// 	}
// }

// func getMark(stnDt *[]StdDtl, nm string) {
// 	var mark int
// 	fmt.Println("enter your marl")
// 	fmt.Scan(&mark)

// 	// Check if student already exists
// 	for i := range *stnDt {
// 		if (*stnDt)[i].Name == nm {
// 			(*stnDt)[i].Score = append((*stnDt)[i].Score, mark) // ✅ append mark to existing student
// 			// fmt.Println("Updated:", (*stnDt)[i])
// 			return
// 		}
// 	}

// 	// If new student → add to slice
// 	*stnDt = append(*stnDt, StdDtl{Name: nm, Score: []int{mark}}) // ✅ fix slice creation
// 	// fmt.Println("Added:", nm, "->", mark)
// }

// func fndAvg(flDet *[]StdDtl) (string, float64) {
// 	fmt.Println("Updated:", (*flDet))
// 	var top float64
// 	var tprNm string

// 	for i := range *flDet {
// 		fmt.Println((*flDet)[i].Name)
// 		total := 0
// 		var avg float64
// 		for _, mrk := range (*flDet)[i].Score {
// 			fmt.Println(mrk)
// 			total += mrk
// 		}
// 		avg = float64(total) / float64(len((*flDet)[i].Score))
// 		if top < avg {
// 			top = avg
// 			tprNm = (*flDet)[i].Name
// 		}
// 	}

// 	return tprNm, top
// }

package main

import (
	"fmt"
	"strings"
)

type Student struct {
	Name   string
	Scores []int
}

// AddScore adds a new score to the student's record
func (s *Student) AddScore(score int) {
	s.Scores = append(s.Scores, score)
}

// Average computes the student's average score
func (s Student) Average() float64 {
	fmt.Println("Avg struct Method", s)
	if len(s.Scores) == 0 {
		return 0
	}
	total := 0
	for _, score := range s.Scores {
		total += score
	}
	return float64(total) / float64(len(s.Scores))
}

func main() {
	var students []Student

	for {
		var name string
		fmt.Print("Enter student name (or 'exit' to finish): ")
		fmt.Scan(&name)

		name = strings.TrimSpace(name)
		if strings.ToLower(name) == "exit" {
			topName, topAvg := findTopStudent(students)
			fmt.Printf("\n🏆 Top student: %s with average %.2f\n", topName, topAvg)
			break
		}

		var mark int
		fmt.Print("Enter score: ")
		fmt.Scan(&mark)

		addOrUpdateStudent(&students, name, mark)
	}
}

func addOrUpdateStudent(students *[]Student, name string, score int) {
	for i := range *students {
		if (*students)[i].Name == name {
			(*students)[i].AddScore(score)
			fmt.Printf("✅ Added score %d for %s\n", score, name)
			return
		}
	}
	// new student
	newStudent := Student{Name: name, Scores: []int{score}}
	*students = append(*students, newStudent)
	fmt.Printf("🆕 Added new student %s with first score %d\n", name, score)
}

func findTopStudent(students []Student) (string, float64) {
	if len(students) == 0 {
		return "N/A", 0
	}

	var topName string
	var topAvg float64

	for _, s := range students {
		avg := s.Average()
		if avg > topAvg {
			topAvg = avg
			topName = s.Name
		}
	}

	return topName, topAvg
}
