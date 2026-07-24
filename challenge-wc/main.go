package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

type Count struct {
	Lines int
	Words int
	Chars int
	Bytes int
}

func main() {
	var file *os.File
	var err error
	lFlag := flag.Bool("l", false, "Count Lines in a File")
	mFlag := flag.Bool("m", false, "Count Number of Characters in a File")
	cFlag := flag.Bool("c", false, "Count Number of Bytes in a File")
	wFlag := flag.Bool("w", false, "Count Number of Words in a File")

	flag.Parse()

	fileName := flag.Args()
	if(len(fileName) > 0){
		file, err = os.Open(fileName[0])
	} else {
		file = os.Stdin
	}
	
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	result := count(file)
	var output []string
	if *lFlag {
		output = append(output, fmt.Sprintf("%8d", result.Lines))
	}
	if *mFlag {
		output = append(output, fmt.Sprintf("%8d", result.Chars))
	}
	if *cFlag {
		output = append(output, fmt.Sprintf("%8d", result.Bytes))
	}
	if *wFlag {
		output = append(output, fmt.Sprintf("%8d", result.Words))
	}

	if len(output) == 0 {
		output = append(output, 
					fmt.Sprintf("%8d", result.Lines),
					fmt.Sprintf("%8d", result.Words),
					fmt.Sprintf("%8d", result.Bytes))
	}

	if len(fileName) > 0 {
		output = append(output, fileName[0])
	} 

	fmt.Println(strings.Join(output, " "))
}

func count(r io.Reader) Count {
	var lines, words, chars = 0, 0, 0
	data, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	str := string(data)
	chars = utf8.RuneCountInString(str)
	words = len(bytes.Fields(data))
	lbytes := len(data)

	for i := 0; i < lbytes; i++ {
		if(data[i] == '\n'){
			lines += 1
		}
	}

	return Count{
		Lines: lines,
		Words: words,
		Chars: chars,
		Bytes: lbytes,
	}
}