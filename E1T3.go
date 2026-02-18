package main

import (
	"fmt"
	"math/rand"
	"time"
)

//rand.Intn(101) numero aleatorio entre 0 y un valor
//time.Sleep(time.Duration(rand(Intn(101))*time.Millisecond)

func prepararse(nombre string, listos chan bool) {

	fmt.Printf("%s ha empezado a prepararse\n", nombre)
	tiempo := 60 + rand.Intn(31)
	time.Sleep(time.Duration(tiempo) * time.Millisecond)
	fmt.Printf("%s tardó %d segundos en prepararse\n", nombre, tiempo)
	listos <- true
}

func ponersezapatos(nombre string, listos chan bool) {
	fmt.Printf("%s ha empezado a ponerse los zapatos\n", nombre)
	tiempo := 35 + rand.Intn(11)
	time.Sleep(time.Duration(tiempo) * time.Millisecond)
	fmt.Printf("%s tardó %d segundos en ponerse los zapatos\n", nombre, tiempo)
	listos <- true
}

func main() {

	listo := make(chan bool)
	alarma := make(chan bool)

	fmt.Println("Vamos de paseo!")

	go prepararse("Noa", listo)
	go prepararse("Alex", listo)
	<-listo
	<-listo

	go func(alarma chan bool) {
		fmt.Println("Activando alarma")
		alarma <- false

		timer := time.NewTimer(60 * time.Millisecond)
		ticker := time.NewTicker(5 * time.Millisecond)
		for {
			select {
			case <-timer.C:
				ticker.Stop()
				fmt.Println("Alarma activada")
				alarma <- true
				return
			case <-ticker.C:
				fmt.Println("Beep")
			}
		}
	}(alarma)

	<-alarma
	go ponersezapatos("Noa", listo)
	go ponersezapatos("Alex", listo)
	<-listo
	<-listo

	fmt.Println("Saliendo y cerrando la puerta")
	<-alarma

}
