package basics

import (
	"errors"
	"fmt"
)

func Pointers() {
	num1 := 8
	num2 := 0

	result, err := matchNumbers(num1, num2)

	if err == nil {
		fmt.Printf("Result: %v\n", result)
	} else {
		fmt.Println(err.Error())
	}
}

func matchNumbers(num1, num2 int) (int, error) {
	if num2 == 0 {
		return -1, errors.New("can not divide by zero")
	}

	return num1 / num2, nil
}
