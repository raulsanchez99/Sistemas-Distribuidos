package main

import "fmt"

func main() {
	for n := 1; n <= 100; n++ {

		fmt.Print("Numero ", n)
		if n%3 == 0 {
			fmt.Println("Fizz ")
		}
		if n%5 == 0 {
			fmt.Println("Buzz ")
		}
		if n%3 == 0 && n%5 == 0 {
			fmt.Println("Fizzbuzz ")
		}
		fmt.Println(" ")
	}
}
