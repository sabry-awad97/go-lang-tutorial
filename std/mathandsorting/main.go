package main

import (
	"math"
	"math/rand"
	"time"
)

func mathFunc() {
	val1 := 279.00
	val2 := 48.95
	Printfln("Abs: %v", math.Abs(val1))
	Printfln("Ceil: %v", math.Ceil(val2))
	Printfln("Copysign: %v", math.Copysign(val1, -5))
	Printfln("Floor: %v", math.Floor(val2))
	Printfln("Max: %v", math.Max(val1, val2))
	Printfln("Min: %v", math.Min(val1, val2))
	Printfln("Mod: %v", math.Mod(val1, val2))
	Printfln("Pow: %v", math.Pow(val1, 2))
	Printfln("Round: %v", math.Round(val2))
	Printfln("RoundToEven: %v", math.RoundToEven(val2))
}

func generateRondomNumbers() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		Printfln("Value %v : %v", i, rand.Int())
	}

	for i := 0; i < 5; i++ {
		Printfln("Value %v : %v", i, rand.Intn(10))
	}
}

func generateIntRange() {
	IntRange := func(min, max int) int {
		return rand.Intn(max-min) + min
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		Printfln("Value %v : %v", i, IntRange(10, 20))
	}
}

func shuffle() {
	var names = []string{"Alice", "Bob", "Charlie", "Dora", "Edith"}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(names), func(first, second int) {
		Printfln("%v, %v", first, second)
		names[first], names[second] = names[second], names[first]
	})
	for i, name := range names {
		Printfln("Index %v: Name: %v", i, name)
	}
}

func main() {
	// mathFunc()
	// generateRondomNumbers()
	// generateIntRange()
	shuffle()
}
