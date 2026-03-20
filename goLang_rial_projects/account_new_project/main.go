package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Expense struct {
	ID          int
	Amount      float32
	Category    string
	Description string
	CreatedAt   time.Time
}

var usrArayData []Expense

func main() {
	rand.Seed(time.Now().UnixNano())

	IdGeneration := rand.Intn(100)

	var userAmount float32
	fmt.Print("Enter the amount: ")
	fmt.Scan(&userAmount)

	var userCategory string
	fmt.Print("Enter the category: ")
	fmt.Scan(&userCategory)

	var description string
	fmt.Print("Enter the description: ")
	fmt.Scan(&description)

	userDef := Expense{IdGeneration, userAmount, userCategory, description, time.Now()}

	usrArayData = append(usrArayData, userDef)
	fmt.Println(usrArayData)

}
