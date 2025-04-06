package helper

import (
	"BOOKING-APP/models"
	"strings"
)

var MrGlob = "global"

func GetFirstName(booking []models.UserData) []string {
	//string logic
	artFirstNames := []string{}
	for _, bkng := range booking {
		// fmt.Printf("index value %v\n", index)
		// fmt.Printf("Index: %v, Booking: %v\n", index, bkng)
		// var fNm = strings.Fields(bkng)
		// artFirstNames = append(artFirstNames, fNm[0])
		artFirstNames = append(artFirstNames, bkng.FirstName)
	}
	return artFirstNames

}

func ValidateUserInput(firstName string, lastName string, email string, userTicket uint, remainingTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isUserTicker := userTicket > 0 && userTicket <= remainingTicket

	return isValidName, isValidEmail, isUserTicker //return multiple value
}
