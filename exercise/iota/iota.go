//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

type Operation int

//* Mathematical operations must be defined as constants using iota
// iota is used to automatically assign values
const (
	Add Operation = iota
	Subtract
	Multiply
	Divide
)

//* Write a receiver function that performs the mathematical operation
//  on two operands
func (op Operation) calculate(lhs, rhs int) int {
	switch op {
	case Add:
		return lhs + rhs
	case Subtract:
		return lhs - rhs
	case Multiply:
		return lhs * rhs
	case Divide:
		return lhs / rhs
	}
	panic("unhandled operation")
}

func (op Operation) String() string {
	switch op {
	case Add:
		return "Add"
	case Subtract:
		return "Subtract"
	case Multiply:
		return "Multiply"
	case Divide:
		return "Divide"
	default:
		return "Other Operation"
	}
}

func main() {
	//* The existing function calls in main() represent the API and cannot be changed
	fmt.Println(Add, Add.calculate(2, 2)) // = 4

	fmt.Println(Subtract, Subtract.calculate(10, 3)) // = 7

	fmt.Println(Multiply, Multiply.calculate(3, 3)) // = 9

	fmt.Println(Divide, Divide.calculate(100, 2)) // = 50
}
