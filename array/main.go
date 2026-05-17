package main

import "fmt"

func main() {
	fmt.Println("Array file")

	var fruitList [4]string;

	fruitList[0] = "Banana";
	fruitList[1] = "Manana";
	// fruitList[2] = "Kiwi";
	fruitList[3] = "Peach";

	fmt.Println("Fruit list:", fruitList)
	fmt.Println("Fruit list:", len(fruitList))

	var vegList = [3]string{"potato", "beans", "gourd"}

	fmt.Println("Veggie list:", vegList);
	fmt.Println("Veggie list:", len(vegList));
}
