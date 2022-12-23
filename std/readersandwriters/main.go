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

func understandingWriters() {
	processData := func(reader io.Reader, writer io.Writer) {
		var buf = make([]byte, 2)

		for {
			count, err := reader.Read(buf)
			if count > 0 {
				writer.Write(buf[0:count])
				Printfln("Read %v bytes: %v", count, string(buf[0:count]))
			}

			if err == io.EOF {
				break
			}
		}
	}

	r := strings.NewReader("Kayak")
	var builder strings.Builder
	processData(r, &builder)
	Printfln("String builder contents: %s", builder.String())
}

func main() {
	// understandingReaders()
	understandingWriters()
}
