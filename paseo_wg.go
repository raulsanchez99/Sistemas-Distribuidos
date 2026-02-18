package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func preparacion(quien string, group *sync.WaitGroup) {
	fmt.Printf("%s ha empezado a prepararse\n", quien)
	tiempo := 60 + rand.Intn(31)
	time.Sleep(time.Duration(tiempo) * time.Millisecond)
	fmt.Printf("%s tardó %d segundos en prepararse\n", quien, tiempo)
	group.Done()
}

func zapatos(quien string, group *sync.WaitGroup) {
	defer group.Done()
	fmt.Printf("%s ha empezado a ponerse los zapatos\n", quien)
	tiempo := 35 + rand.Intn(11)
	time.Sleep(time.Duration(tiempo) * time.Millisecond)
	fmt.Printf("%s tardó %d segundos en ponerse los zapatos\n", quien, tiempo)
}

func main() {

	wg := sync.WaitGroup{}

	rand.Seed(time.Now().UnixNano())
	listo := make(chan bool)  // Fin de tareas
	alarma := make(chan bool) // Eventos de la alarma

	// Lanzar las tares de prepararse en paralelo
	go preparacion("Noa", listo)
	go preparacion("Alex", listo)

	// Esperar a que ambos terminen de prepararse
	wg.Wait()

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

	wg.Add(2)
	// Lanzar las tares de ponese los zapatos en paralelo (también en paralelo a la alarma)
	go zapatos("Noa", listo)
	go zapatos("Alex", listo)

	// Esperamos a los zapatos para salir de la casa
	wg.Wait()
	fmt.Println("Saliendo y cerrando puerta")
}
