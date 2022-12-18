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

func scan() {
	var name string
	var category string
	var price float64
	fmt.Print("Enter text to scan: ")
	n, err := fmt.Scan(&name, &category, &price)
	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func scanSlice() {
	vals := make([]string, 3)
	ivals := make([]interface{}, 3)
	for i := 0; i < len(vals); i++ {
		ivals[i] = &vals[i]
	}
	fmt.Print("Enter text to scan: ")
	fmt.Scan(ivals...)
	Printfln("Name: %v", vals)
}

func scanNewLine() {
	var name string
	var category string
	var price float64
	fmt.Print("Enter text to scan: ")
	n, err := fmt.Scanln(&name, &category, &price)
	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func sourceScan() {
	var name string
	var category string
	var price float64
	source := "Lifejacket Watersports 48.95"
	n, err := fmt.Sscan(source, &name, &category, &price)
	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func templateScan() {
	var name string
	var category string
	var price float64
	source := "Product Lifejacket Watersports 48.95"
	template := "Product %s %s %f"
	n, err := fmt.Sscanf(source, template, &name, &category, &price)
	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func main() {
	// formatting()
	// scan()
	// scanSlice()
	// scanNewLine()
	// sourceScan()
	templateScan()
}
