package main

import (
	"fmt"
	"math/rand"
)

func main() {

	randNum := rand.Intn(10)
	var num int

	for {
		
		fmt.Print("Enter a number to guess: ")
		_, err := fmt.Scanln( &num)

		if err != nil {
			fmt.Println("not a number")
			return
		}

		if num > randNum {
			fmt.Println("too high")
		} else if num < randNum {
			fmt.Println("to low")
		} else {
			fmt.Println("your guess the right number")
			return
		}

	}

}
