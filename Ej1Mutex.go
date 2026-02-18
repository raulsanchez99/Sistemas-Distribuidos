package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mesa = make(map[string]int)
var mutexMesa sync.Mutex

func comer(nombre string, fin chan bool) {
	for {
		mutexMesa.Lock()
		if mesa["total"] == 0 {
			mutexMesa.Unlock()
			fmt.Println(nombre, "ha terminado")
			fin <- true
			break
		}
		raciones := [5]string{"chorizo", "pimientos de padrón", "croquetas", "bravas", "chopitos"}
		idx := rand.Intn(len(raciones))
		tapa := raciones[idx]
		if mesa[tapa] > 0 {
			mesa[tapa] -= 1
			mesa["total"] -= 1
			mutexMesa.Unlock()
			fmt.Println(nombre, "está comiendo", tapa)
			time.Sleep(time.Duration(30+rand.Intn(151)) * time.Millisecond)
		} else {
			mutexMesa.Unlock()
		}
	}

}

func main() {
	comensales := [4]string{"Julia", "Antonio", "María", "Daniel"}

	mutexMesa.Lock()
	mesa["chorizo"] = 5 + rand.Intn(6)
	mesa["pimientos de padrón"] = 5 + rand.Intn(6)
	mesa["croquetas"] = 5 + rand.Intn(6)
	mesa["bravas"] = 5 + rand.Intn(6)
	mesa["chopitos"] = 5 + rand.Intn(6)

	num_porciones := 0
	for _, v := range mesa {
		num_porciones += v
	}
	mesa["total"] = num_porciones
	mutexMesa.Unlock()
	fmt.Println("Total porciones", num_porciones)

	fin := make(chan bool)

	fmt.Println("Que aproveche")
	for _, nombre := range comensales {
		go comer(nombre, fin)
	}

	for range comensales {
		<-fin
		fmt.Println("Aviso recibido")
	}

	fmt.Println("Ha sido una cena increíble")
}
