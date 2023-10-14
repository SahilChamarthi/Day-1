package main

import "fmt"

func main() {

	number := []int{5}
	fmt.Println(number)

	number = append(number, 10)
	fmt.Println(number)

	number = append(number, 15)
	fmt.Println(number)

	number = append(number, 20)
	fmt.Println(number)

	number = append(number, 25)
	fmt.Println(number)

	var findit int

	for i := range number {
		if i == len(number)/2 {

			findit = i
		}
	}

	number = removeIndedx(number, findit)

	fmt.Println(number)

}

func removeIndedx(ar []int, index int) []int {

	return append(ar[:index], ar[index+1:]...)
}
