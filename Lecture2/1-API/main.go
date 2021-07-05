package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Please add params!")
	fmt.Println("This is homepage")
}

func expression(w http.ResponseWriter, r *http.Request) {
	exp := r.URL.Query().Get("exp")

	params1 := r.URL.Query().Get("num1")
	fmt.Println(params1)
	params2 := r.URL.Query().Get("num2")
	num1, _ := strconv.Atoi(params1)
	fmt.Println(num1)
	num2, _ := strconv.Atoi(params2)

	switch exp {
	case "add":
		result := num1 + num2
		fmt.Fprintf(w, "result = %d", result)

	case "sub":
		result := num1 - num2
		fmt.Fprintf(w, "result = %d", result)

	case "mul":
		result := num1 * num2
		fmt.Fprintf(w, "result = %d", result)

	case "div":
		result := num1 / num2
		fmt.Fprintf(w, "result = %d", result)

	}

}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/exp", expression)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

//http://localhost:2800/exp?exp=add&&num1=10&&num2=20


func main() {
	handleRequests()
}
