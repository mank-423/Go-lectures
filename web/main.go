package main

import (
	"encoding/json"
	"fmt"

	"net/http"
)

const url = "https://swoovo-backend.onrender.com/health"

type responseType struct {
	Status    string `json: "status"`
	Service   string `json: "service"`
	Timestamp string `json: "timestamp"`
}

func main() {
	fmt.Println("Web request")

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	}

	fmt.Printf("Response type: %T \n", response)

	defer response.Body.Close() // caller's responsibility

	var content responseType

	err = json.NewDecoder(response.Body).Decode(&content)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Timestamp: %+v \n", content)

}
