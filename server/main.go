package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to server")
	PerformGetRequest()
}

// to do: url in paramter of function
func PerformGetRequest() {
	const myUrl = "https://mt-task.onrender.com"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status:", response.StatusCode)
	fmt.Println("Content length is:", response.ContentLength)

	// content is in byte format
	var responseString strings.Builder

	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	// fmt.Println(string(content));
	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}

func performPostJsonRequest() {
	const url = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"coursename": "PJ GO",
			"price": 0,
			"platform": "learncode.com"
		}
	`)

	response, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func performPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	// form data
	data := url.Values{}
	data.Add("firstname", "mayank")
	data.Add("lastname", "kumar")

	response, err := http.PostForm(myUrl, data);

	defer response.Body.Close();
	
	if err != nil {
		panic(err)
	}

	content, _ := io.ReadAll(response.Body);

	fmt.Println(string(content))
}
