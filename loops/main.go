package main

import "fmt"

func main() {
	fmt.Println("Loops file")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// fmt.Println(days);

	// for d := 0 ; d < len(days); d++{
	// 	fmt.Println(days[d]);
	// }

	// for i := range days {
	// 	fmt.Println(days[i]);
	// }


	for _, day := range days{
		fmt.Printf("value %v \n", day);
	}


	rogueVal := 1;

	for rogueVal < 10 {

		if rogueVal == 5{
			rogueVal++;
			continue;
		}

		fmt.Println("Val:", rogueVal);
		rogueVal++;
	}
}
