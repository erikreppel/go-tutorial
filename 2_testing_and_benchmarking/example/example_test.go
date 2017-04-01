package example_test

import (
	"testing"

	"time"

	"github.com/erikreppel/go-tutorial/2_testing_and_benchmarking/example"
)

func TestSquareProducesCorrectResult(t *testing.T) {
	result := example.Square(2)
	if result != 4 {
		t.Fail()
	}
}

func TestConcatStringsProducesCorrectResult(t *testing.T) {
	result := example.ConcatStrings("SD", "AML")
	if result != "SDAML" {
		t.Fail()
	}
}

func TestUnderAgeUserCannotDrink(t *testing.T) {
	user := example.User{
		Name:     "Jim",
		ID:       "safigjalvn",
		Birthday: time.Date(1998, time.January, 28, 0, 0, 0, 0, time.Local),
	}

	if example.CanDrink(user, 19) == false {
		t.Fail()
	}

	if example.CanDrink(user, 21) {
		t.Fail()
	}

}

// Benchmarks =================================================================

func BenchmarkConcatStrings(b *testing.B) {
	for n := 0; n < b.N; n++ {
		example.ConcatStrings("abc", "def")
	}
}

func BenchmarkSquare(b *testing.B) {
	for n := 0; n < b.N; n++ {
		example.Square(n)
	}
}

func BenchmarkCanDrink(b *testing.B) {
	for n := 0; n < b.N; n++ {
		user := example.User{
			Name:     "Jim",
			ID:       "safigjalvn",
			Birthday: time.Date(1998, time.January, 28, 0, 0, 0, 0, time.Local),
		}
		example.CanDrink(user, n)
	}
}

func BenchmarkSimulateTimeout(b *testing.B) {
	for n := 0; n < b.N; n++ {
		example.SimulateTimeout()
	}
}
