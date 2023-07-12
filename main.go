package main

import "fmt"


func main () {
	conferenceName := "GopherCon"
	const conferenceTickets uint = 50
	var attendeeNumber uint= 0

	// Array of strings
	bookings := []string{}

	fmt.Printf("Welcome to the %v conference 2023.\n",conferenceName)
	fmt.Printf("There are %v tickets available.\n",conferenceTickets-attendeeNumber)
	fmt.Printf("RSVP to all the workshops and talks you would like to attend!\n")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("What is your first name: ")
	fmt.Scan(&firstName)
	fmt.Printf("What is your last name: ")
	fmt.Scan(&lastName)
	fmt.Printf("What is your email: ")
	fmt.Scan(&email)
	fmt.Printf("How many tickets would you like to reserve: ")
	fmt.Scan(&userTickets)
 
	bookings= append(bookings, firstName+" "+ lastName)

	
	attendeeNumber = attendeeNumber + userTickets

	fmt.Printf("%v %v has reserved %v tickets for the conference. An email was sent with your ticket to %v\n",firstName,lastName,userTickets,email)
	fmt.Printf("Only %v tickets are now available.\n",conferenceTickets-attendeeNumber)

}