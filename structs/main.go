package main

import "fmt"

func main() {
	fmt.Println("structs")

	// no inheritence

	hitesh := User{"hitesh", "mk@gm.com", true, 22};

	fmt.Println("hitesh details are: %+v\n ", hitesh);

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
