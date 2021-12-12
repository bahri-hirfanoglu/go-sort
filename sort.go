package main

/*
   @arr - Integer Array
   @Sort Asc
*/
func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		m := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[m] {
				m = j
			}
		}
		temp := arr[i]
		arr[i] = arr[m]
		arr[m] = temp
	}
	return arr
}

/*
   @arr - Integer Array
   @Sort Asc
*/
func bubbleSort(arr []int) []int {
	swap := false
	for swap {
		swap = false
		for i := 2; i < len(arr); i++ {
			if arr[i-1] < arr[i] {
				temp := arr[i]
				arr[i] = arr[i-1]
				arr[i-1] = temp
				swap = true
			}
		}
	}
	return arr
}

/*
   @arr - Integer Array
   @Sort Asc
*/
func insertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

/*
   @arr - Integer Array
   @Sort Asc
*/
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	f := mergeSort(arr[:len(arr)/2])
	s := mergeSort(arr[len(arr)/2:])
	return merge(f, s)
}

/*
   @a - First
   @b - Second
*/
func merge(a []int, b []int) []int {
	arr := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			arr = append(arr, a[i])
			i++
		} else {
			arr = append(arr, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		arr = append(arr, a[i])
	}
	for ; j < len(b); j++ {
		arr = append(arr, b[j])
	}
	return arr
}
