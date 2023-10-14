package main

import "fmt"

func main() {

	slc := []int{6, 1, 9, 2, 5}

	var num int

	fmt.Print("Enter element to search: ")
	_, err := fmt.Scanln(&num)

	if err != nil {
		fmt.Println("not a number")
		return
	}

	for i := 0; i < len(slc); i++ {

		if slc[i] == num {
			fmt.Println(num, "found")
			return
		}
	}
	fmt.Println(num, "not found")

}
