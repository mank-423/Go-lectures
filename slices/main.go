package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video of slices")

	// slice
	var fruitList = []string{"Apple", "Tomato", "Kiwi"}

	// appending the list
	fruitList = append(fruitList, "banana", "mango", "peach")

	fmt.Println("Type of fruitlist", fruitList)

	// slicing
	fruitList = append(fruitList[1:])

	fmt.Println("Type of fruitlist", fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 24
	highScores[2] = 34
	highScores[3] = 4

	// This gives error
	// highScores[4] = 10;

	// But this doesnt gives error
	highScores = append(highScores, 555, 666, 777)

	sort.Ints(highScores)

	fmt.Println("scores", highScores)
	fmt.Println("scores", sort.IntsAreSorted(highScores))

	// how to remove a value from slices based on index

	var courses = []string{"Nodejs", "Reactjs", "Javascript", "Python"}

	// Based on the first char
	// sort.Strings(courses)

	// Removing element
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println("Courses:", courses)

}
