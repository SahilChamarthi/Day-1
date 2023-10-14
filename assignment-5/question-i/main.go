package main

import "fmt"

func main() {

	slc := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(slc); i++ {
		slc[i] = slc[i] * 2
	}
	fmt.Println(slc)

}
