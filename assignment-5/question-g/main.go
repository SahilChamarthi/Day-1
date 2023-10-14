package main

import "fmt"

func main() {

	slc := []int{1, 2, 3, 4, 5}
	fmt.Println("before swap")
	fmt.Println(slc)

	var i, j int = 0, len(slc) - 1

	for i < j {

		var temp int = slc[i]
		slc[i] = slc[j]
		slc[j] = temp
		i++
		j--

	}

	fmt.Println("after swap")
	fmt.Println(slc)

}
