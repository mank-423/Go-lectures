package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")
	
	// result := adder(3,5);

	result, _ := proAdder(2,3,4,4,5,6);

	fmt.Println("sum:", result);
}

func proAdder(values ...int) (int, string){
	total := 0;

	for _, value := range values{
		total += value;
	}

	return total, "This is the string";
}