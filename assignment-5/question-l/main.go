package main

import "fmt"

func main() {

	slc1 := []int{4, 5, 2, 7, 8, 1}

	var sumOfEven int

	for _, v := range slc1 {

		if v%2 == 0 {
			sumOfEven += v
		}
	}

	fmt.Println(sumOfEven)

}
