package main

import "fmt"

func main() {
	fmt.Println("Pointer file")

	// pointer
	// var ptr1 *int

	// memory address
	// fmt.Println("ptr", &ptr1)

	myNumber := 23

	var myNumberPtr = &myNumber
	
	fmt.Println("Value of pointer:", myNumberPtr);
	fmt.Println("Value of pointer:", *myNumberPtr);
}
