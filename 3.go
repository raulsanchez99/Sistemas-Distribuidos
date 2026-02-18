package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var raciones int
var numExploradores = 7
var racionesPorXP = 5
var numCanibales = 9
var racionesMutex sync.Mutex
var vacia sync.Mutex
var llena sync.Mutex

func cocinero(aviso chan bool) {
	for {
		fmt.Println("El cocinero espera")
		vacia.Lock()
		if numExploradores == 0 {
			fmt.Println("No quedan exploradores")
			close(aviso)
			llena.Unlock()
			return
		} else {
			fmt.Println("El cocinero se despierta y se pone a cocinar")
			time.Sleep(time.Duration(1000+rand.Intn(1001)) * time.Millisecond)
			numExploradores -= 1
			raciones += racionesPorXP
			fmt.Println("El cocinero termina de cocinar")
			llena.Unlock()
		}
	}
}

func canibal(i int, aviso chan bool) {

	for {
		racionesMutex.Lock()
		select {
		case <-aviso:
			fmt.Println("El canibal", i, "ha terminado")
			racionesMutex.Unlock()
			return
		default:
			if raciones == 0 {
				fmt.Println("El canibal", i, "avisa al cocinero")
				//aviso <- true
				vacia.Unlock()
				llena.Lock()
				if raciones == 0 {
					racionesMutex.Unlock()
					continue
				}
			}
			raciones -= 1
			racionesMutex.Unlock()
			fmt.Println("El canibal", i, "se pone a comer")
			time.Sleep(time.Duration(500+rand.Intn(501)) * time.Millisecond)
			fmt.Println("El canibal", i, "se va a trabajar")
			time.Sleep(time.Duration(500+rand.Intn(501)) * time.Millisecond)
		}
	}
}

func main() {

	aviso := make(chan bool)
	vacia.Lock()
	llena.Lock()

	go cocinero(aviso)
	for i := 1; i < numCanibales; i++ {
		go canibal(i, aviso)
	}

	time.Sleep(30 * time.Second)
}
