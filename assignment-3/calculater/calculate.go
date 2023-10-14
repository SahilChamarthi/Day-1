package calculater

import "fmt"

func SquareOfNum(n int) int {

	return n * n
}

func SumOfTwo(x, y int) int {
	return x + y
}

func SubOfTwo(x, y int) int {
	return x - y
}

func ProdOfTwo() int {

	var n1, n2 int

	fmt.Print("Enter num1 and num2: ")
	_, err1 := fmt.Scanf("%d", &n1)
	_, err2 := fmt.Scanf("%d", &n2)

	if err1 != nil {
		fmt.Println("num1 not a number")
		return 0
	}
	if err2 != nil {
		fmt.Println("num2 not a number")
		return 0
	}

	fmt.Print("Ans is ")
	return n1 * n2
}

func DivOfNum(num, div int) int {

	return num / div
}
