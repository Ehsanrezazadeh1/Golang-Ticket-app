package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("enter your lastname: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email: ")
		fmt.Scan(&email)
		fmt.Println("How many tickets you want? ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketsNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketsNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. you will receivie a confirmatin email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("there is %v tickets availabe!\n", remainingTickets)

			//call functions print first names
			firstName := getFirstNames(bookings)
			fmt.Printf("tickets reaming for %v\n", firstName)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is bookek out. Come back next year!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short!")
			}
			if !isValidEmail {
				fmt.Println("your email format dose not contain @ sign!")
			}
			if !isValidTicketsNumber {
				fmt.Println("number of tickets you entered is invalid!")
			}

		}

	}

}

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to our %v\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v is available!\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])

	}
	return firstNames
}
