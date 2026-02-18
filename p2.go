package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
   vars{
   	objetivoGala = 20000
   	duracionGala = 4 * time.Second
   	numOperadores = 5
   }
*/
type Llamada struct {
	valor    int
	duracion time.Duration
}

type Aviso struct {
	numOperador   int
	valorDonacion int
	recaudacion   int
}

func centralita(llamadas chan Llamada, duracionTotal time.Duration, fin chan bool) {
	//tiempoInicio := time.Now()

	//Crear un timer de 4 segundos, cuando timer es 4, enviamos el timer a la presentadora
	//timert := time.NewTimer(duracionGala)

	//for time.Since(tiempoInicio) < duracionTotal {
	for {
		nuevaLlamada := Llamada{valor: 1 + rand.Intn(1001), duracion: time.Duration(30+rand.Intn(51)) * time.Millisecond}
		//fmt.Println("PII")
		llamadas <- nuevaLlamada
		time.Sleep(time.Duration(25+rand.Intn(26)) * time.Millisecond)
		//<-timert.C

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
		//objetivo -= nuevaLlamada.valor
		total += nuevaLlamada.valor
		nuevoAviso := Aviso{numOperador: num, valorDonacion: nuevaLlamada.valor, recaudacion: total}
		canalOperador <- nuevoAviso
		sumaMutex.Unlock()

		//fmt.Printf("Operador %d restó %d€ al objetivo de la gala. Nuevo objetivo: %d\n", num, nuevaLlamada.valor, *objetivo)
		//fmt.Printf("Qué gran noche! Hemos conseguido alcanzar la impresionante cifra de %d € donados", total)
		//Ahora tenemos que ver si el valor del objetivo se alcanzó
		//canalOperador <- nuevoAviso

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

func presentadora(canalPresentadora <-chan Aviso, fin chan bool) {
	fmt.Println("¡Bienvenidos a la gala benéfica contra el cambio climático!")
	fmt.Println("La cantidad total recaudada irá destinada a la replantación de árboles en el Amazonas.")
	fmt.Println("En este momento... abrimos las líneas!")

	timer := time.NewTimer(1 * time.Second)
	timert := time.NewTimer(4 * time.Second)
	ticker := time.NewTicker(200 * time.Millisecond)
	objetivoGala := 20000

	for nuevoAviso := range canalPresentadora {
		select {
		case <-timert.C:
			fmt.Println("Y... cerramos líneas! A partir de este momento ya no recibiremos más llamadas")
			fmt.Printf("Qué gran noche! Hemos conseguido alcanzar la impresionante cifra de %d€ donados\n", nuevoAviso.recaudacion)
			fmt.Println("Muchas gracias a todos")
			return

		case <-canalPresentadora:
			fmt.Printf("Nuestro operador %d acaba de recibir una generosa donación de %d€!\n", nuevoAviso.numOperador, nuevoAviso.valorDonacion)
			timer.Reset(1 * time.Second)
			//return

		case <-ticker.C:
			if objetivoGala >= nuevoAviso.recaudacion {
				queda := objetivoGala - nuevoAviso.recaudacion
				fmt.Printf("Estamos a tan solo %d € de llegar al objetivo de la noche\n", queda)
				//return
			} else {
				sobra := nuevoAviso.recaudacion - objetivoGala
				fmt.Printf("Llevamos recaudados %d € más que el objetivo inicial de la noche!\n", sobra)
				//return
			}

		case <-timer.C:
			fmt.Println("Continúen llamando para hacer sus donaciones por esta buena causa")
			timer.Reset(1 * time.Second)
			//return
		}
	}

}

func main() {

	objetivoGala := 20000
	duracionGala := 1 * time.Second

	numOperadores := 5

	canalCentralita := make(chan Llamada)
	canalOperador := make(chan Aviso)
	fin := make(chan bool)
	canalPresentadora := make(chan Aviso)

	go presentadora(canalPresentadora, fin)

	for i := 0; i < numOperadores; i++ {
		go centralita(canalCentralita, duracionGala, fin)
		go operador(i, canalCentralita, canalOperador, objetivoGala)
		go pe(canalOperador, canalPresentadora)
	}

	time.Sleep(5 * time.Second)
}

/*

*   Es decir, operador le pasa a pe, numOperador, valor de la donacion y total recaudado
   	Pe debera restar el total recaudado al objetivo de la gala y eso se lo pasara a la presentadora para que lo anuncie
   	O eso que lo haga el operador y en ese mismo canal pase todos los datos

*   ??? Otro canal distinto para presentadora , que reciba valor donacion, total recaudado y cuanto queda para alcanzar el objetivo

   Canal para los OPERADORES (Ejemplo en FIFO)

   for i := 0; i <= numRepartidores; i++ {
   		canalRepartidor := make(chan Pedido)
   		go repartidor(i, canalTienda, canalRepartidor)
   		repCanales = append(repCanales, canalRepartidor)
   	}

   en pe si donacion > 750 mando mensaje a presentadora, inicio contador


*     *Objetivo cambiarlo

      Time.Now() en centralita

   case <-fin:
   Imprimimos mensajes de fin
   return

      Apartado de variables al inicio


time.ticker -> tick tick tick
time.timer -> ------fin


timer:= time.newtimer(second)
   select
   case <-donaciones alta
		timer.Reset(time.second)
   case 200 ms
   case <-timer.C:
		dona mas
		reset
	case despedir


- mirar las cuentas del if y la resta //TERMINADO??
- revisar el timer
- timer total 4 segundos en centralita, para que no siga recibiendo llamadas
y asi mata todos los programas.

*/
