package main

import (
	"fmt"
	"sync"
	"time"
)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	noOfTickets uint
}

const conferenceTickets uint = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings []UserData

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, userEmail, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicket := validateUserInput(firstName, lastName, userEmail, userTickets)

	if isValidName && isValidEmail && isValidTicket {

		bookTicket(firstName, lastName, userEmail, userTickets)
		wg.Add(1)
		go sendTicket(firstName, lastName, userEmail, userTickets)

		firstNames := getFirstNames()
		fmt.Printf("These are all the bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("our conference is booked out. comeback next year.")
			// break
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
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend.\n")
}

func getFirstNames() []string {
	firstNames := []string{} //slice
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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
	return firstName, lastName, userEmail, userTickets
}

func bookTicket(firstName string, lastName string, userEmail string, userTickets uint) {

	fullName := firstName + " " + lastName

	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       userEmail,
		noOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v for booking %v tickets\n", fullName, userTickets)
	fmt.Printf("You will receive a confirmation email at %v\n", userEmail)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(firstName string, lastName string, userEmail string, userTickets uint) {
	time.Sleep(10 * time.Second)
	var generatedTicket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("-------------------------")
	fmt.Printf("sending ticket: \n %v \n to email address: %v \n", generatedTicket, userEmail)
	fmt.Println("-------------------------")
	wg.Done()
}
