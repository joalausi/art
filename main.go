package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Check if the file exists
	filePath := "example.txt"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("File does not exist: %s", filePath)
	}

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	// Read the file content
	content := make([]byte, 1024)
	n, err := file.Read(content)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	fmt.Printf("File content: %s\n", string(content[:n]))
}
