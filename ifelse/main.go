package main

import "fmt"

func main() {
	fmt.Println("If else")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "regular User"
	} else {
		result = "Main user"
	}

	fmt.Println(result)
}
