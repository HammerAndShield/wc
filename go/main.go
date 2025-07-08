package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	byteCountFlag := flag.Bool("c", false, "Count the number of bytes in the file.")
	flag.Parse()

	if len(flag.Args()) == 0 {
		println("No input file provided. Please use the -h for the help menu.")
		return
	}
	inputFile := flag.Args()[0]

	if *byteCountFlag {
		file, err := os.Open(inputFile)
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		reader := bufio.NewReader(file)
		buffer := make([]byte, 4096)
		byteCount := 0
		for {
			n, err := reader.Read(buffer)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				log.Fatalf("Error reading file: %v", err)
			}
			byteCount += n
		}

		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		fmt.Printf("%d %s\n", byteCount, inputFile)
	} else {
		println("No flags were provided. Please use the -h for the help menu.")
	}

}
