package main

import "fmt"

func binarySearch(items []int, find int) (bool, int) {
	low := 0
	high := len(items) - 1
	for low <= high {
		mid := (low + high) / 2
		if items[mid] < find {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if low == len(items) || items[low] != find {
		return false, -1
		//return false
	}
	return true, low
	//return true
}
func main() {
	//in binary search array must be sorted
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	find := 20
	x, i := binarySearch(items, find)
	if x {
		fmt.Println("Found at Index :", i)
	} else {
		fmt.Println("Not Found")
	}
}
