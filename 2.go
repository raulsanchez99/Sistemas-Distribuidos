package main

import (
	"fmt"
	"sync"
	"time"
)

const numFilosofos = 5

var palillos [numFilosofos]sync.Mutex

func filosofos(id int) {
	for {
		palilloIzq := id
		palilloDch := (id + 1) % numFilosofos
		fmt.Printf("El filosofo %d espera para comer \n", id)

		palillos[palilloIzq].Lock()
		palillos[palilloDch].Lock()

		fmt.Printf("El filosofo %d esta comiendo \n", id)
		time.Sleep(time.Second)

		palillos[palilloDch].Unlock()
		palillos[palilloIzq].Unlock()

		fmt.Printf("El filosofo %d esta pensando \n", id)
		time.Sleep(2 * time.Millisecond)
	}
}

func main() {
	for i := 0; i < numFilosofos; i++ {
		go filosofos(i)
	}

	time.Sleep(5 * time.Second)
}
