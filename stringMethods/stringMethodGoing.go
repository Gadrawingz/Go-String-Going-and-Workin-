package main

import (
	"fmt"
	"strings" // Must be imported
)

func main() {

	// 1. Compare two strings!
	myString1 := "Gad"
	myString2 := "Gorilla"
	myString3 := "Gad"

	// -1 means that myString1 is smaller than myString2
	fmt.Println(strings.Compare(myString1, myString2))

	// 1 means that myString1 is greater than myString2
	fmt.Println(strings.Compare(myString2, myString3))

	// 0 means that myString1 is equal to myString3
	fmt.Println(strings.Compare(myString1, myString3))

	/* How to check if String contains Substring? */
	// To check if a substring is present inside a string, we use the
	// Contains() method of the Go strings package?
	myText := "All Programming Language"
	mySubstring1 := "Program"
	mySubstring2 := "Programmable"
	mySubstring3 := "Programming"

	// Checking if "Program, Programmable,..." present in "myText" variable
	fmt.Println(strings.Contains(myText, mySubstring1)) // true
	fmt.Println(strings.Contains(myText, mySubstring2)) // false
	fmt.Println(strings.Contains(myText, mySubstring3)) // true

	// ==========================================
	// 3. How to replace a String in Go?
	// We use the Replace() method present inside the strings package
	var movieName = "Batman"
	fmt.Println(movieName) // Batman

	// 1st is old char, 2nd is new char, 3rd represent how many old char 2be replaced
	fmt.Println(strings.Replace(movieName, "B", "R", 1)) // Batman
	fmt.Println(strings.Replace(movieName, "B", "C", 1)) // Catman

	// 4. What if we need to replace multiple characters?
	word1 := "GROUNDING"
	word2 := "SALSA"
	fmt.Println(strings.Replace(word1, "G", "X", 2)) // XROUNDINX
	fmt.Println(strings.Replace(word2, "S", "P", 1)) // PALPA

	// 5. How to change the case of Go string?
	fmt.Println(strings.ToLower(word1))    // grounding (Convert to uppercase)
	fmt.Println(strings.ToUpper("eminem")) // EMINEM (Convert to lowercase)

	// 6. How to Split Strings in GoLang?
	// We can split a string into multiple substrings using Split function!
	// This method returns a slice of all the substrings
	greatMessage := "Welcome To Go Land Of Abraham!"
	fmt.Println(greatMessage)
	// We split the string at " " to individual words as output
	splitted := strings.Split(greatMessage, " ")
	fmt.Println(splitted)    // [Welcome To Go Land Of Abraham!]
	fmt.Println(splitted[0]) // Welcome
	fmt.Println(splitted[3]) // Land
}
