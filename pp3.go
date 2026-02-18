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

	timer := time.NewTimer(600 * time.Millisecond)

	for {
		select {
		case <-timer.C:
			fmt.Printf("[Node %d] activa timeout de votación.\n", nodo.id)

			tiempo := time.Duration(100+rand.Intn(101)) * time.Millisecond
			tiempoEspera := time.NewTimer(tiempo)

			select {
			case <-tiempoEspera.C:
				//El nodo se convierte en candidate
				nodo.estado = "candidato"
				fmt.Printf("[Nodo %d] se nombra %s\n", nodo.id, nodo.estado)
				votame := Voto{id: nodo.id, numVoto: 1}
				canalVoto <- votame

			//Comienza votacion
			case <-canalVoto:
				tiempoEspera.Stop()
				//Recibimos una solicitud, el nodo vota
				//Func votacion

			}

			/*Paso los nodos a la votacion
			canalNodo <- nodos[i]*/

		case lider := <-canalHeartbeat:
			//Recibo heartbeat del lider
			fmt.Printf("[Node %d] heartbeat recibido de %d.\n", nodo.id, lider.id)
		}
	}
}

func votacion(canalNodo <-chan Nodo) {
	for {
		// Recibir un slice de nodos del canal
		nodosRecibidos := make([]Nodo, numNodos)
		for i := 0; i < numNodos; i++ {
			nodosRecibidos[i] = <-canalNodo
		}

		// Elegir un número aleatorio entre 1 y 5
		numNodosElegir := rand.Intn(numNodos) + 1

		// Imprimir el número de nodos a elegir
		//fmt.Printf("[Votación] Elegir %d nodos de entre los %d nodos recibidos.\n", numNodosElegir, numNodos)

		// Lógica para elegir nodos aleatorios y cambiar el estado a "candidate" si coincide con el ID del nodo
		nodosElegidos := make(map[int]bool)
		for i := 0; i < numNodosElegir; {
			indiceAleatorio := rand.Intn(numNodos)
			if !nodosElegidos[indiceAleatorio] {
				nodosElegidos[indiceAleatorio] = true

				nodoElegido := &nodosRecibidos[indiceAleatorio]
				//fmt.Printf("[Votación] Nodo %d elegido.\n", nodoElegido.id)

				// Cambiar el estado a "candidate" si coincide con el ID del nodo
				if nodoElegido.id == nodoElegido.id {
					nodosRecibidos[indiceAleatorio].estado = "candidate"
					fmt.Printf("[Votación] Nodo %d cambió su estado a %s.\n", nodoElegido.id, nodoElegido.estado)
				}

				i++
			}
		}

		// Imprimir los datos de cada nodo
		/*for _, nodo := range nodosRecibidos {
			fmt.Printf("[Votación] Nodo %d con estado %s recibido.\n", nodo.id, nodo.estado)
		}*/
	}
}

func lider(canalLider <-chan Nodo, canalHeartbeat chan<- Nodo) {
	for {
		// Esperar a recibir un nodo en el canalLider
		lider := <-canalLider

		// Iniciar un ticker de 400 milisegundos
		ticker := time.NewTicker(400 * time.Millisecond)

		for {
			select {
			case <-ticker.C:
				// Enviar información del nodo lider por el canalHeartbeat
				canalHeartbeat <- lider
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	canalNodo := make(chan Nodo)
	canalHeartbeat := make(chan Nodo)
	canalVoto := make(chan Voto)
	canalLider := make(chan Nodo)

	for i := 0; i < numNodos; i++ {
		go nodo(i, canalNodo, canalHeartbeat, canalVoto)
	}
	go votacion(canalNodo)
	go lider(canalLider, canalHeartbeat)

	time.Sleep(5 * time.Second)
}
