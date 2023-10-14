package main

import "fmt"

func main() {

	studentData := map[string]map[any]any{"sahil": {"Age": 22, "Grade": "A", "City": "kadapa"},
		"king":   {"Age": 21, "Grade": "A", "City": "tirupathi"},
		"vishal": {"Age": 24, "Grade": "A+", "City": "bihar"},
	}

	fmt.Println(studentData)

}
