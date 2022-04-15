package main

import (
	"fmt"
)

func main() {
	name := "Go Conference"
	const conferenceTickets = -50
	var remainingTickets uint = 47

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, name is %T \n", conferenceTickets, remainingTickets, name)

	fmt.Printf("Welcome to our %v booking application\n", name)
	fmt.Printf("We have a total of %v tickets and only %v left!\n", conferenceTickets, remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Print("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name:")
	fmt.Scan(&lastName)
	
	fmt.Print("Enter your email:")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets:")
	fmt.Scan(&userTickets)
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("User %v %v booked %v tickets. You will receive a confirmation at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, name)

}



