package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func captureError(err error) {
	if err != nil {
		if err == io.EOF {
			fmt.Println("File end")
		}
		panic(err)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: gorp <filename> <word>")
		os.Exit(0)
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	captureError(err)
	defer file.Close()

	word := os.Args[2]
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), word) {
			fmt.Println(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
