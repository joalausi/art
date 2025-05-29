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

download and build the project 

cd art
go build -o art

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