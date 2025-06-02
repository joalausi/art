package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"./decoder"
	"./encoder"
	"./utils"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Error")
		return
	}

	flag.Parse()

	encodeMode := false
	multiline := false
	input := ""

	for _, arg := range args {
		if arg == "-e" {
			encodeMode = true
		} else if arg == "-m" {
			multiline = true
		} else {
			// Check if the argument is a .txt file from test_data
			if strings.HasSuffix(arg, ".txt") || strings.HasSuffix(arg, ".art") {
				content, err := utils.ReadFile("test-data/" + arg)
				if err != nil {
					fmt.Println("Error reading file:", err)
					return
				}
				input = content
			} else {
				input = arg
			}
		}
	}

	if input == "" {
		fmt.Println("Error")
		return
	}
	if encodeMode {
		if multiline {
			input = replaceEscapedNewlines(input)
			encoded := encoder.EncodeMultiLine(input)
			fmt.Print(encoded)
		} else {
			encoded := encoder.EncodeSingleLine(input)
			fmt.Print(encoded)
		}
	} else {
		if multiline {
			input = replaceEscapedNewlines(input)
			decoded, err := decoder.DecodeMultiLine(input)
			if err != nil {
				fmt.Println("Error")
			} else {
				fmt.Print(decoded)
			}

		} else {
			decoded, err := decoder.DecodeSingleLine(input)
			if err != nil {
				fmt.Println("Error")
			} else {
				fmt.Print(decoded)
			}
		}
	}

}

func replaceEscapedNewlines(input string) string {
	return strings.ReplaceAll(input, `\n`, "\n")
}

// package main

// import (
// 	"flag"
// 	"fmt"
// 	"os"
// 	// "art/decoder"
// 	// "art/encoder"
// )

// func main() {
// 	encodeMode := flag.Bool("e", false, "encode mode")
// 	multiline := flag.Bool("m", false, "multiline mode")
// 	flag.Parse()

// 	args := flag.Args()
// 	if len(args) == 0 {
// 		fmt.Println("Error")
// 		os.Exit(1)
// 	}

// 	var input string
// 	if *multiline {
// 		data, err := os.ReadFile(args[0])
// 		if err != nil {
// 			fmt.Println("Error")
// 			os.Exit(1)
// 		}
// 		input = string(data)
// 	} else {
// 		input = args[0]
// 	}

// 	var out string
// 	var err error
// 	if *encodeMode {
// 		out = encoder.Encode(input)
// 	} else {
// 		out, err = decoder.Decode(input)
// 	}
// 	if err != nil {
// 		fmt.Println("Error")
// 		os.Exit(1)
// 	}
// 	fmt.Print(out)
// }
