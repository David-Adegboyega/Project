package main

import (
	"fmt"
	"strings"
)

func main() {
	//  Welcome message to the GO conference

	const MN string = "Go conference"
	const CT uint = 50
	var RT uint = 50
	var bookings []string

	fmt.Println("welcome to", MN, "booking application😊")
	fmt.Println("It so nice to have you around")
	fmt.Println("we have a total of", CT, "tickets and", RT, "tickets are still available")
	fmt.Println("Register here to attend")

	for {
		// getting User input

		var firstname string
		var lastname string
		var email string
		var userTickets uint

		fmt.Println("Enter your First Name")
		fmt.Scan(&firstname)

		fmt.Println("Enter your Last Name")
		fmt.Scan(&lastname)

		fmt.Println("Enter your Email address")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets you want to purchase")
		fmt.Scan(&userTickets)

		// No of ticket Remaining

		RT = RT - userTickets

		// list of attendees for the conference

		bookings = append(bookings, firstname+" "+lastname)

		fmt.Printf("Thank you %v %v for booking %v ticket(s). you will receive a confirmation mail via %v \n", firstname, lastname, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", RT, MN)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstname := names[0]
			firstNames = append(firstNames, firstname)
		}
		fmt.Printf("The first name of bookings are: %v\n", firstNames)
		// ending the program
		if RT == 0 {
			fmt.Printf("Alas!😮‍💨  \n")
			fmt.Printf("Ticket no longer available \n")
			break
		}
	}
}
