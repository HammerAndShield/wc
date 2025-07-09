package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	byteCountFlag := flag.Bool("c", false, "Count the number of bytes in the file.")
	lineCountFlag := flag.Bool("l", false, "Count the number of lines in the file.")
	wordCountFlag := flag.Bool("w", false, "Count the numbers of words in the file")
	flag.Parse()

	if len(flag.Args()) == 0 {
		println("No input file provided. Please use -h for the help menu.")
		return
	}
	inputFile := flag.Args()[0]

	if *byteCountFlag || *lineCountFlag || *wordCountFlag {
		file, err := os.Open(inputFile)
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		scanner := bufio.NewScanner(file)

		lineCount := 0
		wordCount := 0
		for scanner.Scan() {
			line := scanner.Text()

			lineCount++
			wordCount += len(strings.Fields(line))
		}

		res := ""
		if *byteCountFlag {
			info, err := file.Stat()
			if err != nil {
				log.Fatalf("Error getting file info: %w", err)
			}
			byteCount := info.Size()
			res += fmt.Sprintf("%d ", byteCount)
		}
		if *lineCountFlag {
			res += fmt.Sprintf("%d ", lineCount)
		}
		if *wordCountFlag {
			res += fmt.Sprintf("%d ", wordCount)
		}
		res += filepath.Base(file.Name())

		print(res)
	} else {
		println("No flags were provided. Please use the -h for the help menu.")
	}

}
