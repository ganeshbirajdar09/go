package main

import "fmt"

func main() {
	fmt.Println("Welcome to functionns in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proResult, myMsg := proAdder(2, 5, 8, 7, 55)
	fmt.Println("ProResult is: ", proResult)
	fmt.Println("ProMessage is: ", myMsg)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hi Pro result function"
}

func greeter() {
	fmt.Println("Namastey from golang")
}
