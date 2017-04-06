package main

import (
	"fmt"
	"sync"
	"time"
)

func waveGenerator(id int, chch chan chan int, wg *sync.WaitGroup) {
	values := make(chan int)
	chch <- values
	for i := id; i < 5; i++ {
		values <- i
		fmt.Printf("Sent %d from Wave Generator %d\n", i, id)
		time.Sleep(time.Second)
	}
	close(values)
	wg.Done()
}

func waveProcessor(id string, chch chan chan int, wg *sync.WaitGroup) {
	for ch := range chch {
		for val := range ch {
			fmt.Printf("%s: %d\n", id, val*2)
		}
	}
	wg.Done()
}

func main() {
	chch := make(chan chan int)

	generatorWG := &sync.WaitGroup{}
	processorWG := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		id := fmt.Sprintf("processor%d", i)
		go waveProcessor(id, chch, processorWG)
		processorWG.Add(1)
	}

	for i := 0; i < 5; i++ {
		go waveGenerator(i, chch, generatorWG)
		generatorWG.Add(1)
	}

	generatorWG.Wait()
	close(chch)
	processorWG.Wait()

}
