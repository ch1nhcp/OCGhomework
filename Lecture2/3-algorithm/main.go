package main

import "fmt"

// Bubble Sort
func bubbleSort(arr []int) {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("\nAfter Bubble Sorting")
	for _, val := range arr {
		fmt.Println(val)
	}
}

//End BubbleSort

// Insertion Sort
func insertionSort(arr []int) {
	len := len(arr)
	for i := 1; i < len; i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}

	fmt.Println("After Insertion Sorting")
	for _, val := range arr {
		fmt.Println(val)
	}
}

// End Insertion Sort

// Merge Sort
func Merge(l, r []int) []int {
	ret := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(ret, r...)
		}
		if len(r) == 0 {
			return append(ret, l...)
		}
		if l[0] <= r[0] {
			ret = append(ret, l[0])
			l = l[1:]
		} else {
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
	return ret
}

func MergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	n := len(s) / 2
	l := MergeSort(s[:n])
	r := MergeSort(s[n:])
	return Merge(l, r)
}

//End MergeSort

//Quick Sort
func partition(a []int, lo, hi int) int {
	p := a[hi]
	for j := lo; j < hi; j++ {
		if a[j] < p {
			a[j], a[lo] = a[lo], a[j]
			lo++
		}
	}

	a[lo], a[hi] = a[hi], a[lo]
	return lo
}

func quickSort(a []int, lo, hi int) {
	if lo > hi {
		return
	}

	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

//End Quick Sort


//Quick Sort rewrite

//End Quick Sort







func main() {

	sample1 := []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
	bubbleSort(sample1)

	sample2 := []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
	insertionSort(sample2)

	sample3 := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
	fmt.Println(MergeSort(sample3))

	sample4 := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
	quickSort(sample4, 0, len(sample4)-1)
	fmt.Println(sample4)
}
