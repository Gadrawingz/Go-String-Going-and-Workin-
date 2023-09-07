package main

import "fmt"

func main() {
	bigCompany := "Samsung"
	charsAndBytePosition(bigCompany)
	fmt.Println("--------------------------")

	schoolName := "New Apple"
	charsAndBytePosition(schoolName)
}

// A fx prints position of the byte where the rune starts along with rune
func charsAndBytePosition(myStr string) {
	for index, rune := range myStr {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}
