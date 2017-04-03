package example

import (
	"log"
	"sync"
)

// just any function that takes time to do something
func double(x int) int {
	// time.Sleep(time.Millisecond * time.Duration(100))
	return x * 2
}

// SerialDouble all the numbers up to n
func SerialDouble(n int, printResult bool) {
	for i := 0; i < n; i++ {
		result := double(i)
		if printResult {
			log.Println(result)
		}
	}
}

// ConcurrentDouble all the numbers up to n
func ConcurrentDouble(n int, nWorkers int, printResult bool) {

	numberChannel := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(nWorkers)
	for i := 0; i < nWorkers; i++ {

		go func(threadID int, numberChannel chan int) {

			for x := range numberChannel {
				if printResult {
					log.Printf("Goroutine #%d: %d\n", threadID, double(x))
				}
			}
			if printResult {
				log.Printf("Goroutine #%d finished", threadID)
			}
			wg.Done()

		}(i, numberChannel)
	}

	for i := 0; i < n; i++ {
		numberChannel <- i
	}

	close(numberChannel)
	wg.Wait()
}

func main() {

}
