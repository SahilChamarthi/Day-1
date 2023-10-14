package main

import "fmt"

func main() {

	slc := []string{"mango", "banana", "apple", "graphs", "watermelon"}

	for _, v := range slc {
		fmt.Print(v, " ")
	}

}
