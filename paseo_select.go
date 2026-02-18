package main

import (
	"fmt"
	"math/rand"
	"time"
)

func preparacion(quien string, ready chan bool) {
	fmt.Printf("%s ha empezado a prepararse\n", quien)
	tiempo := 60 + rand.Intn(31)
	time.Sleep(time.Duration(tiempo) * time.Millisecond)
	fmt.Printf("%s tardó %d segundos en prepararse\n", quien, tiempo)
	ready <- true
}

func zapatos(quien string, ready chan bool) {
	fmt.Printf("%s ha empezado a ponerse los zapatos\n", quien)
	tiempo := 35 + rand.Intn(11)
	time.Sleep(time.Duration(tiempo) * time.Millisecond)
	fmt.Printf("%s tardó %d segundos en ponerse los zapatos\n", quien, tiempo)
	ready <- true
}

func main() {

	rand.Seed(time.Now().UnixNano())
	listo := make(chan bool)  // Fin de tareas
	alarma := make(chan bool) // Eventos de la alarma

	// Lanzar las tares de prepararse en paralelo
	go preparacion("Noa", listo)
	go preparacion("Alex", listo)

	// Esperar a que ambos terminen de prepararse
	<-listo
	<-listo

	// go activaralarma(alarma)
	// Es lo mismo
	go func(alarma chan bool) {
		fmt.Println("Activando alarma")
		alarma <- false // Pulsar el botón: indica que se ha iniciado la cuenta atrás

		timer := time.NewTimer(60 * time.Millisecond)
		ticker := time.NewTicker(5 * time.Millisecond)
		for {
			select {
			case <-timer.C:
				ticker.Stop()
				fmt.Println("Alarma activada")
				alarma <- true // indica que se ha activado (para terminar el programa)
				return
			case <-ticker.C:
				fmt.Println("BEEP")
			}
		}

	}(alarma)

	<-alarma // Esperamos a pulsar el botón para ponernos los zapatos

	// Lanzar las tares de ponese los zapatos en paralelo (también en paralelo a la alarma)
	go zapatos("Noa", listo)
	go zapatos("Alex", listo)

	// Esperamos a los zapatos para salir de la casa
	<-listo
	<-listo
	fmt.Println("Saliendo y cerrando puerta")

	// Espera a que termine la goroutina de la alarma para cerrar el programa
	<-alarma

}
