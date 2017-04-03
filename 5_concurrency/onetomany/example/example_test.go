package example_test

import (
	"testing"

	"github.com/erikreppel/go-tutorial/5_concurrency/onetomany/example"
)

func BenchmarkSerialDouble(b *testing.B) {
	for n := 0; n < b.N; n++ {
		example.SerialDouble(n, false)
	}
}

func BenchmarkConcurrentDouble(b *testing.B) {
	for n := 0; n < b.N; n++ {
		example.ConcurrentDouble(n, 500, false)
	}
}
