package basics

import (
	"errors"
	"fmt"
)

func Structure() {
	e := employee{
		name:   "Mario Yoshi",
		salary: 1200,
		age:    26,
		tax:    10,
	}
	e.employeeDetails()

	taxDenom := 0

	taxResult, err := e.calculateTax(taxDenom, e.tax)

	if err == nil {
		fmt.Printf("Tax Result -> %v\n", taxResult)
	} else {
		fmt.Println(err.Error())
	}
}

type employee struct {
	name   string
	age    int
	salary int
	tax    int
}

func (e employee) calculateTax(myTax, taxDenom int) (int, error) {
	if taxDenom == 0 {
		return -1, errors.New("integer divde by zero error")
	}

	return myTax / taxDenom, nil
}

func (e employee) employeeDetails() {
	fmt.Printf("Name -> %v\n", e.name)
	fmt.Printf("Salary -> %v $\n", e.salary)
	fmt.Printf("Age -> %v\n", e.age)
}
