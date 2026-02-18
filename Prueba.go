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
	// El cocinero se está poniendo el delantal
	fmt.Println("El cocinero se está poniendo el delantal.")
	time.Sleep(200 * time.Millisecond)

	// Esperar a que un cliente haga un pedido
	cocineroListo <- true
	<-ordenes
	fmt.Println("El cocinero comienza a trabajar!")

	// Iniciar el bucle para la preparación continua de platos
	for {
		// Seleccionar un plato aleatorio del menú
		platoElegido := menu[rand.Intn(len(menu))]

		// Simular el tiempo de preparación de la receta
		time.Sleep(300 * time.Millisecond)

		// Determinar cuántos platos se están preparando entre 1 y 3
		numPlatos := rand.Intn(3) + 1

		// Mostrar el número de platos que se están preparando
		for i := 0; i < numPlatos; i++ {
			fmt.Printf("Una nueva ración de %s ha salido de cocina.\n", platoElegido.nombre)
			// Enviar el plato elegido al canal de la mesa
			mesa <- platoElegido
		}
	}
}

func cliente(mesasDisponibles chan int, empiezaTurno chan bool, clienteListo chan int, mesa chan plato, pagos chan int, num int) {
	// Obtener el canal mesasDisponibles
	m := <-mesasDisponibles

	// Avisar que el cliente ha llegado
	fmt.Printf("El cliente %d ha llegado al restaurante y espera su turno.\n", num)
	empiezaTurno <- true

	// Comprobar el número de mesas disponibles
	for m == 0 {
		//numMesa := m > 0
		m = <-mesasDisponibles
	}

	if m >= 1 && m <= 3 {
		// Si hay mesas disponibles, el cliente se sienta en la mesa y espera a que le sirvan
		fmt.Printf("El cliente %d se ha sentado en una mesa, y espera a que le sirvan.\n", num)
		numMesa := m - 1
		mesasDisponibles <- numMesa
		fmt.Printf("Quedan %d mesas \n", numMesa)

		// El cliente espera a que le sirvan la comida y luego come
		platoElegido := <-mesa
		fmt.Printf("Al cliente %d se le ha servido %s. Comienza a comer.\n", num, platoElegido.nombre)
		tiempoComiendo := 200 + rand.Intn(201)
		time.Sleep(time.Duration(tiempoComiendo) * time.Millisecond)
		fmt.Printf("El cliente %d ha terminado de comer. Se levanta de su mesa y se pone en cola para pagar.\n", num)

		// Notificar que el cliente ha terminado de comer
		clienteListo <- num

		// Incrementar el número de mesas disponibles
		m = numMesa + 1
		//mesasDisponibles <- m
		fmt.Printf("Hay %d mesas \n", m)

		costo := platoElegido.precio
		fmt.Printf("El cliente %d ha pagado por su comida %d euros.\n", num, costo)
		pagos <- costo
	}
}

func cajero(pagos chan int, clientes int, cajeroTerminado chan bool) {
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
	fmt.Println("Un día más, El Paraíso de la Papa abre sus puertas")

	menu := []plato{
		{"Papas con mojo", 10},
		{"Tortilla de patatas", 12},
		{"Patatas rellenas", 8},
		{"Gnocchi", 15},
	}

	numClientes := 5 + rand.Intn(5)

	cocineroListo := make(chan bool)
	clientesListos := make(chan int, numClientes)
	ordenes := make(chan int)
	pagos := make(chan int)
	cajeroTerminado := make(chan bool)
	mesasDisponibles := make(chan int, 3)
	mesasDisponibles <- 3
	empiezaServicio := make(chan bool)

	// Crear un canal para la comunicación entre el cocinero y los clientes
	mesa := make(chan plato, 1)

	// Llamar a la función del cocinero
	go cocinero(cocineroListo, menu, mesa, ordenes)
	// Llamar a la función del cajero
	go cajero(pagos, numClientes, cajeroTerminado)

	// Esperar a que el cocinero termine de ponerse el delantal
	<-cocineroListo

	// Simulación de clientes
	for i := 1; i <= numClientes; i++ {
		go cliente(mesasDisponibles, empiezaServicio, clientesListos, mesa, pagos, i)
		<-empiezaServicio
		if i == 1 {
			ordenes <- i
		}
		time.Sleep(time.Duration(150+rand.Intn(101)) * time.Millisecond)
	}
	<-cajeroTerminado
}
