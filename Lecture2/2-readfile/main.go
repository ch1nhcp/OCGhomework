package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	// Doc file text
	content, err := ioutil.ReadFile("test.txt")
	fmt.Println("Open and read file successfully!")

	if err != nil {
		log.Fatal(err)
	}
	// In noi dung file text
	fmt.Println(string(content))

	//Đọc file text, tìm số lớn nhất, nhỏ nhất, trung bình
	// tim so lon nhat
	arr := []int{28, 33, 16, 7, 5, 88}

	max := arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	fmt.Println("So lon nhat : ", max)

	// tim so nho nhat
	min := arr[0]
	for _, value := range arr {
		if value > max {
			min = value
		}
	}
	fmt.Println("So nho nhat : ", min)

	//gia tri trung binh
	var avg int = 0
	for _, value := range arr {
		avg += arr[value]
	}
	var avgValue  int = 5
	fmt.Println("Gia tri trung binh: ", avgValue)

}
