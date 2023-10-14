package main

import (
	"assignment-3/calculater"
	"assignment-3/temperature"
	"fmt"
)

func main() {

	//calculater

	fmt.Println(calculater.SquareOfNum(5))

	fmt.Println(calculater.SumOfTwo(10, 20))

	fmt.Println(calculater.SubOfTwo(20, 10))

	fmt.Println(calculater.ProdOfTwo())

	fmt.Println(calculater.DivOfNum(6, 3))

	//temperature

	fmt.Println(temperature.FarinToCelsius(78))

}
