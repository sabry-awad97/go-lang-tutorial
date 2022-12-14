package main

import (
	"math"
	"math/rand"
	"sort"
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

func sorting() {
	ints := []int{9, 4, 2, -1, 10}
	Printfln("Ints: %v", ints)
	sort.Ints(ints)
	Printfln("Ints Sorted: %v", ints)

	floats := []float64{279, 48.95, 19.50}
	Printfln("Floats: %v", floats)
	sort.Float64s(floats)
	Printfln("Floats Sorted: %v", floats)

	strings := []string{"Kayak", "Lifejacket", "Stadium"}
	Printfln("Strings: %v", strings)
	if !sort.StringsAreSorted(strings) {
		sort.Strings(strings)
		Printfln("Strings Sorted: %v", strings)
	} else {
		Printfln("Strings Already Sorted: %v", strings)
	}
}

func sortedCopy() {
	ints := []int{9, 4, 2, -1, 10}
	sortedInts := make([]int, len(ints))
	copy(sortedInts, ints)
	sort.Ints(sortedInts)
	Printfln("Ints: %v", ints)
	Printfln("Ints Sorted: %v", sortedInts)
}

func searchingSortedData() {
	ints := []int{9, 4, 2, -1, 10}
	sortedInts := make([]int, len(ints))
	copy(sortedInts, ints)
	sort.Ints(sortedInts)
	Printfln("Ints: %v", ints)
	Printfln("Ints Sorted: %v", sortedInts)

	indexOf4 := sort.SearchInts(sortedInts, 4)
	indexOf3 := sort.SearchInts(sortedInts, 3)
	Printfln("Index of 4: %v", indexOf4)
	Printfln("Index of 3: %v", indexOf3)

	Printfln("Index of 4: %v (present: %v)", indexOf4, sortedInts[indexOf4] == 4)
	Printfln("Index of 3: %v (present: %v)", indexOf3, sortedInts[indexOf3] == 3)
}

func sortingCustomDataTypes() {
	products := []Product{
		{"Kayak", 279},
		{"Lifejacket", 49.95},
		{"Soccer Ball", 19.50},
	}

	ProductSlices(products)

	for _, p := range products {
		Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
	}
}

func sortWithField() {
	products := []Product{
		{"Kayak", 279},
		{"Lifejacket", 49.95},
		{"Soccer Ball", 19.50},
	}

	ProductSlicesByName(products)

	for _, p := range products {
		Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
	}
}

func ComparisonFunction() {
	products := []Product{
		{"Kayak", 279},
		{"Lifejacket", 49.95},
		{"Soccer Ball", 19.50},
	}

	SortWith(products, func(p1, p2 Product) bool {
		return p1.Price < p2.Price
	})

	for _, p := range products {
		Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
	}
}

func main() {
	// mathFunc()
	// generateRondomNumbers()
	// generateIntRange()
	// shuffle()
	// sorting()
	// sortedCopy()
	// searchingSortedData()
	// sortingCustomDataTypes()
	// sortWithField()
	ComparisonFunction()
}
