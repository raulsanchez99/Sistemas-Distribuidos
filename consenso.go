package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	numNodos = 5
)

var nodos [numNodos]Nodo

type Nodo struct {
	id     int
	estado string
}

type Voto struct {
	id      int
	numVoto int
}

func nodo(i int, canalNodo chan<- Nodo, canalHeartbeat chan Nodo, canalVoto chan Voto) {
	nodo := Nodo{id: i, estado: "follower"}
	nodos[i] = nodo
	fmt.Printf("[Node %d] iniciado como %s.\n", nodo.id, nodo.estado)

	// Create a new timer for each iteration
	timer := time.NewTimer(600 * time.Millisecond)

	select {
	case <-timer.C:
		fmt.Printf("[Node %d] activa timeout de votación.\n", nodo.id)

		// Iniciar temporizador de 100-200ms
		tiempo := time.Duration(100+rand.Intn(101)) * time.Millisecond
		tiempoEspera := time.NewTimer(tiempo)

		select {
		case <-tiempoEspera.C:
			canalNodo <- nodo

		case votoRecibido := <-canalVoto:
			// Se recibió un voto, parar el temporizador
			tiempoEspera.Stop()
			fmt.Printf("[Nodo %d] recibio por el nodo %d. Temporizador detenido.\n", nodo.id, votoRecibido.id)
		}

	case lider := <-canalHeartbeat:
		// Recibo heartbeat del lider
		fmt.Printf("[Node %d] heartbeat recibido de %d.\n", nodo.id, lider.id)
	}

}

func candidato(canalNodo chan Nodo, canalVoto chan Voto) {
	nodo := <-canalNodo

	// Cambiar el estado del nodo a "candidato"
	nodo.estado = "candidato"
	fmt.Printf("[Nodo %d] cambió su estado a %s.\n", nodo.id, nodo.estado)

	// Crear un voto y enviarlo por el canalVoto global
	voto := Voto{id: nodo.id, numVoto: 1}

	canalVoto <- voto

	fmt.Printf("[Nodo %d] comenzo una votacion (num 1).\n", nodo.id)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	canalNodo := make(chan Nodo)
	canalHeartbeat := make(chan Nodo)
	canalVoto := make(chan Voto)

	for i := 0; i < numNodos; i++ {
		go nodo(i, canalNodo, canalHeartbeat, canalVoto)
		go candidato(canalNodo, canalVoto)
	}

	time.Sleep(1 * time.Second)
}
