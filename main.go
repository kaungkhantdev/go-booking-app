package main

import (
	"fmt"
	"strings"
)


func booking() {

	conferenceName := "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking applications\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	
	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName);

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		fmt.Println("Enter your ticket: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {

			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets.\n", remainingTickets, userTickets)
			break;
		}

		remainingTickets = remainingTickets - userTickets;
		bookings = append(bookings, firstName + " " + lastName);

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n ", remainingTickets, conferenceName);

		var firstNames = []string{}
		for _, booking := range bookings {
			name := strings.Fields(booking);
			firstNames = append(firstNames, name[0]);
		}

		fmt.Printf("The first name of bookings are: %v\n", firstNames);

		if remainingTickets == 0 {
			fmt.Print("Our conference is booked out. Come back next year.")
			break;
		}
	}
}

func main() {
	booking()
}