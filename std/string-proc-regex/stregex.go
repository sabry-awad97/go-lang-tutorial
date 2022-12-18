package main

import (
	"bytes"
	"fmt"
	"regexp"
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

func inspectStrings() {
	description := "A boat for one person"
	fmt.Println("Count: o", strings.Count(description, "o"))
	fmt.Println("Index: o", strings.Index(description, "o"))
	fmt.Println("LastIndex: o", strings.LastIndex(description, "o"))
	fmt.Println("IndexAny: abcd", strings.IndexAny(description, "abcd"))
	fmt.Println("LastIndex: o", strings.LastIndex(description, "o"))
	fmt.Println("LastIndexAny: abcd", strings.LastIndexAny(description, "abcd"))

	fmt.Println("IndexFunc:", strings.IndexFunc(description, func(r rune) bool {
		return r == 'B' || r == 'b'
	}))
}

func splitStrings() {
	description := "A  boat  for  one  person"

	for _, x := range strings.Fields(description) {
		fmt.Println("Fields >>" + x + "<<")
	}

	for _, x := range strings.FieldsFunc(description, func(r rune) bool {
		return r == ' '
	}) {
		fmt.Println("FieldsFunc >>" + x + "<<")
	}

	for _, x := range strings.Split(description, " ") {
		fmt.Println("Split >>" + x + "<<")
	}

	for _, x := range strings.SplitN(description, " ", 3) {
		fmt.Println("SplitN >>" + x + "<<")
	}

	for _, x := range strings.SplitAfter(description, " ") {
		fmt.Println("SplitAfterN >>" + x + "<<")
	}

	for _, x := range strings.SplitAfterN(description, " ", 3) {
		fmt.Println("SplitAfterN >>" + x + "<<")
	}
}

func trimStrings() {
	description := "A boat for one person"
	fmt.Println("Trimmed:", ">>"+strings.TrimSpace("  Alice  ")+"<<")
	fmt.Println("Trimmed:", ">>"+strings.Trim(description, "Asno ")+"<<")

	fmt.Println("Trimmed:", strings.TrimPrefix(description, "A boat "))
	fmt.Println("Not trimmed:", strings.TrimPrefix(description, "A hat "))

	fmt.Println("Trimmed:", strings.TrimFunc(description, func(r rune) bool {
		return r == 'A' || r == 'n'
	}))
}

func alterStrings() {
	text := "It was a boat. A small boat."
	fmt.Println("Replace:", strings.Replace(text, "boat", "canoe", 1))
	fmt.Println("Replace All:", strings.ReplaceAll(text, "boat", "truck"))

	fmt.Println("Mapped:", strings.Map(func(r rune) rune {
		if r == 'b' {
			return 'c'
		}
		return r
	}, text))

	fmt.Println("Replaced:", strings.NewReplacer("boat", "kayak", "small", "huge").Replace(text))

	fmt.Println("Joined:", strings.Join(strings.Fields(text), "--"))
}

func buildString() {
	text := "It was a boat. A small boat."
	var builder strings.Builder
	for _, field := range strings.Fields(text) {
		if field == "small" {
			builder.WriteString("very ")
		}
		builder.WriteString(field)
		builder.WriteRune(' ')
	}
	fmt.Println("String:", builder.String())
}

func matchString() {
	description := "A boat for one person"
	match, err := regexp.MatchString("[A-z]oat", description)
	if err == nil {
		fmt.Println("Match:", match)
	} else {
		fmt.Println("Error:", err)
	}

	pattern, compileErr := regexp.Compile("[A-z]oat")
	question := "Is that a goat?"
	preference := "I like oats"

	if compileErr == nil {
		fmt.Println("Description:", pattern.MatchString(description))
		fmt.Println("Question:", pattern.MatchString(question))
		fmt.Println("Preference:", pattern.MatchString(preference))
	} else {
		fmt.Println("Error:", compileErr)
	}

}

func findSubStrings() {
	getSubstring := func(s string, indices []int) string {
		return string(s[indices[0]:indices[1]])
	}

	pattern := regexp.MustCompile("K[a-z]{4}|[A-z]oat")
	description := "Kayak. A boat for one person."
	firstIndex := pattern.FindStringIndex(description)
	allIndices := pattern.FindAllStringIndex(description, -1)
	fmt.Println("First index", firstIndex[0], "-", firstIndex[1],
		"=", getSubstring(description, firstIndex))
	for i, idx := range allIndices {
		fmt.Println("Index", i, "=", idx[0], "-",
			idx[1], "=", getSubstring(description, idx))
	}

	fmt.Println("First match:", pattern.FindString(description))

	for i, m := range pattern.FindAllString(description, -1) {
		fmt.Println("Match", i, "=", m)
	}
}

func splitRegex() {
	pattern := regexp.MustCompile(" |boat|one")
	description := "Kayak. A boat for one person."
	split := pattern.Split(description, -1)
	for _, s := range split {
		if s != "" {
			fmt.Println("Substring:", s)
		}
	}
}

func main() {
	// compareStrings()
	// compareBytes()
	// convertCase()
	// inspectStrings()
	// splitStrings()
	// trimStrings()
	// alterStrings()
	// buildString()
	// matchString()
	// findSubStrings()
	splitRegex()
}
