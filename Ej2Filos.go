package main

import (
	"fmt"
	"math/rand"
	"time"
)

const numFilosofos = 5

var palillos [numFilosofos]chan bool

func filosofo(id int) {
	for {
		palilloIzq := id
		palilloDch := (id + 1) % numFilosofos

		fmt.Printf("El filósofo %d espera a poder comer\n", id)

		// Cojo los dos palillos que necesito
		<-palillos[palilloIzq]
		select {
		case <-palillos[palilloDch]:
			fmt.Printf("El filósofo %d está comiendo\n", id)
			time.Sleep(100 * time.Millisecond)

			// Suelto los palillos
			palillos[palilloIzq] <- true
			palillos[palilloDch] <- true

			fmt.Printf("El filósofo %d está pensando\n", id)
			time.Sleep(50 * time.Millisecond)

		case <-time.After(time.Duration(70+rand.Intn(10)) * time.Millisecond):
			palillos[palilloIzq] <- true
			fmt.Printf("El filósofo %d deja su palillo porque el segundo está cogido.\n", id)
		}

	}

}

func main() {

	for i := 0; i < numFilosofos; i++ {
		palillos[i] = make(chan bool, 1)
		palillos[i] <- true
	}

	for i := 0; i < numFilosofos; i++ {
		go filosofo(i)
	}

	time.Sleep(1 * time.Second)
}
