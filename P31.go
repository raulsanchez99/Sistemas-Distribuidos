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
			canalNodo <- nodos[i]

			/*Paso los nodos a la votacion
			canalNodo <- nodos[i]*/

		case lider := <-canalHeartbeat:
			//Recibo heartbeat del lider
			fmt.Printf("[Node %d] heartbeat recibido de %d.\n", nodo.id, lider.id)
		}
	}
}

func candidatos(canalNodo <-chan Nodo, canalVoto chan<- Voto, canalVotacion <-chan Voto) {
	nodosRecibidos := make([]Nodo, numNodos)

	for {
		// Recibir un slice de nodos del canal
		for i := 0; i < numNodos; i++ {
			nodosRecibidos[i] = <-canalNodo
		}

		// Iniciar temporizador de 100-200ms
		tiempo := time.Duration(100+rand.Intn(101)) * time.Millisecond
		tiempoEspera := time.NewTimer(tiempo)

		select {
		case <-tiempoEspera.C:
			// Cambiar el estado del nodo a "candidato"

			randnum := rand.Intn(numNodos)
			nodosRecibidos[randnum].estado = "candidato"
			fmt.Printf("[Nodo %d] cambió su estado a %s.\n", nodosRecibidos[randnum].id, nodosRecibidos[randnum].estado)

			// Crear un voto y enviarlo por el canalVoto
			votame := Voto{id: nodosRecibidos[randnum].id, numVoto: 1}
			canalVoto <- votame
			fmt.Printf("[Votación] Voto enviado por el nodo %d.\n", nodosRecibidos[randnum].id)

			//Comienza votacion
		case votoRecibido := <-canalVotacion:
			// Se recibió un voto, parar el temporizador
			//tiempoEspera.Stop()
			fmt.Printf("[Votación] Voto recibido por el nodo %d. Temporizador detenido.\n", votoRecibido.id)

		}
	}
}

func votacion(canalVoto chan Voto, canalVotacion chan Voto, canalLider chan Nodo) {

	<-canalVoto
	fmt.Println("Voto recibido")

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
	canalVotacion := make(chan Voto)
	canalLider := make(chan Nodo)

	for i := 0; i < numNodos; i++ {
		go nodo(i, canalNodo, canalHeartbeat, canalVoto)
	}
	go candidatos(canalNodo, canalVoto, canalVotacion)
	go votacion(canalVoto, canalVotacion, canalLider)
	go lider(canalLider, canalHeartbeat)

	time.Sleep(5 * time.Second)
}
