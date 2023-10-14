package main

import "fmt"

func match(s []int, k int) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == k {
			return false
		}
	}
	return true
}

func main() {

	slc1 := []int{3, 4, 7, 1, 8, 4, 7, 1}
	slc2 := []int{}
	slc2 = append(slc2, slc1[0])

	for i := 1; i < len(slc1); i++ {
		if match(slc2, slc1[i]) {
			slc2 = append(slc2, slc1[i])
		}
	}
	fmt.Println("before:", slc1)
	fmt.Println("after", slc2)

}
