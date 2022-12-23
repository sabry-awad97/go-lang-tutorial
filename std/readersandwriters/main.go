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

func copyData() {
	processData := func(reader io.Reader, writer io.Writer) {
		count, err := io.Copy(writer, reader)
		if err == nil {
			Printfln("Read %v bytes", count)
		} else {
			Printfln("Error: %v", err.Error())
		}
	}

	r := strings.NewReader("Kayak")
	var builder strings.Builder
	processData(r, &builder)
	Printfln("String builder contents: %s", builder.String())
}

func pipe() {
	pipeReader, pipeWriter := io.Pipe()
	go GenerateData(pipeWriter)
	ConsumeData(pipeReader)
}

func multipleReader() {
	r1 := strings.NewReader("Kayak")
	r2 := strings.NewReader("Lifejacket")
	r3 := strings.NewReader("Canoe")
	concatReader := io.MultiReader(r1, r2, r3)
	ConsumeData(concatReader)
}

func main() {
	// understandingReaders()
	// understandingWriters()
	// copyData()
	// pipe()
	multipleReader()
}
