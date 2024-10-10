package main

import "fmt"

func sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	numbers := []int{5, 10, 15, 20, 25}

	result := sum(numbers)
	fmt.Println("numbers: ", numbers)
	fmt.Printf("The sum of the numbers is: %d\n", result)

	numbers = append(numbers, 30, 35, 40)
	fmt.Println("Updated slice:", numbers)

	updatedResult := sum(numbers)
	fmt.Printf("The sum of the updated numbers is: %d\n", updatedResult)
}
