package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
		fmt.Println("It is not Prime array!")

	} else {
		fmt.Println("It is Prime array!")
	}
}

func readFile(file string) []int {
	content, err := ioutil.ReadFile(file)

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


func WriteFileByIO() {
	//Bước 1: tạo file
	file, err := os.Create("writed.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//Bước 2: ghi file
	message := []byte("Hi there!")
	err = ioutil.WriteFile("writed.txt", message, 0644)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}

	//Bước 3: đóng fileText
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}

func main() {

	var textFile []int
	textFile = readFile("textfile.txt")
	PrimeNumber(textFile)
}
