package main

import (
	"fmt"
	"main.go/helper"
)
	
var conferenceName = "Go Conference"
const totalTickets uint  = 50
var remainTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName string
	lastName  string
	email     string
	numberOfTickets uint
}

func main() {
	

	greetUser()
	
	var firstName string
	var lastName string
	var emailAdress string
	var userTickets uint

	for {

		firstName, lastName, emailAdress, userTickets := getUserInput(firstName, lastName, emailAdress, userTickets)
		
		isvalidName, isvalidEmailAddress, isvalidTickets := helper.ValidUserInput(firstName, lastName, emailAdress, userTickets, remainTickets)

		if isvalidName && isvalidEmailAddress && isvalidTickets {

			bookTicket(userTickets, firstName, lastName, emailAdress)

			firstNames := printFirstnames()
			fmt.Printf("the whole bookings: %v\n", firstNames)

			if remainTickets == 0 {
				fmt.Printf("our tickets have booked out, please come next year")
				break
			}

		
		} else {

			if !isvalidName {
				fmt.Println("your firstname or lastname is too short")
			}

			if !isvalidEmailAddress {
				fmt.Println("your email address should contains @")
			}

			if !isvalidTickets {
				fmt.Println("your tickets is not valid")
			}
		}
	}
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have %v tickets, we remain %v are avaiable.\n", totalTickets, remainTickets)
	fmt.Println("Get your ticktes here to attend")
}

func printFirstnames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput(firstName string, lastName string, emailAdress string, userTickets uint) (string, string, string, uint) {

	fmt.Println("please enter you first name")
	fmt.Scan(&firstName)

	fmt.Println("please enter you last name")
	fmt.Scan(&lastName)

	fmt.Println("please enter you email adress")
	fmt.Scan(&emailAdress)

	fmt.Println("please enter you number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, emailAdress, userTickets
}


func bookTicket(userTickets uint, firstName string, lastName string, emailAdress string) {
	remainTickets = remainTickets - userTickets
	var userData = userData {
		firstName: firstName,
		lastName: lastName,
		email: emailAdress,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings %v\n", userData)
	fmt.Printf("Thank you %v %v booked %v tickets. you will receive confirmation email at %v\n", firstName, lastName, userTickets, emailAdress)
	fmt.Printf("%v remain tickets for %v\n", remainTickets, conferenceName)
}