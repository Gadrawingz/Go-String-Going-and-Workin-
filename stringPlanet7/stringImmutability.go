package main

import "fmt"

func mutateStringFailed(str string) string {
	// Any valid unicode character within single quote is a rune
	// str[0] = 'A' ?? Error here
	return str
}

// Dealing with immutability, strings are converted to a slice of runes.
// Then that slice is mutated with whatever changes are needed and converted
// back to a new string.
func mutateStringItWorks(str []rune) string {
	str[0] = 'G' // No error as seen in mutateStringFailed function
	return string(str)
}

func mutate1stStringLetter(str []rune) string {
	str[0] = '?' // No error as seen in mutateStringFailed function
	return "Incomplete..." + string(str)
}

func main() {
	message1 := "Gappy new year!"
	message2 := "Gad On!"

	// Lemme use my functions:
	fmt.Println(mutateStringItWorks([]rune(message1)))
	// Aappy new year
	fmt.Println(mutate1stStringLetter([]rune(message2)))
	// Output: Incomplete...?ad On!

}
