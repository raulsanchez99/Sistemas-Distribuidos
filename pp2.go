package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const duracionGala = 4 * time.Second

var numOperadores = 5
var objetivoGala = 20000
var total = 0
var recaudacionTotal = 0

var canalCentralita = make(chan Llamada)
var canalOperador = make(chan Aviso)
var fin = make(chan bool)
var canalPresentadora = make(chan Aviso)

type Llamada struct {
	valor    int
	duracion time.Duration
}

type Aviso struct {
	numOperador   int
	valorDonacion int
	recaudacion   int
}

func centralita(llamadas chan Llamada, fin chan bool) {
	timer := time.NewTimer(duracionGala)

	for {
		select {
		case <-timer.C:
			// Cuando el temporizador alcanza la duración de la gala, envía fin = true al canal
			fin <- true
			return

		default:
			//Generacion de llamadas
			nuevaLlamada := Llamada{valor: 1 + rand.Intn(1001), duracion: time.Duration(30+rand.Intn(51)) * time.Millisecond}
			llamadas <- nuevaLlamada
			time.Sleep(time.Duration(25+rand.Intn(26)) * time.Millisecond)
		}
	}
	close(llamadas)
}

func operador(num int, llamadas <-chan Llamada, canalOperador chan<- Aviso) {
	sumaMutex := sync.Mutex{}

	for nuevaLlamada := range llamadas {
		time.Sleep(nuevaLlamada.duracion)

		sumaMutex.Lock()
		total += nuevaLlamada.valor
		recaudacionTotal = total

		//Genero aviso con nuevos datos
		nuevoAviso := Aviso{numOperador: num, valorDonacion: nuevaLlamada.valor, recaudacion: total}
		canalOperador <- nuevoAviso
		sumaMutex.Unlock()
	}
	close(canalOperador)
}

func pe(canalOperador <-chan Aviso, canalPresentadora chan<- Aviso) {
	for nuevoAviso := range canalOperador {
		if nuevoAviso.valorDonacion >= 750 {
			//Manda aviso a la presentadora si valorDonacion > 750
			canalPresentadora <- nuevoAviso
		}
	}
	close(canalPresentadora)
}

func presentadora(canalPresentadora <-chan Aviso, fin chan bool, objetivo int) {
	mutex := sync.Mutex{}

	fmt.Println("¡Bienvenidos a la gala benéfica contra el cambio climático!")
	fmt.Println("La cantidad total recaudada irá destinada a la replantación de árboles en el Amazonas.")
	fmt.Println("En este momento... abrimos las líneas!")

	timer := time.NewTimer(1 * time.Second)
	ticker := time.NewTicker(200 * time.Millisecond)

	for {
		select {
		case <-fin:
			fmt.Println("Y... cerramos líneas! A partir de este momento ya no recibiremos más llamadas")
			fmt.Printf("Qué gran noche! Hemos conseguido alcanzar la impresionante cifra de %d€ donados\n", recaudacionTotal)
			fmt.Println("Muchas gracias a todos")
			return

		case nuevoAviso := <-canalPresentadora:
			fmt.Printf("Nuestro operador %d acaba de recibir una generosa donación de %d€!\n", nuevoAviso.numOperador, nuevoAviso.valorDonacion)
			timer.Reset(1 * time.Second)

		case <-ticker.C:
			mutex.Lock()
			if objetivo > recaudacionTotal {
				//queda := objetivo - recaudacionTotal
				//fmt.Printf("El objetivo es %d, la recaudacion total es %d, queda %d \n", objetivo, recaudacionTotal, objetivo-recaudacionTotal)
				fmt.Printf("Estamos a tan solo %d € de llegar al objetivo de la noche\n", objetivo-recaudacionTotal)

			} else {
				//sobra := recaudacionTotal - objetivo
				//fmt.Printf("El objetivo es %d, la recaudacion total es %d, queda %d \n", objetivo, recaudacionTotal, objetivo-recaudacionTotal)
				fmt.Printf("Llevamos recaudados %d € más que el objetivo inicial de la noche!\n", recaudacionTotal-objetivo)
			}
			mutex.Unlock()

		case <-timer.C:
			fmt.Println("Continúen llamando para hacer sus donaciones por esta buena causa")
			timer.Reset(1 * time.Second)
		}
	}

}

func main() {

	go presentadora(canalPresentadora, fin, objetivoGala)

	for i := 0; i < numOperadores; i++ {
		go centralita(canalCentralita, fin)
		go operador(i, canalCentralita, canalOperador)
		go pe(canalOperador, canalPresentadora)
	}

	time.Sleep(5 * time.Second)
}
