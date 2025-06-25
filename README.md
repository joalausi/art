# Art Encoder/Decoder

This project contains both a command line utility and a small web server for
encoding and decoding text based art.  The encoding format uses a simple scheme
that repeats a line `STR` `N` times.  In the `examples` directory you’ll find
original and encoded ASCII-art files to try out the program.

## Installation

```bash
git clone https://gitea.kood.tech/artemchornonoh1/art.git
cd ~/art
```

## Build Terminal

To build the executable, run:

```bash
go build -o art
```

You can also run the tool directly with:

```bash
go run .
```

## Running the Web Interface

The repository also includes a small HTTP server that hosts a web page for
encoding and decoding art.

```bash
# build the server
go build -o art-server server_main.go
```
Start with: `./art-server`

```bash
# or run it directly
go run server_main.go
```

Navigate to http://localhost:22459 in your browser to use the interface. Use
the radio buttons to switch between decoding and encoding modes and submit your
text using the form.

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
