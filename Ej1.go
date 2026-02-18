package main

import (
	"fmt"
	"time"
)

func atleta(num int, testigo chan bool, terminados chan bool) {

	//Salta a la pista
	fmt.Printf("La atleta %d está posicionada en su línea\n", num)

	//Recogemos testigo
	<-testigo
	fmt.Printf("La atleta %d corre con el testigo\n", num)

	//Corremos la mitas de la carrera
	time.Sleep(time.Millisecond * 100)

	//Salta la siguiente atleta
	if num < 4 {
		go atleta(num+1, testigo, terminados)
	}

	//Corremos la mitas de la carrera
	time.Sleep(time.Millisecond * 100)

	//Entrega el testigo
	if num < 4 {
		testigo <- false
		fmt.Printf("La atleta %d pasa el testigo a la atleta%d\n", num, num+1)
	} else {

	}
	/*
		1.Posicionar (print)
		2.Tomar testigo
		3.Empezar a correr (print)
		4.Correr (sleep)
		5.Siguiente atleta
		6.Correr (sleep)
		7.Pasar testigo
	*/
}

func main() {
	relevo := make(chan bool)
	final := make(chan bool)
	fmt.Println("Bienvenid@s a la final femenina de 4x200")
	fmt.Println("Comienza la carrera!")

	go atleta(1, relevo, final)

	relevo <- false
	time.Sleep(time.Second * 5)

	//Esperamos a que termine

	fmt.Println("Carrera terminada")

}
