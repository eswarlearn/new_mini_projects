package main

import (
	"fmt"
	"math/rand"
	"time"
)

type indDetail struct {
	ID          int
	Amount      float32
	Category    string
	Description string
	CreatedAt   time.Time
}

var fullDetail []indDetail

func main() {
	rand.Seed(time.Now().UnixNano())

	IdGeneration := rand.Intn(100)

	var getAmount float32
	fmt.Println("Enter ID :")
	fmt.Scan(&getAmount)

	var cat string
	fmt.Println("Enter Category:")
	fmt.Scan(&cat)

	var dis string
	fmt.Println("enter discription:")
	fmt.Scan(&dis)

	usrDet := indDetail{IdGeneration, getAmount, cat, dis, time.Now()}

	fullDetail = append(fullDetail, usrDet)

	fmt.Println(fullDetail)

}
