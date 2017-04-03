package example

import "log"

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	// time.Sleep(time.Microsecond * time.Duration(100))
	return n + factorial(n-1)
}

// SerialFactorial computes factorials for a number of numbers
func SerialFactorial(n int, printResult bool) {
	for i := 0; i < n; i++ {
		result := factorial(i % 8)
		if printResult {
			log.Println(result)
		}
	}
}
