package main

import "fmt"

func chanPasser(ch chan<- chan int) {
	var channels [5]chan int
	fmt.Println("Starting channel passer")

	total := 0

	for i := 0; i < 5; i++ {
		tmpCh := make(chan int)
		channels[i] = tmpCh
	}

	for index, chanPtr := range channels {
		ch <- chanPtr
		for j := total; j < total+5; j++ {
			chanPtr <- j
			fmt.Printf("Sent Value %d to channel %d\n", j, index)
		}
		total += 5
		close(chanPtr)
	}

	close(ch)

}

func chanReceiver(ch <-chan chan int, done chan bool) {
	fmt.Println("Starting channel receiver")
	for newChan := range ch {
		for val := range newChan {
			fmt.Println("Recevied value:", val)
		}
	}
	done <- true
}

func main() {
	chch := make(chan chan int)
	done := make(chan bool)
	go chanReceiver(chch, done)
	go chanPasser(chch)
	<-done
}
