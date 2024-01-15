package tourofgo

import "fmt"

func MyFunctions() {
	result := addTwo(12, 14)
	fmt.Printf("Result -> %v\n", result)
}

func addTwo(x int, y int) int {
	return x + y
}
