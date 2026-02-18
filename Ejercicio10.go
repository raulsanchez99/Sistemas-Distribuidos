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

func (nuevo *Vigilante) incrementarturno() {
	nuevo.turnos++
}

func (nuevo *Vigilante) cambioestado() {
	nuevo.trabajando = !nuevo.trabajando
}

func añadirvigilante(equipo *[]Vigilante) {
	var v Vigilante
	rellenarInfo(&v)
	*equipo = append(*equipo, v)
}

func eliminarvigilante(equipo *[]Vigilante) {
	var indice int
	fmt.Printf("Que vigilante quieres eliminar [0-%d]: ", len(*equipo))
	fmt.Scanf("%d", &indice)
	*equipo = append((*equipo)[:indice], (*equipo)[indice+1:]...)
}

func main() {
	var equipo []Vigilante

	for {
		var opcion int
		fmt.Println("Elige una opción:")
		fmt.Println("Opción 1: Añadir vigilante")
		fmt.Println("Opción 2: Eliminar vigilante")
		fmt.Println("Opción 3: Info")
		fmt.Println("Opción 4: Turno")
		fmt.Scanf("%d", &opcion)

		switch opcion {
		case 1:
			añadirvigilante(&equipo)
		case 2:
			eliminarvigilante(&equipo)
		case 3:
			trabajador_activo := ""
			for i := 0; i < len(equipo); i++ {
				fmt.Printf("%s ha trabajado %d turnos\n", equipo[i].nombre, equipo[i].turnos)
				if equipo[i].trabajando {
					trabajador_activo = equipo[i].nombre
				}
			}

			fmt.Printf("El turno actual lo cubre %s\n", trabajador_activo)

		case 4:
			aleatorio := rand.Int() % 4 // Obtener un entero aleatorio entre 0 y 3
			for j := 0; j < len(equipo); j++ {
				if j == aleatorio {
					if !equipo[j].trabajando {
						equipo[j].cambioestado()
					}
					equipo[j].incrementarturno()
				} else {
					if equipo[j].trabajando {
						equipo[j].cambioestado()
					}
				}
			}
		}
	}
}
