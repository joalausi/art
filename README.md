# Art Encoder/Decoder

This is a command-line utility written in Go for encoding and decoding text files. The encoding format uses a simple scheme that repeats a line `STR` `N` times. In the `examples` directory you’ll find original and encoded ASCII-art files to try out the program.

## Installation

## Build

To build the executable, run:

```bash
go build -o art
```

You can also run the tool directly with:

```bash
go run .
```

## Usage

```
./art [options] [lines...]
```

- If you use the `-i` flag, input is read from a file.
- Without `-i`, any command-line arguments are treated as the input lines.
- By default, the tool runs in **decoding** mode.

### Main Flags

- `-encode` — switch to **encoding** mode.  
- `-multi`  — treat each argument as a separate line (multi-line mode).  
- `-i FILE` — path to the input file.  
- `-o FILE` — path to the output file.  
- `-showout` — print the result on screen even if you’ve used `-o`.  

## Examples

**Decoding a file:**

```bash
./art -i examples/cats.encoded.txt
```

**Encoding and saving the result:**

```bash
./art -encode -i examples/alpha.art.txt -o alpha.txt
```

**Multi-line input from arguments:**

```bash
./art -encode -multi "AAA" "BBB"
```

## Project Structure

```bash

art/
├── go.mod                  
├── main.go            — the CLI entry point.
├── examples           — sample files to demonstrate how it works.
└── processing/        — package with the encoding, decoding, and parsing logic.
    ├── parser.go
    ├── decoder.go
    ├── encoder.go
    └── processing.go
```

## Contact

For more information, feel free to reach out on Discord: **joalausi**
