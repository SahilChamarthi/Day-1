package main

import "fmt"

var (
	Name string
	Age  int
	Sal  float64
	All  []any
)

func main() {

	Name = "sahil"
	Age = 22
	Sal = 0
	All = []any{Name, Age, Sal}

	fmt.Println(Name)
	fmt.Println(Age)
	fmt.Println(Sal)
	fmt.Println(All)

}
