package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Fibonacci(number int) {
	a := 0
	b := 1
	for a < number {
		fmt.Println(a)
		c := a + b
		a = b
		b = c
	}
}

func PrimeNumber(a []int) {
	error := 0
	for _, v := range a {
		if v < 2 {
			error += 1
		}
		for i := 2; i < v/2; i++ {
			if v%i == 0 {
				error += 1
			}
		}
	}
	if error != 0 {
		fmt.Println("It is Prime array!")

	} else {
		fmt.Println("It is not Prime array!")
	}
}

func readFile(file string) []int {
	content, err := ioutil.ReadFile("textfile.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
	arr := strings.Split(string(content), ",")
	var newarray []int
	for _, v := range arr {
		a, _ := strconv.Atoi(v)
		newarray = append(newarray, a)
	}
	return newarray
}

func main() {

}
