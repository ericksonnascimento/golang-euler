package main

import "fmt"

/*
Multiples of 3 and 5

Problem 1
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

//SumMultiplesOf3Or5 ...This function calculate the sum of the multiples of 3 or 5 below limit
func SumMultiplesOf3Or5(limit int) int {
	sum := 0

	for i := 0; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}

func main() {
	limit := 10
	sum := SumMultiplesOf3Or5(limit)
	fmt.Printf("The sum of the multiples of 3 or 5 below %v is %v\n", limit, sum)
}
