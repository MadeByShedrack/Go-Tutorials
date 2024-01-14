package basics

import "fmt"

func Conditionals() {

	x := 7

	if x > 6 {
		fmt.Println("More than 6")
	} else if x < 2 {
		fmt.Println("Less than 2")
	} else {
		fmt.Println("No number specify")
	}
}
