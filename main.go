package main

import (
	"flag"
	"fmt"
	"os"
	// "art/decoder"
	// "art/encoder"
)

func main() {
	// флаги
	encodeMode := flag.Bool("e", false, "encode mode")
	multiline := flag.Bool("m", false, "multiline mode")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Error")
		os.Exit(1)
	}

	var input string
	if *multiline {
		data, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Println("Error")
			os.Exit(1)
		}
		input = string(data)
	} else {
		input = args[0]
	}

	var out string
	var err error
	if *encodeMode {
		out = encoder.Encode(input)
	} else {
		out, err = decoder.Decode(input)
	}
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}
	fmt.Print(out)
}
