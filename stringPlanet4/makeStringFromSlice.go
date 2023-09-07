package main

import "fmt"

func main() {

	// 1. How to create a string from a slice of Bytes?

	// theByteSlice contains the UTF-8 Encoded hex bytes of the string
	theByteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	gadSlice := []byte{0x47, 0x40, 0x44}
	myString1 := string(theByteSlice) // Café
	myString2 := string(gadSlice)     // G@D
	fmt.Println(myString1)
	fmt.Println(myString2)

	// 2. How to Create string from a slice of runes?
	theLuneSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	myString := string(theLuneSlice)
	fmt.Println(myString) // Señor

}
