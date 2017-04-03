package main

import "fmt"

// Lets look at structs:

// Human is a struct
// Uppercase properties are public, lowercase are private
type Human struct {
	Name  string
	Age   uint8
	House string
	mood  string
}

// Greet is a function on struct Human
func (h *Human) Greet() {
	fmt.Println("Greetings from house", h.House, "my name is", h.Name)
}

// go has type embedding so we can do things like this:

// RoyalHuman is a Human with a throne. It has all the properties and functions
// of Human, but also get the propery Title, and the function Proclaim.
type RoyalHuman struct {
	Human
	Title string
}

// Proclaim your royal status!
func (rh *RoyalHuman) Proclaim() {
	fmt.Printf("My name is %s, %s of %s, age %d, and I am %s\n",
		rh.Name, rh.Title, rh.House, rh.Age, rh.mood)
}

// structs can also have interfaces embedded in them

// WarriorInterface can Fight, then will inevitably die
type WarriorInterface interface {
	Fight()
	Die()
}

// WarriorRoyal is a royal at war time
type WarriorRoyal struct {
	RoyalHuman
	WarriorInterface
}

// Warrior is a bit of a stretch but I'm using it to fulfil WarriorInterface
type Warrior struct{}

// Die ...
func (w Warrior) Die() {
	fmt.Println("NOOOOOOOO")
}

// Fight ...
func (w Warrior) Fight() {
	fmt.Println("arrggggg")
}

func main() {
	// We can create a struct or create a struct and get a reference to it
	jon := Human{
		Name:  "Jon Snow",
		Age:   14,
		House: "Stark",
		mood:  "angsty",
	}

	samwell := &Human{
		Name:  "Samwell Tarley",
		Age:   16,
		House: "Tarley",
		// if a property is not set for a struct,
		// the default value for that properties type is used
	}

	jon.Greet()
	samwell.Greet()

	// structs with embedded fields still need the embedded type explicitly stated
	// Can't do:
	// danny := RoyalHuman{
	// 	Name:  "Daenerys Targaryen",
	// 	Age:   15,
	// 	House: "Targaryen",
	// 	mood:  "vengful",
	// 	Title: "Daenerys Stormborn of the House Targaryen, First of Her Name, the Unburnt, Queen of the Andals and the First Men, Khaleesi of the Great Grass Sea, Breaker of Chains, and Mother of Dragons",
	// }
	// Need to do:
	danny := RoyalHuman{
		Human: Human{
			Name:  "Daenerys Targaryen",
			Age:   15,
			House: "Targaryen",
			mood:  "vengful",
		},
		Title: "Daenerys Stormborn of the House Targaryen, First of Her Name, the Unburnt, Queen of the Andals and the First Men, Khaleesi of the Great Grass Sea, Breaker of Chains, and Mother of Dragons",
	}

	// You can also do this:
	joffery := Human{
		Name:  "Joffrey Baratheon",
		Age:   7,
		House: "Baratheon/Lanister, depends who you ask",
		mood:  "\"Lets kill stuff\"",
	}

	kingJoffery := RoyalHuman{
		Human: joffery,
		Title: "King",
	}

	danny.Greet()
	danny.Proclaim()

	kingJoffery.Proclaim()

	// Like with anything, structs can be abused and made complicated
	rhaegar := WarriorRoyal{
		RoyalHuman: RoyalHuman{
			Human: Human{
				Name:  "Rhaegar Targaryen",
				Age:   21,
				House: "Targaryen",
				mood:  "who knows",
			},
			Title: "Price",
		},
		WarriorInterface: Warrior{},
	}

	rhaegar.Fight()
	rhaegar.Proclaim()
	rhaegar.Greet()
	rhaegar.Die()

}
