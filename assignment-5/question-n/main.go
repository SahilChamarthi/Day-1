package main

import "fmt"

func main() {

	slc := []int{4, 2, 6, 3, 9, 1}
	var findNum int = 2

	fmt.Println(findIndex(slc, findNum))

}

func findIndex(a []int, n int) int {

	for i := 0; i < len(a); i++ {
		if a[i] == n {
			return i
		}
	}
	return -1

}
