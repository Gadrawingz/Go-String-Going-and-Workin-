package main

import "fmt"

func main() {

	// 1st way: String Concatenation
	// The + operator is used 2 join(concatenates) strings together
	theFirstName := "Micheal"
	theLastName := "Jackson"
	fmt.Println(theFirstName + " " + theLastName)

	// 2nd way: Using the Sprintf function
	greetInitial := "Good"
	greetSuffix := "Afternoon"
	person1 := "Irving"
	person2 := "Gadrawin"

	// It formats a string according to the input format specifier
	// It then, returns the resulting string!
	var done1 = fmt.Sprintf("%s %s %s!", greetInitial, greetSuffix, person1)
	var done2 = fmt.Sprintf("%s %s %s!", greetInitial, greetSuffix, person2)
	fmt.Println(done1)
	fmt.Println(done2)

	// NB1: String is immutable
	// Once a string is created, It's impossible to change it
}
