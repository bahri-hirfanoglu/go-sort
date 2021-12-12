package main

import "fmt"

func main() {
	arr := []int{6, 3, 16, 744, 5, 5, 4, 6, 4, 2, 3, 5, 66, 78, 34, 5, 67, 46, 6, 7, 345, 67, 5, 345, 67, 45, 345, 6, 7}
	fmt.Printf("Seleciton Sort: %+v\n", selectionSort(arr))
	fmt.Printf("Bubble Sort: %+v\n", bubbleSort(arr))
	fmt.Printf("Insertion Sort: %+v\n", insertionSort(arr))
	fmt.Printf("Merge Sort: %+v\n", mergeSort(arr))
}
