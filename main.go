package main

// go mod init folder name for initialization (in our case it is go mod init Go)

// fmt -> collection of source files like Print, Scan, Printf, Println, etc
// strings -> collection of source files like Contains, Fields, etc
import (
	"fmt"
	"strings"
)

var conferenceName = "Go Conference" 

func greet(firstName string) {
	fmt.Printf("Hello %v, Welcome to the %v\n", firstName, conferenceName)
}

func validateInput(firstName string, lastName string, email string) (bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	return isValidName, isValidEmail   // can return multiple values in go
}

func main() {
	// fmt.Print("Hello");  // use Prinln for new line  
	// short form conferenceName := "Go Conference"  // use := for short form of var  (but can't use for constant values)
	const conferenceTickets = 50  // use const for constant values
	var remainingTickets uint = 50  // use uint if you don't want negative values

	// can use Println as well then you don't need to write \n explicitly
	// fmt.Println("Welcome to", conferenceName, "\n")  // use Printf for formatted string

	fmt.Printf("conferentTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)  // for getting the datatypes of variables 

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// you have to give data type for the variable if you are not assigning any value to it (not initializing at the time of declaration)
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	
	
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)   // for taking input from user
	
	fmt.Println("Enter you last name: ")
	fmt.Scan(&lastName)
	
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	
	fmt.Println("Enter number of tickets you want to book: ")
	fmt.Scan(&userTickets)

	// user input validation
	isValidName, isValidEmail := validateInput(firstName, lastName, email)
	fmt.Println(isValidName, isValidEmail)
	
	remainingTickets = remainingTickets - userTickets
	
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	var bookings = [50]string{}
	// alternative syntax for empty array ->  var bookings = [50]string
	bookings[0] = firstName + " " + lastName
	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first element of the array: %v\n", bookings[0])
	fmt.Printf("Array type %T\n", bookings)
	fmt.Printf("Array length %v\n", len(bookings))

	// slice is a dynamic array

	var dynamicBookings = []string{}
	dynamicBookings = append(dynamicBookings, firstName + " " + lastName)
	// alternative syntax for empty slice ->  var dynamicBookings = []string
	
	fmt.Printf("The whole slice: %v\n", dynamicBookings)
	fmt.Printf("Slice type %T\n", dynamicBookings)
	fmt.Printf("Slice length %v\n", len(dynamicBookings))

	// loops in go (only for loop is there in go)

	// for {
	// 	// whatever inside get's repeated until break is called
	// 	fmt.Println("This is an infinite loop")
	// }

	firstNames := []string{}
	for _, booking := range dynamicBookings {
		var names = strings.Fields(booking)
		var firstName = names[0]
		firstNames = append(firstNames, firstName)
	}

	fmt.Printf("The first names of our bookings are: %v\n", firstNames)

	// if-else in go

	var age uint
	fmt.Println("Enter you age: ")
	fmt.Scan(&age)

	if age < 18 {
		fmt.Println("You are not eligible to attend the conference")
	} else if age >= 18 && age <= 60 {
		fmt.Println("You are eligible to attend the conference")
	} else {
		fmt.Println("You are not eligible to attend the conference")
	}

	// Functions in Go
	greet(firstName)
}