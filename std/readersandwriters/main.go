package main

import (
	"bufio"
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

func multipleWriters() {
	var w1 strings.Builder
	var w2 strings.Builder
	var w3 strings.Builder
	combinedWriter := io.MultiWriter(&w1, &w2, &w3)
	GenerateData(combinedWriter)

	// the same data is written to the three individual writers
	Printfln("Writer #1: %v", w1.String())
	Printfln("Writer #2: %v", w2.String())
	Printfln("Writer #3: %v", w3.String())
}

func EchoData() {
	r1 := strings.NewReader("Kayak")
	r2 := strings.NewReader("Lifejacket")
	r3 := strings.NewReader("Canoe")
	concatReader := io.MultiReader(r1, r2, r3)
	var writer strings.Builder
	teeReader := io.TeeReader(concatReader, &writer)
	ConsumeData(teeReader)
	Printfln("Echo data: %v", writer.String())
}

func limitData() {
	r1 := strings.NewReader("Kayak")
	r2 := strings.NewReader("Lifejacket")
	r3 := strings.NewReader("Canoe")
	concatReader := io.MultiReader(r1, r2, r3)
	limited := io.LimitReader(concatReader, 5)
	ConsumeData(limited)
}

func readerWrapper() {
	text := "It was a boat. A small boat."
	var reader io.Reader = NewCustomReader(strings.NewReader(text))
	var writer strings.Builder
	slice := make([]byte, 5)

	buffered := bufio.NewReader(reader)
	for {
		count, err := buffered.Read(slice)
		if count > 0 {
			Printfln("Buffer size: %v, buffered: %v",
				buffered.Size(), buffered.Buffered())
			writer.Write(slice[0:count])
		}
		if err != nil {
			break
		}
	}
	Printfln("Read data: %v", writer.String())
}

func main() {
	// understandingReaders()
	// understandingWriters()
	// copyData()
	// pipe()
	// multipleReader()
	// multipleWriters()
	// EchoData()
	// limitData()
	readerWrapper()
}
