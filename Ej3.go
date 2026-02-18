package main

import (
	"fmt"
	"math/rand"
	"time"
)

func comer(nombre string, mesa chan string, fin chan bool) {
	for porcion := range mesa {
		fmt.Println(nombre, "está comiendo", porcion)
		time.Sleep(time.Duration(30+rand.Intn(151)) * time.Millisecond)
	}
	fin <- true
}

func main() {
	comensales := [4]string{"Julia", "Antonio", "María", "Daniel"}
	raciones := [5]string{"chorizo", "pimientos de padrón", "croquetas", "bravas", "chopitos"}

	/*porciones := [5]int{
		5 + rand.Intn(6),
		5 + rand.Intn(6),
		5 + rand.Intn(6),
		5 + rand.Intn(6),
		5 + rand.Intn(6),
	}*/

	cuenta := make(map[string]int)
	cuenta["chorizo"] = 5 + rand.Intn(6)
	cuenta["pimientos de padrón"] = 5 + rand.Intn(6)
	cuenta["croquetas"] = 5 + rand.Intn(6)
	cuenta["bravas"] = 5 + rand.Intn(6)
	cuenta["chopitos"] = 5 + rand.Intn(6)

	num_porciones := 0
	for _, v := range cuenta {
		num_porciones += v
	}
	fmt.Println("Total porciones", num_porciones)

	mesa := make(chan string, num_porciones)
	fin := make(chan bool)

	for i := 0; i < num_porciones; i++ {
		/*idx := rand.Intn(len(raciones))
		if porciones[idx] > 0 {
			mesa <- raciones[idx]
			porciones[idx] -= 1
		} else {
			i--
		}*/

		plato := raciones[rand.Intn(len(raciones))]
		porciones_restantes := cuenta[plato]
		if porciones_restantes > 0 {
			mesa <- plato
			cuenta[plato] -= 1
		} else {
			i--
		}
	}
	close(mesa)

	fmt.Println("Que aproveche")
	for _, nombre := range comensales {
		go comer(nombre, mesa, fin)
	}

	for range comensales {
		<-fin
	}

	fmt.Println("Ha sido una cena increíble")
}
