package main

import (
	"fmt"
	"math"
	"time"
)

// Función para determinar si un número es primo
func esPrimo(numero int) bool {
	if numero < 2 {
		return false
	}
	raizCuadrada := int(math.Sqrt(float64(numero)))
	for i := 2; i <= raizCuadrada; i++ {
		if numero%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	limiteSuperior := 10000000
	numerosPrimos := []int{}

	// Medir el tiempo de ejecución
	inicio := time.Now()

	// Encontrar números primos en el rango especificado
	for i := 2; i <= limiteSuperior; i++ {
		if esPrimo(i) {
			numerosPrimos = append(numerosPrimos, i)
		}
	}

	// Calcular el tiempo transcurrido
	tiempoTranscurrido := time.Since(inicio)

	// Mostrar el resultado en el formato especificado
	fmt.Printf("Números primos encontrados (caso #): %d\n", len(numerosPrimos))
	fmt.Printf("Tiempo transcurrido (ms): %d\n", tiempoTranscurrido.Milliseconds())
}
