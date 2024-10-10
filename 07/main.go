package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func handleDivision(a, b float64) {
	result, err := divide(a, b)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Result: %.2f\n", result)
}

type NegativeNumberError struct {
	Number float64
}

func (e *NegativeNumberError) Error() string {
	return fmt.Sprintf("negative number encountered: %.2f", e.Number)
}

func checkForNegativeNumber(num float64) error {
	if num < 0 {
		return &NegativeNumberError{Number: num}
	}
	return nil
}

func main() {
	handleDivision(10, 2)
	handleDivision(10, 0)

	num := -5.0
	err := checkForNegativeNumber(num)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("The number is positive:", num)
	}

	num = 10.0
	err = checkForNegativeNumber(num)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("The number is positive:", num)
	}
}
