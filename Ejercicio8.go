package main

import (
	"fmt"
	"math/rand"
)

type Vigilante struct {
	nombre     string
	turnos     int
	trabajando bool
}

func rellenarInfo(vig *Vigilante) {
	fmt.Println("Introduce el nombre:")
	fmt.Scanf("%s", &vig.nombre)
	vig.turnos = 0
	vig.trabajando = false
}

func main() {
	var equipo [4]Vigilante
	for i := 0; i < len(equipo); i++ {
		rellenarInfo(&equipo[i])
	}

	for i := 0; i < 10; i++ {
		aleatorio := rand.Int() % 4 // Obtener un entero aleatorio entre 0 y 3
		for j := 0; j < len(equipo); j++ {
			if j == aleatorio {
				equipo[j].trabajando = true
				equipo[j].turnos += 1
			} else {
				equipo[j].trabajando = false
			}
		}
	}

	trabajador_activo := ""
	for i := 0; i < len(equipo); i++ {
		fmt.Printf("%s ha trabajado %d turnos\n", equipo[i].nombre, equipo[i].turnos)
		if equipo[i].trabajando {
			trabajador_activo = equipo[i].nombre
		}
	}

	fmt.Printf("El turno actual lo cubre %s\n", trabajador_activo)

}
