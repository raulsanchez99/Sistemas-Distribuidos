package main

import (
	"fmt"
	"sync"
	"time"
)

const numFilosofos = 5

var palillos [numFilosofos]sync.Mutex

func filosofo(id int) {
	for {
		palilloIzq := id
		palilloDch := (id + 1) % numFilosofos

		fmt.Printf("El filósofo %d espera a poder comer\n", id)

		// Cojo los dos palillos que necesito
		palillos[palilloIzq].Lock()
		palillos[palilloDch].Lock()

		fmt.Printf("El filósofo %d está comiendo\n", id)
		time.Sleep(time.Second)

		// Suelto los palillos
		palillos[palilloDch].Unlock()
		palillos[palilloIzq].Unlock()

		fmt.Printf("El filósofo %d está pensando\n", id)
		time.Sleep(2 * time.Second)
	}

}

func main() {
	for i := 0; i < numFilosofos; i++ {
		go filosofo(i)
	}

	time.Sleep(5 * time.Second)
}
