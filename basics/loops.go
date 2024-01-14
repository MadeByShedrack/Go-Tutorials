package basics

import "fmt"

func Loops() {

	cars := []string{
		"Corvette",
		"Toyota",
		"Bentley",
		"Porsche",
		"Mercedes",
	}

	for index := 0; index < len(cars); index++ {
		fmt.Printf("Cars -> %v\n", cars[index])
	}

	myIndex := 0
	mySum := 0

	for myIndex < 20 {
		mySum += myIndex
		myIndex += 1
	}

	fmt.Printf("Total Sum -> %v\n", mySum)

	for carIndex, car := range cars {
		fmt.Printf("%v is car number -> %v\n", car, carIndex)
	}

}
