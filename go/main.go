package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {

	byteCountFlag := flag.Bool("c", false, "Count the number of bytes in the file.")
	lineCountFlag := flag.Bool("l", false, "Count the number of lines in the file.")
	wordCountFlag := flag.Bool("w", false, "Count the numbers of words in the file")
	charCountFlag := flag.Bool("m", false, "The number of characters in the file")
	processingTimeFlag := flag.Bool("p", false, "The processing time will be displayed")
	flag.Parse()

	if len(flag.Args()) == 0 {
		println("No input file provided. Please use -h for the help menu.")
		return
	}
	inputFile := flag.Args()[0]

	now := time.Now()

	if *charCountFlag {
		count, err := countRunes(inputFile)
		if err != nil {
			log.Fatalf("Error reading the char count: %v", err)
		}

		fmt.Printf("  %d %s", count, inputFile)
	} else {
		file, err := os.Open(inputFile)
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}
		defer file.Close()

		noFlags := !*byteCountFlag && !*lineCountFlag && !*wordCountFlag

		scanner := bufio.NewScanner(file)

		lineCount := 0
		wordCount := 0
		charCount := 0
		for scanner.Scan() {
			line := scanner.Text()
			lineCount++
			wordCount += len(strings.Fields(line))

			if *charCountFlag {
				charCount += utf8.RuneCountInString(line)
			}
		}

		res := "  "
		if *lineCountFlag || noFlags {
			res += fmt.Sprintf("%d  ", lineCount)
		}
		if *wordCountFlag || noFlags {
			res += fmt.Sprintf("%d  ", wordCount)
		}
		if *byteCountFlag || noFlags {
			info, err := file.Stat()
			if err != nil {
				log.Fatalf("Error getting file info: %v", err)
			}
			byteCount := info.Size()
			res += fmt.Sprintf("%d  ", byteCount)
		}

		res += filepath.Base(file.Name())

		println(res)
	}

	if *processingTimeFlag {
		p := time.Since(now)
		fmt.Printf("  Processing took %dms", p.Milliseconds())
	}
}

func countRunes(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	var n int
	for {
		_, _, err := rd.ReadRune()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return 0, err
		}
		n++
	}
	return n, nil
}
