package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const NUMERO_PARTICIPANTES = 4

var wg4 sync.WaitGroup
var listaIntentos [4]int
var mux5 sync.Mutex
var mux6 sync.Mutex
var mux7 sync.Mutex

func presentador(listaCanales []chan int) {
	numAleatorio := rand.Intn(40) + 1
	fmt.Printf("-- Comienza el concurso.  NUM: %d\n", numAleatorio)
	//Notificamos a todos los participantes para que empiecen
	wg4.Done()
	for i := 0; i < NUMERO_PARTICIPANTES; i++ {
		go func(listaCanales []chan int, index int) {
			for {
				n := <-listaCanales[index]
				if n > numAleatorio {
					fmt.Printf("Conc. %d dice %d y se pasa!\n", index, n)
					listaCanales[index] <- 1
					mux5.Lock()
					listaIntentos[index]++
					mux5.Unlock()
				} else if n < numAleatorio {
					fmt.Printf("Conc. %d dice %d y se queda corto!\n", index, n)
					listaCanales[index] <- -1
					mux5.Lock()
					listaIntentos[index]++
					mux5.Unlock()
				} else {
					fmt.Printf("Conc. %d dice %d y acierta!\n", index, n)
					listaCanales[index] <- 0
				}
			}
		}(listaCanales, i)
	}
}

func participante(canalConcursante chan int) {
	//Esperamos la señal del presentador para empezar a adivinar el numero
	wg4.Wait()
	for {
		numAleatorio := rand.Intn(40) + 1
		canalConcursante <- numAleatorio
		n := <-canalConcursante
		if n == -1 || n == 1 {
			continue
		} else {
			return
		}
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())
	var listaCanes []chan int
	wg4.Add(1)
	for i := 0; i < NUMERO_PARTICIPANTES; i++ {
		canalConcursante := make(chan int)
		go participante(canalConcursante)
		listaCanes = append(listaCanes, canalConcursante) //nos guardamos los canales de los concursantes
	}
	go presentador(listaCanes)

	time.Sleep(5 * time.Second)

	fmt.Printf("--Puntuaciones finales--\n")
	for i := 0; i < NUMERO_PARTICIPANTES; i++ {
		fmt.Printf("concursante %d: %d ptos\n", i, 100/listaIntentos[i])

	}

}
