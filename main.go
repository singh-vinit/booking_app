package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var bookings []string //slice -> dynamic array in golang

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		var firstName string
		var lastName string
		var userTickets uint
		var userEmail string

		fmt.Print("enter your firstName: ")
		fmt.Scan(&firstName)
		fmt.Print("enter you lastName: ")
		fmt.Scan(&lastName)
		fmt.Print("enter your email: ")
		fmt.Scan(&userEmail)
		fmt.Print("enter the no of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) > 2 && len(lastName) > 2
		isValidEmail := strings.Contains(userEmail, "@")
		isValidTicket := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicket {

			fullName := firstName + " " + lastName
			bookings = append(bookings, fullName)

			remainingTickets = remainingTickets - userTickets

			fmt.Printf("Thank you %v for booking %v tickets\n", fullName, userTickets)
			fmt.Printf("You will receive a confirmation email at %v\n", userEmail)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := getFirstNames(bookings)
			fmt.Printf("These are all the bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("our conference is booked out. comeback next year.")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("either firstName or lastName is too short")
			}
			if !isValidEmail {
				fmt.Println("invalid email input")
			}
			if !isValidTicket {
				fmt.Println("invalid ticket input")
			}
		}
	}

}

func greetUsers(conferenceName string, conferenceTickets uint, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend.\n")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{} //slice
	for _, v := range bookings {
		var name = strings.Fields(v)
		firstNames = append(firstNames, name[0])
	}
	return firstNames
}
