package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func preparacion(nombre string, calle int, group *sync.WaitGroup) {
	defer group.Done()
	fmt.Printf("%s saldrá desde la calle %d\n", nombre, calle)
}

func carrera(nombre string, group *sync.WaitGroup) {
	defer group.Done()
	tiempo := 50 + rand.Intn(21)
	time.Sleep(time.Duration(tiempo) * time.Millisecond)
	fmt.Printf("%s ha terminado con un tiempo de %dms \n", nombre, tiempo)
}

func main() {

	rand.Seed(time.Now().UnixNano())
	wg := sync.WaitGroup{}

	fmt.Println("¡Empieza la gran final de 200m mariposa femenino en los Juegos Olímpicos de Río 2016!")

	nombres := [4]string{"Natsumi Hoshi", "Mireia Belmonte", "Cammile Adams", "Madeline Groves"}

	wg.Add(4)

	for i, n := range nombres {
		go preparacion(n, i, &wg)
	}
	wg.Wait()

	fmt.Println("Comienza la carrera!")
	wg.Add(4)
	for _, n := range nombres {
		go carrera(n, &wg)
	}

	wg.Wait()
	fmt.Println("La carrera ha terminado")
}
