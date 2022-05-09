package main

import (
	"fmt"
	"strings"
)

var conferencName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {

		//Cakk get user func
		firstName, lastName, email, userTickets := getUserInput()

		//Call validate input func
		isValidName, isValidEmail, isValidUserTickets := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidUserTickets {

			//Call book func to book tickets
			bookTicket(userTickets, firstName, lastName, email)

			//Call getFirstNames function to print first name of the user booking the system
			firstNames := getFirstNames()
			fmt.Printf("Your first name is %v\n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("We are fully booked! See yah next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("First or last name is not correct \n")
			}
			if isValidName {
				fmt.Printf("Your email is wrong \n")
			}
			if isValidUserTickets {
				fmt.Printf("Number of tickets you entered is invalid\n")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking app\n", conferencName)
	fmt.Printf("We have a total of %v tickets and  %v are still available\n", conferenceTickets, remainingTickets)
}

func getFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string,
	email string, userTickets uint) (bool, bool, bool) {

	//Validate Name
	isValidName := len(firstName) >= 2 && len(lastName) >= 2

	//Validate Email
	isValidEmail := strings.Contains(email, "@")

	//Validate Number of tickets is greater than 0
	isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidUserTickets
}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//Ask user for their name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("How many tickets did you buy:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint,
	firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferencName)

}
