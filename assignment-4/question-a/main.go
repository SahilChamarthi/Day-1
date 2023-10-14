package main

import "fmt"

func main() {

	var num int = 9

	if num&2 == 0 {
		fmt.Println(num, "is even")
		return
	}
	fmt.Println(num, "is odd")

}
