package main

import (
	"fmt"
	"strings"
)

func main(){
	
	conferenceName := "Go conference"
	conferenceTickets := 66
	remainingTickets := 66
	bookings := []string{}

	greetUsers( conferenceName, conferenceTickets, remainingTickets)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		fmt.Println("Enter your first name :")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name :")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address :")
		fmt.Scan(&email)
		fmt.Println("Enter the number of tickets :")
		fmt.Scan(&userTickets)

		isVaildName := len(firstName) >=2 && len(lastName) >=2
		isVaildEmail := strings.Contains(email,"@")
		isVaildTicket := userTickets >0 && userTickets <= remainingTickets

		if isVaildName && isVaildEmail && isVaildTicket {

			remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName + " " + lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive an email confirmation at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v remaining tickets for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _,booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("First names of bookings are: %v\n", firstNames)

		if remainingTickets ==0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}
		} else {
			if !isVaildName {
				fmt.Println("First name and last name should be at least 2 characters long")
			}
			if !isVaildEmail {
				fmt.Println("Please enter a valid email address")
			}
			if !isVaildTicket {
				fmt.Println("Please enter a valid number of tickets (1-66)")
			}
		}

	} 
	
}

func greetUsers( confName string, ConfTickets int, remTickets int) {
	fmt.Printf("Welcome to %v booking application\n",confName)
}