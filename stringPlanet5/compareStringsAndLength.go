package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// 1. How to get string length?
	// The RuneCountInString(s string) (n int) function of the utf8 package
	// can be used to find the length of the string
	// This method takes a string as an arg & returns the number of runes in it

	var myString = "élémentaire"
	fmt.Printf("The string is %s \n", myString)
	fmt.Printf("Has The Length = %d \n", utf8.RuneCountInString(myString))
	fmt.Printf("Number of Bytes= %d \n\n", len(myString))

	fullName := "IRADUFASHA"
	fmt.Printf("The string is %s \n", fullName)                            // GAD IRADUFASHA
	fmt.Printf("Has The Length = %d \n", utf8.RuneCountInString(fullName)) // 14
	fmt.Printf("Number of Bytes= %d \n", len(fullName))                    // 14

	justName := "Définition"
	fmt.Printf("The string is %s \n", justName)
	fmt.Printf("Has The Length = %d \n", utf8.RuneCountInString(justName))
	fmt.Printf("Number of Bytes= %d \n", len(justName))

	leVerbe := "ÉCOLE"
	fmt.Printf("The string is %s \n", leVerbe)
	fmt.Printf("Has The Length = %d \n", utf8.RuneCountInString(leVerbe))
	fmt.Printf("Number of Bytes= %d \n", len(leVerbe))

	// How to make string comparison?
	// The == operator is used to compare two strings for equality
	// If both the strings are equal, the result is true else its false
	englishName, frenchName := "Egide", "Égide"
	fmt.Println(englishName == frenchName) // false
	compareStrings(englishName, frenchName)

	stringName1, stringName2 := "Java", "Dart"
	compareStrings(stringName1, stringName2)

	stringName3, stringName4 := "Java", "JAVA"
	compareStrings(stringName3, stringName4)

	stringName5, stringName6 := "Kotlin1", "Kotlin1"
	compareStrings(stringName5, stringName6)

}

// Gadefined function
func compareStrings(s1 string, s2 string) {
	if s1 == s2 {
		fmt.Printf("%s and %s are equal!\n", s1, s2)
		return
	} else {
		fmt.Printf("%s and %s ain't equal!\n", s1, s2)
	}
}
