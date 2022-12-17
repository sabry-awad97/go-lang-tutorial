package main

import (
	"fmt"
	"strings"
)

func compareStrings() {
	product := "Kayak"
	fmt.Println("Product:", product)

	fmt.Println("Contains: yak", strings.Contains(product, "yak"))
	fmt.Println("ContainsAny: abc", strings.ContainsAny(product, "abc"))
	fmt.Println("ContainsRune: K", strings.ContainsRune(product, 'K'))
	fmt.Println("EqualFold: KAYAK", strings.EqualFold(product, "KAYAK"))
	fmt.Println("HasPrefix: Ka", strings.HasPrefix(product, "Ka"))
	fmt.Println("HasSuffix: yak", strings.HasSuffix(product, "yak"))
}

func main() {
	compareStrings()
}
