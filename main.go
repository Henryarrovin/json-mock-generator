package main

import (
	"fmt"
	"json-mock-generator/client"
)

func main() {
	fmt.Println("OLLAMA URL:", client.BaseURL)
	fmt.Println("Model:", client.Model)

	response, err := client.CallOllama("Say hello in one sentence")
	if err != nil {
		fmt.Println("Error Connecting:", err)
		return
	}

	fmt.Println("Response:")
	fmt.Println(response)
}
