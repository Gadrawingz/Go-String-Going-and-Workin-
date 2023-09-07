package main

import "fmt"

func main() {

	// Without using Lune Datatype!
	givenString := "Monsieur Señor"
	outputCharactersNoRune(givenString)
	// Output: M o n s i e u r   S e Ã ± o r
	// Why: The unicode point for ñ is U+00F1 that occupies
	// 2 bytes c3(Ã) and b1(±)

	// By Using builtin "Lune" type aka int32, it will represent
	// a Unicode point in Go, Doesn't matter how many bytes the code points
	// Contains
	outputCharactersUseRune(givenString)

}

// Making 2 functions for difference
func outputCharactersNoRune(str string) {
	fmt.Printf("The characters: ")
	for g := 0; g < len(str); g++ {
		fmt.Printf("%c ", str[g])
	}
	fmt.Printf("\n")
}

func outputCharactersUseRune(str string) {
	fmt.Printf("The characters: ")
	myRunes := []rune(str)
	for k := 0; k < len(myRunes); k++ {
		fmt.Printf("%c", myRunes[k])
	}
	fmt.Printf("\n")
}
