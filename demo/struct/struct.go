package main

import "fmt"

type Passenger struct {
	Name string
	TicketNumber int
	Boarded bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	sabry := Passenger{"Sabry", 1, false}
	fmt.Println(sabry)

	var (
		ahmed = Passenger{Name: "Ahmed", TicketNumber: 2}
		mostafa = Passenger{Name: "Mostafa", TicketNumber: 3}
	)

	fmt.Println(ahmed, mostafa)

	var mohammed Passenger

	mohammed.Name = "Mohammed"
	mohammed.TicketNumber = 4

	fmt.Println(mohammed)

	sabry.Boarded = true

	bus := Bus{sabry}
	
	fmt.Println(bus)

	fmt.Println(bus.FrontSeat.Name, "is in the front seat")
}