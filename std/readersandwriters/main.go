package main

import (
	"io"
	"strings"
)

func understandingReaders() {
	processData := func(reader io.Reader) {
		var buf = make([]byte, 2)

		for {
			count, err := reader.Read(buf)
			if count > 0 {
				Printfln("Read %v bytes: %v", count, string(buf[0:count]))
			}

			if err == io.EOF {
				break
			}
		}
	}

	r := strings.NewReader("Kayak")
	processData(r)
}

func main() {

}
