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

	var bookings  = [50]string{"Nana", "Nicole", "James"}
	// array must contain the number / length of the array and the type of what will go in..
	
	
	fmt.Print(bookings)
	
	// var firstName string
	// var lastName string
	// var email string
	// var userTickets uint
	// var bookingsTwo [50]string
	var bookingsThree[]uint

	bookingsThree = append(bookingsThree, 50, 20, 30)
	fmt.Println(bookingsThree)

	// bookingsTwo[0] = "Dad"
	// bookingsTwo[1] = "Mom"
	// // ask user for their name
	// fmt.Print("Enter your first name:")
	// fmt.Scan(&firstName)

	// fmt.Print("Enter your last name:")
	// fmt.Scan(&lastName)
	
	// fmt.Print("Enter your email:")
	// fmt.Scan(&email)

	// fmt.Print("Enter number of tickets:")
	// fmt.Scan(&userTickets)
	// remainingTickets = remainingTickets - userTickets

	// fmt.Printf("User %v %v booked %v tickets. You will receive a confirmation at %v\n", firstName, lastName, userTickets, email)
	// fmt.Printf("%v tickets remaining for %v\n", remainingTickets, name)

	// bookingsTwo[2] = firstName + " " + lastName

	// fmt.Println(bookingsTwo[2], "Hi")
	// fmt.Println(bookingsTwo[2], "Hi again")
	// fmt.Printf("Array length: %v\n", len(bookingsTwo))
	


}



