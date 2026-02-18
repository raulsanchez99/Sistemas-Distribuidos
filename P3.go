package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	numNodos = 5
	contador = 0
)

var nodos [numNodos]Nodo

type Nodo struct {
	id     int
	estado string
}

type Beat struct {
	idLider int
}

func nodo(i int, canalNodo chan<- Nodo, canalHeartbeat chan Beat) {

	nodo := Nodo{id: i, estado: "follower"}
	//nodos[i] = nodo
	fmt.Printf("[Node %d] iniciado como %s.\n", nodo.id, nodo.estado)
	//canalNodo <- nodos[i]
	canalNodo <- nodo

	timer := time.NewTimer(600 * time.Millisecond)
	for {
		select {
		case <-timer.C:
			for i := 0; i > numNodos; i++ {
				fmt.Printf("[Node %d] activa timeout de votación.\n", nodo.id)
				canalNodo <- nodo
			}
		case heartbeat := <-canalHeartbeat:
			fmt.Printf("[Node %d] heartbeat recibido de %d.\n", nodo.id, heartbeat.idLider)
		}
	}

}

func votacion(canalNodo <-chan Nodo) {

	for {
		// Recibir un nodo del canal
		nodoRecibido := <-canalNodo

		// Imprimir el mensaje con el ID y el estado del nodo
		//fmt.Printf("Nodo %d es %s\n", nodoRecibido.id, nodoRecibido.estado)

		numAleatorio := rand.Intn(numNodos)
		if numAleatorio == nodoRecibido.id {
			nodoRecibido.estado = "candidate"
			fmt.Printf("[Votación] Nodo %d cambió a estado %s.\n", nodoRecibido.id, nodoRecibido.estado)
		}

	}
}

func lider(i int, canalHeartbeat chan Beat) {
	heartbeat := Beat{idLider: i}
	canalHeartbeat <- heartbeat
}

func main() {

	canalNodo := make(chan Nodo)
	canalHeartbeat := make(chan Beat)

	for i := 0; i < numNodos; i++ {
		go nodo(i, canalNodo, canalHeartbeat)
		go votacion(canalNodo)
		go lider(i, canalHeartbeat)
	}

	//	go lider(comunicacionNodos, respuestas)

	time.Sleep(1 * time.Second)

}
