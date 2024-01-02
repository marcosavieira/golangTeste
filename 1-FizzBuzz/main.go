package main

import "fmt"

func main() {
	number := 100

	for i := 1; i <= number; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		} else {
			fmt.Println(i)
		}
	}
}
