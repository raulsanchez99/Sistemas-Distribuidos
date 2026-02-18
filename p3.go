package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	NUM_NODOS        = 5
	HEARTBEAT_DELAY  = 600 * time.Millisecond
	VOTO_TIMEOUT_MIN = 100 * time.Millisecond
	VOTO_TIMEOUT_MAX = 200 * time.Millisecond
)

type EstadoNodo int

const (
	Follower EstadoNodo = iota
	Leader
	Candidate
)

type Mensaje struct {
	Tipo   string
	Origen int // Identificador del nodo que envió el mensaje
}

func nodo(id int, canalHeartbeat chan Mensaje, estadoNodo chan EstadoNodo, nuevoLider chan int, votoRequest chan Mensaje) {
	estado := Follower
	heartbeatTimer := time.NewTimer(HEARTBEAT_DELAY)

	for {
		select {
		case <-heartbeatTimer.C:
			// No se recibió el heartbeat a tiempo, buscar un nuevo líder
			fmt.Printf("Nodo %d: No se recibió el heartbeat a tiempo. Buscando nuevo líder...\n", id)
			estadoNodo <- Follower                // Cambiar a estado Follower
			nuevoLider <- id                      // Informar que este nodo busca ser líder
			heartbeatTimer.Reset(HEARTBEAT_DELAY) // Reiniciar temporizador para el próximo heartbeat
		case mensaje := <-canalHeartbeat:
			// Se recibió un heartbeat del líder
			fmt.Printf("Nodo %d: Recibido mensaje del líder: %s\n", id, mensaje.Tipo)
			heartbeatTimer.Reset(HEARTBEAT_DELAY) // Reiniciar temporizador para el próximo heartbeat
		case <-time.After(time.Duration(rand.Intn(int(VOTO_TIMEOUT_MAX-VOTO_TIMEOUT_MIN)+1) + int(VOTO_TIMEOUT_MIN))):
			// Iniciar cuenta atrás aleatoria
			fmt.Printf("Nodo %d: No se recibió solicitud de voto a tiempo. Iniciando votación...\n", id)
			estadoNodo <- Candidate                                  // Cambiar a estado Candidate
			votoRequest <- Mensaje{Tipo: "VOTO_REQUEST", Origen: id} // Iniciar solicitud de voto
		case voto := <-votoRequest:
			// Se recibió una solicitud de voto
			fmt.Printf("Nodo %d: Solicitud de voto recibida de Nodo %d. Respondiendo favorablemente.\n", id, voto.Origen)
			votoRequest <- Mensaje{Tipo: "VOTO_RESPUESTA", Origen: id} // Responder favorablemente
		}
	}
}

func lider(canalHeartbeat []chan Mensaje, estadoNodo chan EstadoNodo, votoRequest chan Mensaje) {
	for {
		// Esperar un tiempo antes de enviar el primer heartbeat
		time.Sleep(500 * time.Millisecond)

		// Enviar el primer heartbeat a todos los nodos
		mensaje := Mensaje{Tipo: "HEARTBEAT"}
		for _, canal := range canalHeartbeat {
			canal <- mensaje
		}

		// Esperar un tiempo antes de enviar el próximo
		time.Sleep(200 * time.Millisecond)

		// Iniciar solicitud de voto a un nodo aleatorio
		nodoSolicitado := rand.Intn(NUM_NODOS)
		fmt.Printf("Enviando solicitud de voto a Nodo %d\n", nodoSolicitado)
		votoRequest <- Mensaje{Tipo: "VOTO_REQUEST", Origen: nodoSolicitado}

		// Esperar un tiempo antes de enviar el próximo heartbeat
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	canalHeartbeat := make([]chan Mensaje, NUM_NODOS)
	estadoNodo := make(chan EstadoNodo, NUM_NODOS)
	nuevoLider := make(chan int)
	votoRequest := make(chan Mensaje)

	for i := 0; i < NUM_NODOS; i++ {
		canalHeartbeat[i] = make(chan Mensaje)
		go nodo(i, canalHeartbeat[i], estadoNodo, nuevoLider, votoRequest)
	}

	go lider(canalHeartbeat, estadoNodo, votoRequest)

	// Simular cambio de líder cada cierto tiempo
	go func() {
		for {
			time.Sleep(3000 * time.Millisecond) // Cambio de líder cada 3 segundos
			nuevoLider <- -1                    // Indicar que se debe elegir un nuevo líder al azar
		}
	}()

	// Manejar cambios en el estado de los nodos
	for {
		select {
		case nuevoEstado := <-estadoNodo:
			fmt.Printf("Estado de los nodos actualizado: %v\n", nuevoEstado)
		case idNuevoLider := <-nuevoLider:
			if idNuevoLider == -1 {
				// Elegir un nuevo líder al azar
				nuevoLiderID := rand.Intn(NUM_NODOS)
				fmt.Printf("Seleccionando nuevo líder al azar: Nodo %d\n", nuevoLiderID)
				estadoNodo <- Leader // Cambiar a estado Leader
				// Detener temporizador de heartbeat para el nuevo líder
				canalHeartbeat[nuevoLiderID] <- Mensaje{Tipo: "STOP_TIMER"}
			} else {
				// Nodo específico busca ser líder
				fmt.Printf("Nodo %d busca ser líder\n", idNuevoLider)
			}
		}
	}
}
