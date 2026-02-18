package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Llamada struct {
	valor    int
	duracion time.Duration
}

type Aviso struct {
	numOperador   int
	valorDonacion int
	actual        int
	recaudacion   int
}

func centralita(llamadas chan Llamada, duracionTotal time.Duration) {
	tiempoInicio := time.Now()

	for time.Since(tiempoInicio) < duracionTotal {
		nuevaLlamada := Llamada{valor: 1 + rand.Intn(1001), duracion: time.Duration(30+rand.Intn(51)) * time.Millisecond}
		//fmt.Println("PII")
		llamadas <- nuevaLlamada
		time.Sleep(time.Duration(25+rand.Intn(26)) * time.Millisecond)
	}
}

func operador(num int, llamadas <-chan Llamada, canalOperador chan<- Aviso, objetivo int) {
	sumaMutex := sync.Mutex{}
	total := 0

	for nuevaLlamada := range llamadas {
		//fmt.Println("Llamada cogida")
		time.Sleep(nuevaLlamada.duracion)

		//Restar valor al total
		sumaMutex.Lock()
		objetivo -= nuevaLlamada.valor
		total += nuevaLlamada.valor
		sumaMutex.Unlock()

		nuevoAviso := Aviso{numOperador: num, valorDonacion: nuevaLlamada.valor, actual: objetivo, recaudacion: total}

		//fmt.Printf("Operador %d restó %d€ al objetivo de la gala. Nuevo objetivo: %d\n", num, nuevaLlamada.valor, *objetivo)
		//fmt.Printf("Qué gran noche! Hemos conseguido alcanzar la impresionante cifra de %d € donados", total)

		//Ahora tenemos que ver si el valor del objetivo se alcanzó

		canalOperador <- nuevoAviso
	}
	close(canalOperador)
}

func pe(canalOperador <-chan Aviso, canalPresentadora chan<- Aviso) {
	for nuevoAviso := range canalOperador {

		if nuevoAviso.valorDonacion >= 750 {
			//Manda aviso a la presentadora
			canalPresentadora <- nuevoAviso
		}
	}
}

func presentadora(canalPresentadora <-chan Aviso) {

	fmt.Println("Buenas noches y bienvenidos a la gala benéfica contra el cambio climático")
	fmt.Println("La cantidad total recaudada irá destinada a la replantación de árboles en el Amazonas")
	fmt.Println("En este momento... abrimos las líneas!")

	//En este momento inicia la centralita

	//Mandar cierre a centralita
	for nuevoAviso := range canalPresentadora {
		fmt.Printf("Nuestro operador %d acaba de recibir una generosa donación de %d€!\n", nuevoAviso.numOperador, nuevoAviso.valorDonacion)
		/*
			fmt.Println("Continúen llamando para hacer sus donaciones por esta buena causa")

			fmt.Println("Estamos a tan solo 15966 € de llegar al objetivo de la noche")
			fmt.Println("Llevamos recaudados 9380 € más que el objetivo inicial de la noche!")

			fmt.Println("Y... cerramos líneas! A partir de este momento ya no recibiremos más llamadas")
			fmt.Printf("Qué gran noche! Hemos conseguido alcanzar la impresionante cifra de %d€ donados\n", nuevoAviso.numOperador)
			fmt.Println("Muchas gracias a todos")*/
	}
}

func main() {

	objetivoGala := 20000
	duracionGala := 1 * time.Second

	numOperadores := 5

	canalCentralita := make(chan Llamada)
	canalOperador := make(chan Aviso)
	canalPresentadora := make(chan Aviso)

	go presentadora(canalPresentadora)

	for i := 0; i < numOperadores; i++ {
		go centralita(canalCentralita, duracionGala)
		go operador(i, canalCentralita, canalOperador, objetivoGala)
		go pe(canalOperador, canalPresentadora)
	}

	time.Sleep(duracionGala)
}

/*

Canal para los OPERADORES (Ejemplo en FIFO)

for i := 0; i <= numRepartidores; i++ {
		canalRepartidor := make(chan Pedido)
		go repartidor(i, canalTienda, canalRepartidor)
		repCanales = append(repCanales, canalRepartidor)
	}


   *Objetivo cambiarlo

   Time.Now() en centralita

   Apartado de variables al inicio

*/
