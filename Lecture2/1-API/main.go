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
		fmt.Fprintf(w, "Add exp result = %d", result)

	case "sub":
		result := num1 - num2
		fmt.Fprintf(w, "Sub exp result = %d", result)

	case "mul":
		result := num1 * num2
		fmt.Fprintf(w, "Mul exp result = %d", result)

	case "div":
		if num2 != 0 {

			result := num1 / num2
			fmt.Fprintf(w, "Div exp result = %d", result)
		}
		fmt.Fprintf(w, "Div exp cannot run when num2 = 0")
	}

}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/exp", expression)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

//http://localhost:3000/exp?exp=add&num1=10&num2=20     : test link

func main() {
	handleRequests()
}
