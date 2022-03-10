package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Running server ...")
	http.HandleFunc("/", logEnvironmentVariables)
	http.ListenAndServe(":8090", nil)
}

func logEnvironmentVariables(w http.ResponseWriter, req *http.Request) {
	firstVar := os.Getenv("FIRST_VAR")
	secondVar := os.Getenv("SECOND_VAR")
	thirdVar := os.Getenv("THIRD_VAR")

	fmt.Printf("First variable: %v\n", firstVar)
	fmt.Printf("Second variable: %v\n", secondVar)
	fmt.Printf("Third variable: %v\n", thirdVar)
}
