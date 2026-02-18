package main

import "fmt"

func r(n int, c chan string) {
	for i := 0; i < n; i++ {
		c <- "ping"
	}
	close(c)
}

func main() {

	c := make(chan string)

	go r(5, c)

	for i := range c {
		fmt.Println(i)
	}

}
