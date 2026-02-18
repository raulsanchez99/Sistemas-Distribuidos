package main

import (
	"fmt"
	"time"
)

func talk(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	var lista = []string{"Hola", "que", "tal", "estas"}

	// Llamar a la función talk de manera secuencial
	for _, str := range lista {
		go talk(str)
	}

	// Esperar un poco para permitir que las goroutines terminen
	time.Sleep(2 * time.Second)
}
