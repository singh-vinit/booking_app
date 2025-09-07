package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend.\n")

	var userName string
	var userTickets uint

	//taking the userinput by passing the reference of variable
	fmt.Print("enter your name: ")
	fmt.Scan(&userName)
	fmt.Print("enter the no of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("username: %v and userTickets: %v", userName, userTickets)

}
