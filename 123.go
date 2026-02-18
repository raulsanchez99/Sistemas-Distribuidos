package main

import (
	"fmt"
	"time"
	"math/rand"
)

const numeroinicial = 0
var tiempoEspera = time.Duration(10+rand.Intn(21))*time.Millisecond
const numTotal = 5


func numero(numeroinicial int){
	suma := numeroinicial + 1
	fmt.Printf("La suma es: %d\n", suma)

}


func main(){

	for i:= 0; i < numTotal; i++{
		fmt.Printf("-- Comienza la suma.  NUM: %d\n", numeroinicial)
		go numero(numeroinicial)
		time.Sleep(tiempoEspera)
	}

}
