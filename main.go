package main

import (
	"art/processing"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// --- CLI флаги ---
	encodePtr := flag.Bool("encode", false, "Switch to encoding mode.")
	showoutPtr := flag.Bool("showout", false, "Show output even if -o is defined.")
	multiPtr := flag.Bool("multi", false, "Enable multi-line input in the command line.")
	infilePtr := flag.String("i", "", "Specify input file to read as input.")
	outfilePtr := flag.String("o", "", "Specify output file to write output to.")
	flag.Parse()

	// --- Сбор входа ---
	var inputLines []string
	if *infilePtr != "" {
		lines, err := ReadLines(*infilePtr)
		if err != nil {
			log.Fatalf("Error reading input file %q: %v", *infilePtr, err)
		}
		inputLines = lines
	} else {
		args := flag.Args()
		if len(args) == 0 {
			log.Fatal("No input provided.")
		}
		if *multiPtr {
			inputLines = args
		} else {
			inputLines = args[:1]
		}
	}

	// --- Обработка (decode/encode) с поддержкой concurrency ---
	resultLines := processing.ProcessData(inputLines, *encodePtr)

	// --- Вывод ---
	if *showoutPtr || *outfilePtr == "" {
		for _, line := range resultLines {
			fmt.Println(line)
		}
	}
	if *outfilePtr != "" {
		if err := WriteLines(*outfilePtr, resultLines); err != nil {
			log.Fatalf("Error writing output file %q: %v", *outfilePtr, err)
		}
	}
}

// ReadLines читает все строки из файла в срез строк.
func ReadLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// WriteLines записывает все строки в файл (по одной на строку).
func WriteLines(path string, lines []string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, line := range lines {
		if _, err := w.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return w.Flush()
}
