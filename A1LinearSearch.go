package main

import "fmt"

func linearSearch(items []int, find int) (bool, int) {
	x := 0
	for _, v := range items {
		x++
		if v == find {
			return true, x
		}
	}
	return false, -1
}
func main() {
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	find := 45
	x, i := linearSearch(items, find)
	if x {
		fmt.Println("Found at Index :", i)
	} else {
		fmt.Println("Not Found")
	}
}
