package main

import (
	"fmt"
	"math/rand"
	"time"
)

type plato struct {
	nombre string
	precio int
}

func cocinero(cocineroListo chan bool, menu []plato, mesa chan plato, ordenes chan int) {

	fmt.Println("El cocinero se está poniendo el delantal.")
	time.Sleep(200 * time.Millisecond)

	cocineroListo <- true
	<-ordenes

	fmt.Println("El cocinero comienza a trabajar!")

	// Bucle para cocinar platos aleatoriamente
	for {

		platoElegido := menu[rand.Intn(len(menu))]
		time.Sleep(300 * time.Millisecond)

		numPlatos := rand.Intn(3) + 1

		// Mostrar el número de platos que se prepararon
		for i := 0; i < numPlatos; i++ {
			fmt.Printf("Una nueva ración de %s ha salido de cocina.\n", platoElegido.nombre)
			mesa <- platoElegido
		}
	}
}

func cliente(numMesas chan int, empiezaTurno chan bool, clienteListo chan int, mesa chan plato, pagos chan int, num int) {

	m := <-numMesas

	// Cliente ha llegado
	fmt.Printf("El cliente %d ha llegado al restaurante y espera su turno.\n", num)
	empiezaTurno <- true

	fmt.Printf("El cliente %d se ha sentado en una mesa, y espera a que le sirvan.\n", num)

	//fmt.Printf("Quedan %d mesas \n", m)

	// El cliente espera a que le sirvan la comida y luego come
	platoElegido := <-mesa
	fmt.Printf("Al cliente %d se le ha servido %s. Comienza a comer.\n", num, platoElegido.nombre)
	tiempoComiendo := 200 + rand.Intn(201)
	time.Sleep(time.Duration(tiempoComiendo) * time.Millisecond)
	fmt.Printf("El cliente %d ha terminado de comer. Se levanta de su mesa y se pone en cola para pagar.\n", num)

	// Cliente ha terminado de comer
	clienteListo <- num

	// Incrementar el número de mesas disponibles
	numMesas <- m
	//fmt.Printf("Hay %d mesas \n", m)

	//Calculamos el precio del plato
	costo := platoElegido.precio
	fmt.Printf("El cliente %d ha pagado por su comida %d euros.\n", num, costo)
	pagos <- costo

}

func cajero(pagos chan int, clientes int, cajeroTerminado chan bool) {

	//Sumamos todos los precios de los platos consumidos
	total := 0
	for i := 0; i < clientes; i++ {
		pago := <-pagos
		total += pago
	}
	fmt.Println("Todos los clientes han pagado. Se cierra la caja.")
	fmt.Printf("El Paraíso de la Papa cierra por hoy con una facturación de %d euros.\n", total)
	cajeroTerminado <- true
}

func main() {

	menu := []plato{
		{"Papas con mojo", 10},
		{"Tortilla de patatas", 12},
		{"Patatas rellenas", 8},
		{"Gnocchi", 15},
	}

	numClientes := 5 + rand.Intn(5)

	cocineroListo := make(chan bool)
	empiezaTurno := make(chan bool)
	clientesListos := make(chan int, numClientes)
	ordenes := make(chan int)
	mesa := make(chan plato, 1)
	mesasDisponibles := 3
	numMesas := make(chan int, 3)
	pagos := make(chan int)
	cajeroTerminado := make(chan bool)

	fmt.Println("Un día más, El Paraíso de la Papa abre sus puertas")

	// Llamadas a las funciones del cocinero y cajero
	go cocinero(cocineroListo, menu, mesa, ordenes)
	go cajero(pagos, numClientes, cajeroTerminado)

	<-cocineroListo

	//Envio tres mesas al canal
	for i := 0; i < 3; i++ {
		numMesas <- mesasDisponibles
	}

	// Simulación de los clientes
	for i := 1; i <= numClientes; i++ {
		go cliente(numMesas, empiezaTurno, clientesListos, mesa, pagos, i)
		<-empiezaTurno
		//<-numMesas
		if i == 1 {
			ordenes <- i
		}
		time.Sleep(time.Duration(150+rand.Intn(101)) * time.Millisecond)
	}
	<-cajeroTerminado
}
