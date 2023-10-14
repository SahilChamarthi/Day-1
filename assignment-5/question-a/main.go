package main

import "fmt"

func main() {

	ar := [5]int{3, 4, 1, 7, 8}

	for i := 0; i < len(ar); i++ {
		fmt.Print(ar[i], " ")
	}
}
