package main

import (
	"fmt"
)

func largestPrimeFactor(n int) int {
	largest := -1

	for i := 2; i <= n; i++ {
		for n%i == 0 {
			largest = i
			n /= i
		}
	}

	return largest
}

func main() {
	num := 38 // Replace this with the number for which you want to find the largest prime factor
	result := largestPrimeFactor(num)
	fmt.Printf("The largest prime factor of %d is %d\n", num, result)
}
