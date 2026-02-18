package main

import (
	"fmt"
)

type Animal interface {
	hacerSonido() string
}

type Perro struct {
}

type Gato struct {
}

func (p Perro) hacerSonido() string {
	return "Guau"
}

func (g Gato) hacerSonido() string {
	return "Miau"
}

func main() {

	/*
		var lista = make([]Animal, 2)
		var animal1 Perro
		var animal2 Gato

		//lista = append(animales, animal1, animal2)

		lista[0] = animal1
		lista[1] = animal2
	*/

	// 3. Declara una lista mediante un slice para almacenar animales e inserta al menos un elemento de cada tipo
	// Creamos un slice de tipo Animal
	animales := []Animal{Perro{}, Gato{}}

	// 4. Itera el slice con range y llama al método de la interfaz para cada elemento de la lista, imprimiendo su sonido
	/*
		for i:= range lista {
			fmt.Println(animales[i].hacerSonido())
		}
	*/

	for _, animal := range animales {
		fmt.Println(animal.hacerSonido())
	}
}
