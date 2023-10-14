package main

import "fmt"

func main() {

	stud := map[string]string{"sahil": "A", "vikalp": "A+", "king": "B+"}

	for k, v := range stud {

		fmt.Println("Name:", k, " Grade:", v)

	}

	delete(stud, "king")

	fmt.Println(stud)

}
