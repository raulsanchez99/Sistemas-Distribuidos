package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numPedidos = 7
const numRepartidores = 4

type Pedido struct {
	id               int
	duracionEstimada time.Duration
	duracionReal     time.Duration
}

func tienda(canalTienda chan Pedido) {
	for i := 0; i <= numPedidos; i++ {
		tiempo := time.Duration(11+rand.Intn(90)) * time.Minute
		p := Pedido{id: i, duracionEstimada: tiempo}
		fmt.Printf("Tienda: nuevo pedido %d recibido. Tiempo estimado de entrega %v\n", p.id, p.duracionEstimada)
		canalTienda <- p
	}
	close(canalTienda)
}

func repartidor(id int, canalTienda <-chan Pedido, canalRepartidor chan<- Pedido) {
	for p := range canalTienda {
		fmt.Printf("Repartidor %d: repartiendo %d ... \n", id, p.id)
		p.duracionReal = p.duracionEstimada + time.Duration(-10+rand.Intn(20))*time.Minute
		tiempoEntregaFake := p.duracionReal / (60 * 1000)
		time.Sleep(tiempoEntregaFake)
		canalRepartidor <- p
		fmt.Printf("Repartidor %d ha vuelto a la tienda \n", id)
	}
	close(canalRepartidor) //Repartidor indica que no va a entegar más paquetes
}

func estadistico(repartidores []chan Pedido, canalEstadistico chan<- time.Duration) {
	tiempoTotal := time.Duration(0)
	tiempoMutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(repartidores))

	for _, canal := range repartidores { //Iterar sobre el slice repartidores
		go func() {
			for p := range canal {
				fmt.Printf("Estadistico: el pedido %d ha sido entregado en %v\n", p.id, p.duracionReal)
				tiempoMutex.Lock()
				tiempoTotal += p.duracionReal
				tiempoMutex.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Estadistico: todos los pedidos han sido procesados. ")
	canalEstadistico <- tiempoTotal / time.Duration(numPedidos)
}

func main() {

	canalTienda := make(chan Pedido)

	go tienda(canalTienda)

	var repCanales []chan Pedido

	for i := 0; i <= numRepartidores; i++ {
		canalRepartidor := make(chan Pedido)
		go repartidor(i, canalTienda, canalRepartidor)
		repCanales = append(repCanales, canalRepartidor)
	}

	canalEstadistico := make(chan time.Duration)
	go estadistico(repCanales, canalEstadistico)
	tiempoPromedio := <-canalEstadistico
	fmt.Printf("Main: tiempo promedio de entrega: %v\n", tiempoPromedio)

	/*

		// Tiempo de entrega estimado
		tiempo := time.Duration(50+rand.Intn(100)) * time.Minute
		fmt.Printf("%v\n", tiempo)

		siesta := tiempo / (60 * 1000)
		fmt.Printf("Siesta %v\n", siesta)*/

}
