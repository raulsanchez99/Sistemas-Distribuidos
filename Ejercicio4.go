package main

import (
	"fmt"
	"strconv"
)

func main() {
	for n := 0; n <= 1000; {

		var entrada string
		// Solicitar al usuario un numero
		fmt.Print("Ingresa un numero: ")
		fmt.Scan(&entrada)

		switch entrada {
		case "stop", "parar":
			break
		default:
			num, _ := strconv.ParseFloat(entrada, 10)
			if num >= 0 {
				fmt.Printf("El numero %f es positivo")
			} else {
				fmt.Printf("El numero %f es negativo ")
			}

		}
	}
}
