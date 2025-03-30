package main

import (
	"BOOKING-APP/helper"
	"fmt"
	"strconv"
)

// package variable
var conferenceName = "Go Conference" //replace var and data type
const conferenceTicket = 50

var remainingTicket uint = conferenceTicket

// var booking []string
var booking = []string{} // this also used

func main() {

	fmt.Printf("global variable %v \n", helper.MrGlob)

	greeting()

	for {

		// data type
		firstName, lastName, email, userTicket := getUserDetails()
		fmt.Println(firstName, lastName, email, userTicket, "trse ")

		isValidName, isValidEmail, isUserTicker := helper.ValidateUserInput(firstName, lastName, email, userTicket, remainingTicket)

		if isValidName && isValidEmail && isUserTicker {

			bookingTicket(firstName, lastName, email, userTicket)

			fmt.Printf("our first name booking %v\n", helper.GetFirstName(booking))

			if remainingTicket == 0 {
				fmt.Printf("sorry all %v has been sold", conferenceTicket)
				break
			} else {
				fmt.Println()
			}

			fmt.Printf("our firstNm lastNm booking %v\n\n", booking)

		} else {
			fmt.Printf("input data is in valid pls check again\n")
		}
	}

}

func greeting() {
	fmt.Printf("Welcome to %v ticket booking\n", conferenceName)
	fmt.Printf("We have number %v of tickets now avaliable %v tickets\n", conferenceTicket, remainingTicket)
	fmt.Println("get yout tickets here")
}

func getUserDetails() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Print("please enter your First name : \n")
	fmt.Scan(&firstName)
	fmt.Print("please enter your last name : \n")
	fmt.Scan(&lastName)
	fmt.Print("please enter your email: \n")
	fmt.Scan(&email)
	fmt.Print("please enter number of tickets: \n")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookingTicket(firstName string, lastName string, email string, userTicket uint) {

	remainingTicket -= userTicket
	// booking[0] = firstName //array syntax
	var fullName = firstName + " " + lastName
	fmt.Printf("hi %v hom many tickets tou like to book :\n ", fullName)

	booking = append(booking, fullName)

	//map
	var useData = make(map[string]string)
	useData["firstName"] = firstName
	useData["lastName"] = lastName
	useData["email"] = email
	useData["userBookingTicket"] = strconv.FormatUint(uint64(userTicket), 10)

	fmt.Println(useData, " map in function")

	/*array syntax*/
	// fmt.Printf("full array %v \n", booking)
	// fmt.Printf("1 at array %v \n", booking[0])
	// fmt.Printf("type array %T \n", booking)
	// fmt.Printf("length of array %v \n", len(booking))

	fmt.Printf("hi %v you have boolked  %v tickets\n", firstName, userTicket)
	fmt.Printf("remaining %v tickets are available\n", remainingTicket)
}
