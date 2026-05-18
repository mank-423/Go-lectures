package main

import "fmt"

func main() {
	fmt.Println("methods")

	// no inheritence

	hitesh := User{"hitesh", "mk@gm.com", true, 22}

	hitesh.GetStatus()
	hitesh.NewMail()

	fmt.Println(hitesh)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// method declaration
func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "new@gm.com"
	fmt.Println("Email:", u.Email)
}
