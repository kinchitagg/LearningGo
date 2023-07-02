package main

import (
	"fmt"
	"strings"
)

var conferenceName string = "Cloud Cube workshop"
var totalTickets uint = 50
var remainingTickets uint = 50
var totalBookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	userEmail   string
	ticketCount uint
}

func main() {

	for {

		var isBook = greetUserAndCheck()

		if isBook {

			fmt.Println("getting user data")
			firstName, lastName, userEmail, ticketCount := getUserData()
			isAllFieldValid := validateUserInput(firstName, lastName, userEmail, ticketCount)

			if isAllFieldValid {
				bookTickets(firstName, lastName, userEmail, ticketCount)

			} else {
				fmt.Println("Your input is wrong kindly start again !")
				fmt.Println("Bye Bye")
				fmt.Println("*******************************************")
				continue
			}

		} else {

			fmt.Println("Thanks for Visit. We expect you in our next workshop")

		}

	}

}

func greetUserAndCheck() bool {

	fmt.Println("##############################################")
	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Printf("Attend our workshop on platform engineering\n")
	fmt.Printf("Hurry!! only %v/%v ticket left !\n", remainingTickets, totalTickets)
	fmt.Println("##############################################")

	var isBook bool
	var userBookValue string
	for {

		fmt.Printf("Would like to book our tickets (y/n)")
		fmt.Scan(&userBookValue)

		if userBookValue == "y" {
			isBook = true
			return isBook
		} else if userBookValue == "n" {
			isBook = false
			return isBook
		} else {
			fmt.Print("wrong input kindly enter again \n\n")

		}
	}

}

func getUserData() (string, string, string, uint) {

	var firstName string
	var lastName string
	var userEmail string
	var ticketCount uint
	fmt.Println("Thanks for Booking ! Kindly enter user data")
	fmt.Println("Enter your first name.")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name.")
	fmt.Scan(&lastName)
	fmt.Println("Enter your valid email.")
	fmt.Scan(&userEmail)
	fmt.Println("Enter your ticket count.")
	fmt.Scan(&ticketCount)
	return firstName, lastName, userEmail, ticketCount

}

func validateUserInput(firstName string, lastName string, userEmail string, ticketCount uint) bool {

	var isAllFieldValid bool
	isAllFieldValid = true

	if len(firstName) < 2 || len(lastName) < 2 {
		fmt.Println("Name is too short")
		isAllFieldValid = false
	}

	if !strings.Contains(userEmail, "@") {
		fmt.Println("Not a valid email address")
		isAllFieldValid = false
	}
	if ticketCount <= 0 {
		fmt.Println("Invalid ticket count")
		isAllFieldValid = false
	}
	if ticketCount > remainingTickets {
		fmt.Printf("Sorry we have only %v tickets left\n", remainingTickets)
		isAllFieldValid = false
	}

	if isAllFieldValid {
		return true
	} else {
		return false
	}

}

func bookTickets(firstName string, lastName string, userEmail string, ticketCount uint) {

	remainingTickets = remainingTickets - ticketCount
	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		userEmail:   userEmail,
		ticketCount: ticketCount,
	}

	totalBookings = append(totalBookings, userData)

	fmt.Printf("Thank you %v %v for booking. we will send confirmation of %v ticket at %v", firstName, lastName, ticketCount, userEmail)
}
