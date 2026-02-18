package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mapa = make(map[int]int)
var mutex sync.Mutex

func incrementarcontador(indice int, final chan bool) {
	for {
		select {
		case <-final:
			fmt.Println("Saliendo")
			return
		default:
			time.Sleep(time.Duration(50+rand.Intn(251)) * time.Millisecond)
			mutex.Lock()
			mapa[indice]++
			mutex.Unlock()
		}
	}
}

func main() {

	var fin = make(chan bool)

	for i := 0; i < 7; i++ {
		go incrementarcontador(i, fin)
	}
	time.Sleep(5 * time.Second)
	close(fin)

	mutex.Lock()
	for k, v := range mapa {
		fmt.Printf("%d, %d\n", k, v)
	}
	mutex.Unlock()
}
