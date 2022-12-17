package main

import (
	"bytes"
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

func compareBytes() {
	price := "€100"
	fmt.Println("Price in bytes: ", []byte(price))
	fmt.Println("Strings Prefix €: ", strings.HasPrefix(price, "€"))
	fmt.Println("Bytes Prefix:", bytes.HasPrefix([]byte(price),
		[]byte{226, 130}))
}

func main() {
	// compareStrings()
	compareBytes()
}
