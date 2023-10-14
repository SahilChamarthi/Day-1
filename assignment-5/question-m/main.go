package main

import "fmt"

func main() {

	slc1, slc2 := []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}

	for i := 0; i < len(slc1); i++ {

		if slc1[i] != slc2[i] {
			fmt.Println("not equal")
			return
		}

	}

	fmt.Println("both are equal")

}
