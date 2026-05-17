package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome"
	fmt.Println(welcome);

	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Etner rating:");

	// comma ok
	input, _ := reader.ReadString('\n');
	fmt.Println("Reader:", input);

}