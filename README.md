# Art Encoder and Decoder

A simple CLI tool to encode and decode art and test files.

## Project Structure 

art/
├── main.go # CLI entry point
├── decoder/ # Decoder package
│ └── decoder.go
├── encoder/ # Encoder package
│ └── encoder.go
├── utils/ # Utility functions (e.g. file reading)
│ └── utils.go
└── test-data/ # Sample text files for testing

## Usage

I recommed building the program to a binary with `go build -o art`. From now on `./art` can be considered equal to `go run .`

- Decode the cats art: `./art -i examples/cats.encoded.txt`
- Encode the doomer art: `./art -encode -i examples/doomer.art.txt`
- Decode the plane art and write output to a file: `./art -i examples/plane.encoded.txt -i output.txt`
- Encode the globe art, write output to a file and also display the result in the command line: `./art -encode -showout -i examples/globe.art.txt -o output.txt`

## flags

-e encoder mode on
-m multiline mode on

## Example

go run . -m "lion.encoded.txt"
go run . -e -m "lion.art.txt"


## Extras
 It features a multi-line decoder.
 It can encode text-based art.
 It can encode multi-line text-based art.


Individual testing has completed all the test's in the project by me.