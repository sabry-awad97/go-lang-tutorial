package main

import (
	"tutorial/demo/pkg/display"
	"tutorial/demo/pkg/msg"
)

func main() {
	msg.Hi()
	display.Display("hello from display")
	msg.Exciting("calling a package function")
}
