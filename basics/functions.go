package basics

import (
	"errors"
	"fmt"
	"math"
)

func Examples() {
	result := sum(12, 18)
	fmt.Printf("The result is -> %v\n", result)

	squareRoot, err := square(0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(squareRoot)
	}

	num1 := 12
	num2 := 0

	answer, err := divideNumber(num1, num2)

	if err == nil {
		fmt.Printf("Result -> %v\n", answer)
	} else {
		fmt.Println(err.Error())
	}
}

func sum(x, y int) int {
	return x + y
}

func square(root float64) (float64, error) {
	if root < 0 {
		return 0, errors.New("undefined for negative numbers")
	}

	return math.Sqrt(root), nil
}

func divideNumber(num, denom int) (int, error) {
	if denom == 0 {
		return -1, errors.New("divide by zero is not allowed")
	}
	return num / denom, nil
}
