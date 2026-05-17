package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome time study")

	presentTime := time.Now()

	fmt.Println(presentTime);

	// get date with a format
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.February, 10, 23, 23, 0, 0, time.UTC);
	fmt.Println("Created:", createdDate);
}
