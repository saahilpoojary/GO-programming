package main

import "fmt"

func isPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-(i+1)] {
			return false
		}
	}
	return true
}

func main() {
	var m1, m2, product int
	largest := 0

	for i := 100; i <= 999; i++ {
		for j := 101; j <= 999; j++ {
			product = i * j
			strm := fmt.Sprintf("%d", product)
			if isPalindrome(strm) && product > largest {
				largest = product
				m1 = i
				m2 = j
			}
		}
	}

	fmt.Printf("Largest palindrome product is %d\n", largest)
	fmt.Printf("The multiplicands are %d and %d\n", m1, m2)
}
