package main

import "fmt"

func main() {

	carrito := make(map[string]int) // Mapa vacío con make

	var producto string

	for i := 0; i < 6; i++ {
		fmt.Println("Que producto quieres comprar?:")
		fmt.Scanf("%s\n", &producto)

		p, ok := carrito[producto]
		if ok {
			carrito[producto] = p + 1
		} else {
			carrito[producto] = p
		}
	}

	fmt.Println(carrito)
	for k, v := range carrito {
		fmt.Printf("%s, %d\n", k, v)
	}

	fmt.Println("Borra un producto")
	fmt.Scanf("%s", &producto)

	_, ok := carrito[producto]
	if ok {
		delete(carrito, producto)
	} else {
		fmt.Println("El producto no existe")
	}

	fmt.Println(carrito)
}
