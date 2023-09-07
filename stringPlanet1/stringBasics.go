package main

import "fmt"

func main() {

	// 1. String making @gadiradufasha
	developerName := "Gad Iradufasha"
	fmt.Println(developerName)
	fmt.Println()

	// 1.B We can also represent a String using a backtick
	newString := `I need to work at Google`
	fmt.Println(newString)

	// 2. Accessing individual bytes of a string
	firstName := "IRADUFASHA"

	// %s is the format specifier to print a string.
	// len(s) returns the number of bytes in the string.
	// %x is the format specifier for hexadecimal.

	fmt.Printf("String : %s \n", firstName) // String: Iradufasha
	printStringBytes(firstName)
	fmt.Printf("\n")
	printStringCharacters(firstName)

	// 3. Accessing individual characters of a string:
	name := "Hello World"
	// This firstName[0] returns the 1st character of the string: I
	fmt.Printf("\n%c \n", firstName[0])
	// This firstName[4] returns the 4th character of the string: U
	fmt.Printf("%c \n", firstName[4])

	fmt.Printf("String: %s\n", name)
	printStringCharacters(name)
	fmt.Printf("\n")
	printStringBytes(name)
	fmt.Printf("\n\n")
	name = "Señor"
	fmt.Printf("String: %s\n", name)
	printStringCharacters(name)
	fmt.Printf("\n")
	printStringBytes(name)

}

// Gadefined function
func printStringBytes(myString string) {
	fmt.Printf("The bytes : ")
	for a := 0; a < len(myString); a++ {
		fmt.Printf("%x \n", myString[a])
	}
}

func printStringCharacters(myString string) {
	fmt.Printf("The Characters : ")
	for a := 0; a < len(myString); a++ {
		fmt.Printf("%c ", myString[a])
	}
}
