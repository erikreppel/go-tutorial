package example

import (
	"math/rand"
	"time"
)

// Square multiplies a number by its self
func Square(x int) int {
	return x * x
}

// SimulateTimeout simulates an operation that takes some period of time by
// sleeping for a random value in (0, 500] ms.
func SimulateTimeout() {
	r := rand.Intn(500)
	time.Sleep(time.Millisecond * time.Duration(r))
}

// User is a struct describing a hypothetical user of a website
type User struct {
	Name     string
	ID       string
	Birthday time.Time
}

// CanDrink checks if a user can drink based on their birthday
func CanDrink(user User, drinkingAge int) bool {
	userAge := time.Since(user.Birthday)
	drinkingAgeHours := time.Hour * time.Duration(24*365*drinkingAge)
	if userAge > drinkingAgeHours {
		return true
	}
	return false
}

// ConcatStrings is an example of an implicit return
func ConcatStrings(string1, string2 string) (concated string) {
	concated = string1 + string2 // notice "=" not ":="
	return
}
