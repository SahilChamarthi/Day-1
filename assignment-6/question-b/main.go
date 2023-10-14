package main

import "fmt"

func main() {

	fruit := map[string]int{"mango": 5, "apple": 3, "jack": 1, "graphs": 100, "watermelon": 1}

	for k, v := range fruit {
		fmt.Println("fruit:", k, " quantity:", v)

	}

}
