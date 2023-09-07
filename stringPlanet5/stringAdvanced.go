package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// How to get string length?
	// The RuneCountInString(s string) (n int) function of the utf8 package
	// can be used to find the length of the string
	// This method takes a string as an arg & returns the number of runes in it

	var myString = "Hold On!"
	fmt.Printf("The string is %s \n", myString)
	fmt.Printf("Has The Length = %d \n", utf8.RuneCountInString(myString))
	fmt.Printf("Number of Bytes= %d \n\n", len(myString))

	fullName := "GAD IRADUFASHA"
	fmt.Printf("The string is %s \n", fullName)                            // GAD IRADUFASHA
	fmt.Printf("Has The Length = %d \n", utf8.RuneCountInString(fullName)) // 14
	fmt.Printf("Number of Bytes= %d \n", len(fullName))                    // 14

	justName := "Être tablissement où élémentaire Définition de ÉCOLE"
	fmt.Printf("The string is %s \n", justName)                            // GAD IRADUFASHA
	fmt.Printf("Has The Length = %d \n", utf8.RuneCountInString(justName)) // 14
	fmt.Printf("Number of Bytes= %d \n", len(justName))                    // 14

}
