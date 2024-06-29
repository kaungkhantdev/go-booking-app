package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go conference"
const conferenceTickets int = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func booking() {

	greeting()
	


	firstName, lastName, email, userTickets := userInput()
	
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets);
	
	if isValidEmail && isValidName && isValidTicketNumber {
		
		bookingTickets(userTickets, firstName, lastName, email)
		
		wg.Add(1)
		go sendTickets(userTickets, firstName, lastName, email)

		firstNames := printFirstNames();
		
		fmt.Printf("The first name of bookings are: %v\n", firstNames);

		if remainingTickets == 0 {
			fmt.Print("Our conference is booked out. Come back next year.")
		}

		
	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
	}

	wg.Wait()
}

func greeting() {
	fmt.Printf("Welcome to %v booking applications\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
}

func userInput() (string, string, string, uint) {
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
	return firstName, lastName, email, userTickets
}



func printFirstNames() []string {
	var firstNames = []string{}
	for _, booking := range bookings {
	
		firstNames = append(firstNames, booking.firstName);
	}

	return firstNames;

}

func bookingTickets( userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets;

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData);
	fmt.Printf("List of bookings is %v\n", bookings);

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n ", remainingTickets, conferenceName);

}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}

func main() {
	booking()
}