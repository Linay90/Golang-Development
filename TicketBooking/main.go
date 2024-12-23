package main

import (
	"fmt"
)

func main() {

	var conferenceName = "Go Conference" //go is statically typed it will imply the datatype of assigned value to given variable but when we are only declaring variable then we have to specify the datatype
	const conferenceTickets = 50
	var remainingTickets = 50

	// remainingTkt:=60       //shorthand syntax for variable declaration  cannot work for constants
	fmt.Println("Welcome to", conferenceName, "ticket booking app")
	fmt.Println("we have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")

	fmt.Printf("conference tickets is %T,remaining ticket is %T,conferrencename is %T\n", conferenceTickets, remainingTickets, conferenceName) //  %T is a placeholder for variable types it will automatically interpret the types of that variables.

	// fmt.Printf("Welcome to %v ticket booking app\n", conferenceName)
	// fmt.Printf("we have total of %v tickets and %v are still remaining",conferenceTickets,remainingTickets)  //make sure values you pass are in correct order as you want
	fmt.Println("Get your tickets here to attend")

	var bookings = []string{}
	//loop statements
	for {
		var userName string
		var userTickets int
		// userName = "Tom"
		// fmt.Println(userName)
		fmt.Scan(&userName) // accepting user input  if we don't use & here then the program will exit immediately without asking for input it needs the memory address for accepting input
		fmt.Println("Enter usernamem")
		fmt.Scan(userName)
		fmt.Println("Enter number of tickets")
		fmt.Scan(userTickets)

		remainingTickets = conferenceTickets - remainingTickets
		bookings[0] = userName
		bookings = append(bookings, userName)
	}

}
