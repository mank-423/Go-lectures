package main

import "fmt"

func main() {

	// LIFO structure to print: three, two, one
	defer fmt.Println("one");
	defer fmt.Println("two");
	defer fmt.Println("three");
	 
	fmt.Println("Hello");
	fmt.Println("World");

	myDefer();
}

func myDefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Println(i);
	}
}