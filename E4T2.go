package main

import "fmt"

func suma(nums []int, ch1 chan int, ch2 chan bool) {
	resultado := 0
	for _, n := range nums {
		resultado += n
	}
	ch1 <- resultado

	fmt.Println("Suma parcial", resultado)

	ch2 <- true
	fmt.Println("Goroutime terminated")
}

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sumasParciales := make(chan int, 2)
	terminados := make(chan bool, 2)

	go suma(numeros[:5], sumasParciales, terminados)
	go suma(numeros[5:], sumasParciales, terminados)

	//Esperamos a que termine
	<-terminados
	<-terminados
	close(sumasParciales)

	// Recuperamos resultados
	total := 0
	for s := range sumasParciales {
		total += s
	}

	fmt.Println("Suma total: ", total)

}
