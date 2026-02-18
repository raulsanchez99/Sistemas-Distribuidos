package main

import (
	"fmt"
	"math"
)

// Definición del tipo de dato para errores
type MiError struct {
	valor   float64
	mensaje string
}

func (e *MiError) Error() string {
	return e.mensaje
}

// Función para calcular la raíz cuadrada y manejar errores
func raiz(numero float64) (float64, error) {
	if numero < 0 {
		return 0, &MiError{valor: numero, mensaje: "No se puede calcular la raíz cuadrada de un número negativo"}
	}
	return math.Sqrt(numero), nil
}

func main() {

	v, err := raiz(12)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

	v, err = raiz(-16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
}
