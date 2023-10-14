package main

import "fmt"

func main() {

	slc := []int{2, 9, 1, 5, 6}

	great := slc[0]

	for _, v := range slc {
		if v > great {
			great = v
		}
	}

	fmt.Println(great)

}
