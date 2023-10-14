package main

import "fmt"

func main() {

	integer := []int{5, 2, 7, 8, 3}

	var sum, avg int

	for _, v := range integer {
		sum += v
	}

	avg = sum / len(integer)

	fmt.Println("sum is", sum)
	fmt.Println("avg is", avg)

}
