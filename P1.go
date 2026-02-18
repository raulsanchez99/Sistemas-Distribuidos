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

func cocinero(cocineroListo chan bool, menu []plato, mesa chan plato, clienteEspera chan bool) {
	// El cocinero se está poniendo el delantal
	fmt.Println("El cocinero se está poniendo el delantal.")
	time.Sleep(200 * time.Millisecond)

	// Indicar que el cocinero está listo
	cocineroListo <- true

	<-clienteEspera
	fmt.Println("El cocinero comienza a trabajar!")

	// Iniciar el bucle para la preparación continua de platos
	for {
		// Seleccionar un plato aleatorio del menú
		platoElegido := menu[rand.Intn(len(menu))]

		// Simular el tiempo de preparación de la receta
		time.Sleep(300 * time.Millisecond)

		// Determinar cuántos platos se están preparando entre 1 y 3
		numPlatos := 1 + rand.Intn(3)

		// Mostrar el número de platos que se están preparando
		for i := 0; i < numPlatos; i++ {
			fmt.Printf("Una nueva ración de %s ha salido de cocina.\n", platoElegido.nombre)
			// Enviar el plato elegido al canal de la mesa
			mesa <- platoElegido
		}
	}
}

func cliente(num int, mesa chan plato, clienteListo chan bool, clienteEspera chan bool) {
	// El cliente llega al restaurante y espera pacientemente
	fmt.Printf("El cliente %d ha llegado al restaurante y espera su turno.\n", num)
	// El cliente se sienta en la mesa y espera a que le sirvan
	fmt.Printf("El cliente %d se ha sentado en una mesa, y espera a que le sirvan.\n", num)

	if num == 1 {
		clienteEspera <- true
	}

	// El cliente espera a que le sirvan la comida y luego come
	platoElegido := <-mesa
	fmt.Printf("Al cliente %d se le ha servido %s. Comienza a comer.\n", num, platoElegido.nombre)
	tiempoComiendo := 200 + rand.Intn(201)
	time.Sleep(time.Duration(tiempoComiendo) * time.Millisecond)
	fmt.Printf("El cliente %d ha terminado de comer  Se levanta de su mesa y se pone en cola para pagar..\n", num)

	// Notificar que el cliente ha terminado de comer
	clienteListo <- true
}

func main() {

	fmt.Println("Un día más, El Paraíso de la Papa abre sus puertas")

	menu := []plato{
		{"Papas con mojo", 10},
		{"Tortilla de patatas", 12},
		{"Patatas rellenas", 8},
		{"Gnocchi", 15},
	}

	numClientes := 5 + rand.Intn(5)
	numMesas := 3

	cocineroListo := make(chan bool)
	clientesListos := make(chan bool, numClientes)
	clienteEspera := make(chan bool)

	// Crear un canal para la comunicación entre cocinero y clientes
	mesa := make(chan plato, numMesas)

	// Llamar a la función del cocinero
	go cocinero(cocineroListo, menu, mesa, clienteEspera)

	// Esperar a que el cocinero termine de ponerse el delantal
	<-cocineroListo

	// Simulación de clientes
	for i := 1; i <= numClientes; i++ {
		go cliente(i, mesa, clientesListos, clienteEspera)
		time.Sleep(time.Duration(150+rand.Intn(101)) * time.Millisecond)
		<-clientesListos
	}

}
