package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=1214213"

func main() {
	fmt.Println("URL")
	fmt.Println(myUrl)

	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Host)

	// fmt.Println(strings.Split(result.RawQuery, "&"));
	qparams := result.Query()
	fmt.Printf("Type of query: %T\n", qparams)
	
	fmt.Println(qparams["coursename"]);

	for _, val := range qparams{
		fmt.Println("Params:", val); 
	}
	
	// var url URL

	// partsOfUrl := &url.URL{
	// 	Scheme: "https",
	// 	Host: "lco.dev",
	// 	Path: "/awfwa",
	// 	RawPath: "user=mank",
	// }
}
