package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to pizza app")
	fmt.Println("Please reate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin);

	input, _ := reader.ReadString('\n');

	fmt.Println("Thansk for rating:", input);


	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64);

	if err != nil {
		fmt.Print(err);
	}else{
		fmt.Print("Rating change:", numRating + 1); 
	}

}
