package main

import "fmt"

func main() {

	var even, odd int

	slc := []int{8, 1, 3, 6, 2, 9, 7}

	for _, v := range slc {

		if v%2 == 0 {
			even++
		} else {
			odd++
		}

	}

	fmt.Println("count of even's are", even)
	fmt.Println("count of odd's are", odd)

}
