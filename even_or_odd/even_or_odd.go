package main

import "fmt"

func even_or_odd(numbers []int) {
	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(number, "is even!")
		} else {
			fmt.Println(number, "is odd!")
		}
	}
}
