package main

import (
	"fmt"
	"math/rand"
	"time"
)

func agente(ingrediente [3]string) {
	var ingrediente1 int
	var ingrediente2 int

	for {
		ingrediente1 = rand.Intn(3)
		ingrediente2 = rand.Intn(3)

		if ingrediente1 != ingrediente2 {
			break
		}
	}

	fmt.Printf("El agente pone en la mesa los ingredientes: %d y %d\n", ingrediente1, ingrediente2)
}

func fumador(num int, ingrediente [3]string) {

	fmt.Printf("El fumador %d [%s] coge los ingredientes y fuma\n", num, ingrediente)
}

func main() {
	ingredientes := [3]string{"papel", "tabaco", "fósforos"}

	numFumadores := 3

	for i := 0; i <= numFumadores; i++ {
		go agente(ingredientes)
		go fumador(i, ingredientes)
		time.Sleep(time.Duration(15) * time.Millisecond)

	}

}
