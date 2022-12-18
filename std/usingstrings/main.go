package main

import "fmt"

func getProductName(index int) (name string, err error) {
	if len(Products) > index {
		name = fmt.Sprintf("Name of product: %v", Products[index].Name)
	} else {
		err = fmt.Errorf("error for index %v", index)
	}
	return
}

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func formatting() {
	fmt.Println("Product:", Kayak.Name, "Price:", Kayak.Price)
	fmt.Print("Product:", Kayak.Name, "Price:", Kayak.Price, "\n")
	fmt.Printf("Product: %v, Price: $%4.2f\n", Kayak.Name, Kayak.Price)

	name, _ := getProductName(1)
	fmt.Println(name)

	_, err := getProductName(10)
	fmt.Println(err.Error())

	Printfln("Value: %v", Kayak)
	Printfln("Value with fields: %+v", Kayak)
	Printfln("Go syntax: %#v", Kayak)
	Printfln("Type: %T", Kayak)

	intNumber := 250
	Printfln("Binary: %b", intNumber)
	Printfln("Decimal: %d", intNumber)
	Printfln("Octal: %o, %O", intNumber, intNumber)
	Printfln("Hexadecimal: %x, %X", intNumber, intNumber)

	floatNumber := 279.00
	Printfln("Decimalless with exponent: %b", floatNumber)
	Printfln("Decimal with exponent: %e", floatNumber)
	Printfln("Decimal without exponent: %f", floatNumber)
	Printfln("Hexadecimal: %x, %X", floatNumber, floatNumber)
	Printfln("Decimal without exponent: >>%8.2f<<", floatNumber)
	Printfln("Decimal without exponent: >>%.2f<<", floatNumber)
	Printfln("Sign: >>%+.2f<<", floatNumber)
	Printfln("Zeros for Padding: >>%010.2f<<", floatNumber)
	Printfln("Right Padding: >>%-8.2f<<", floatNumber)

	Printfln("String: %s", name)
	Printfln("Character: %c", []rune(name)[0])
	Printfln("Unicode: %U", []rune(name)[0])

	Printfln("Bool: %t", len(name) > 1)
	Printfln("Bool: %t", len(name) > 100)

	Printfln("Pointer: %p", &name)
}

func main() {
	// formatting()
}