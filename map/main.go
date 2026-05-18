package main

import (
	"fmt"
)

func main() {
	lang := make(map[string]string)

	lang["Js"] = "JS"
	lang["Py"] = "Python"
	lang["Ts"] = "TS"

	fmt.Println("list of all:", lang)

	delete(lang, "Js")

	fmt.Println("list of all:", lang)

	// loop over map
	for _, value := range lang{
		fmt.Printf("value %v \n", value);
	}

}
