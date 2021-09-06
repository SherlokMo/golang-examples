package main

import (
	"fmt"
	"time"
)

func devide(arr []int, pivot int, closure func(a int, b int) bool) []int {

	var res []int

	for _, value := range arr {
		if closure(value, pivot) {
			res = append(res, value)
		}
	}

	return res
}

func qsort(arr []int) []int {

	var res []int

	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]

	less := devide(arr, pivot, func(value int, pivot int) bool {
		if value < pivot {
			return true
		}
		return false
	})

	greater := devide(arr, pivot, func(value int, pivot int) bool {
		if value > pivot {
			return true
		}
		return false
	})

	res = append(qsort(less), pivot)
	res = append(res, qsort(greater)...)

	return res
}

func main() {

	start := time.Now()

	fmt.Println(qsort([]int{9, 8, 6, 5, 4, 2}))

	fmt.Printf("qsort took %s", time.Since(start))
}
