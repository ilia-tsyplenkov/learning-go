// package provide solution for conveyor
// with so many gorouitnes as you want
// finally it shows how long it has been executed
// with temporaty log printing
package main

import (
	"flag"
	"log"
	"time"
)

var channels []chan int
var firstCh = make(chan int)
var goroutNum = flag.Int("gn", 10, "goroutines number in conveyor")
var logStep = flag.Int("step", 100000, "step for printing log")

func f(i int, channels ...chan int) {
	if i%(*logStep) == 0 {
		log.Println("current index is", i)
	}
	switch len(channels) {
	case 0:
		log.Fatal("no channels provided for function in conveyor")
		return
	case 1:
		initValue := <-firstCh
		out := channels[0]
		out <- i + initValue
		return
	case 2:
		out := channels[1]
		in := channels[0]
		val := <-in
		out <- i + val
		return
	default:
		log.Fatal("too many channels for function in conveyor")
	}
}

func main() {
	flag.Parse()
	channels = make([]chan int, *goroutNum)
	start := time.Now()
	for i := range channels {
		ch := make(chan int)
		channels[i] = ch
		if i == 0 {
			go f(i, firstCh, ch)
			continue
		}
		go f(i, channels[i-1], ch)
	}

	firstCh <- 0
	res := <-channels[len(channels)-1]
	log.Printf("conveyor len:%d\tresult:%d\tduration: %v\n", len(channels), res, time.Since(start))
}
