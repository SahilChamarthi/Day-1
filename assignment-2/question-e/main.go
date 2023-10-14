package main

import "fmt"

const (
	EmpName = "sahil"
	EmpAge  = 22
	EmpSal  = 0
)

func main() {

	fmt.Println(EmpName)
	fmt.Println(EmpAge)
	fmt.Println(EmpSal)
	// EmpName = 23 // const cannot be redeclared
	// fmt.Println(EmpName)

}
