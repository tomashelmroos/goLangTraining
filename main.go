package main

import "fmt"

func main() {
fmt.Println(fibonacci())
}

func fibonacci() int {
	var totalSum int
	var nextNumber int
	var fibonacci []int
	var i int
	for nextNumber < 4000000 {
		if i < 2 {
			nextNumber = i+1 // needed to build the initial array
		} else {
			fibonacciLength := len(fibonacci) - 1
			nextNumber = fibonacci[fibonacciLength] + fibonacci[fibonacciLength-1]
		}
		fibonacci = append(fibonacci,nextNumber)
		if isEven(nextNumber) {
			totalSum += nextNumber
		}
		i++
	}
	return totalSum
}

func isEven(x int) bool {
	return x % 2 == 0
}