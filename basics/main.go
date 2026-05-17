package main

import "fmt"

const loginToken string = "awfaw";

func main() {
	var username string = "Mayank"
	fmt.Println("Hey buddy", username)
	fmt.Printf("Varible if of type: %T \n", username)

	var isUsername bool = true
	fmt.Printf("Varible if of type: %T \n", isUsername)

	var smallVal uint = 256
	fmt.Printf("Varible if of type: %T \n", smallVal)

	var website = "ealrcode"
	fmt.Println("val:", website)

	numberOfUser := 3000
	fmt.Println("no:", numberOfUser)
}
