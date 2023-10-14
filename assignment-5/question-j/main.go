package main

import "fmt"

func main() {

	slc1, slc2 := []int{1, 2, 3}, []int{4, 5, 6}

	slc1 = append(slc1, slc2[0], slc2[1], slc2[2])

	fmt.Println(slc1)

}
