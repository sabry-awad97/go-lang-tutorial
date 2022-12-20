package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f := os.Stdin
	defer f.Close()

	scanner := bufio.NewReader(f)

	sum := 0

	for {
		input, inputErr := scanner.ReadString(' ')

		num, convErr := strconv.Atoi(strings.TrimSpace(input))
		if convErr != nil {
			fmt.Println(convErr)
		} else {
			sum += num
		}

		if inputErr == io.EOF {
			break
		}
		if inputErr != nil {
			fmt.Println("Error reading Stdin:", inputErr)
		}
	}
	fmt.Printf("Sum: %v\n", sum)
}
