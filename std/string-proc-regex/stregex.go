package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

func convertCase() {
	description := "A boat for sailing"
	fmt.Println("Original:", description)

	caser := cases.Title(language.English)
	fmt.Println("Title:", caser.String(description))

	specialChar := "\u01c9"
	fmt.Println("Original:", specialChar, []byte(specialChar))
	upperChar := strings.ToUpper(specialChar)
	fmt.Println("Upper:", upperChar, []byte(upperChar))
	titleChar := strings.ToTitle(specialChar)
	fmt.Println("Title:", titleChar, []byte(titleChar))

	product := "Kayak"
	for _, char := range product {
		fmt.Println(string(char), "Upper case:", unicode.IsUpper(char))
	}
}

func main() {
	// compareStrings()
	// compareBytes()
	convertCase()
}
