package main

import (
	"bufio"
	"fmt"
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

func unbufferedWrites() {
	text := "It was a boat. A small boat."
	var builder strings.Builder
	var writer = NewCustomWriter(&builder)
	for i := 0; true; {
		end := i + 5
		if end >= len(text) {
			writer.Write([]byte(text[i:]))
			break
		}
		writer.Write([]byte(text[i:end]))
		i = end
	}
	Printfln("Written data: %v", builder.String())
}

func bufferedWrites() {
	text := "It was a boat. A small boat."
	var builder strings.Builder
	var writer = bufio.NewWriterSize(NewCustomWriter(&builder), 20)
	for i := 0; true; {
		end := i + 5
		if end >= len(text) {
			writer.Write([]byte(text[i:]))
			writer.Flush()
			break
		}
		writer.Write([]byte(text[i:end]))
		i = end
	}
	Printfln("Written data: %v", builder.String())
}

func scanFromReader(reader io.Reader, template string, vals ...interface{}) (int, error) {
	return fmt.Fscanf(reader, template, vals...)
}

func scan() {
	reader := strings.NewReader("Kayak Watersports $279.00")
	var name, category string
	var price float64
	scanTemplate := "%s %s $%f"
	_, err := scanFromReader(reader, scanTemplate, &name, &category, &price)
	if err != nil {
		Printfln("Error: %v", err.Error())
	} else {
		Printfln("Name: %v", name)
		Printfln("Category: %v", category)
		Printfln("Price: %.2f", price)
	}
}

func scanSingle(reader io.Reader, val interface{}) (int, error) {
	return fmt.Fscan(reader, val)
}

func scanGradual() {
	reader := strings.NewReader("Kayak Watersports $279.00")
	for {
		var str string
		_, err := scanSingle(reader, &str)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		Printfln("Value: %v", str)
	}
}

func writeFormatted() {
	writeFormatted := func(writer io.Writer, template string, vals ...interface{}) {
		fmt.Fprintf(writer, template, vals...)
	}

	var writer strings.Builder
	template := "Name: %s, Category: %s, Price: $%.2f"
	writeFormatted(&writer, template, "Kayak", "Watersports", float64(279))
	fmt.Println(writer.String())
}

func writeReplaced() {
	writeReplaced := func(writer io.Writer, str string, subs ...string) {
		replacer := strings.NewReplacer(subs...)
		replacer.WriteString(writer, str)
	}

	text := "It was a boat. A small boat."
	subs := []string{"boat", "kayak", "small", "huge"}

	var writer strings.Builder
	writeReplaced(&writer, text, subs...)
	fmt.Println(writer.String())
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
	// readerWrapper()
	// unbufferedWrites()
	// bufferedWrites()
	// scan()
	// scanGradual()
	// writeFormatted()
	writeReplaced()
}
