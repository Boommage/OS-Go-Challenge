package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
)

func errc(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func grep(file *os.File, term string) {
	scanner := bufio.NewScanner(file)
	
	//read line-by-line
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, term) {
			fmt.Println(line)
		}
	}
	errc(scanner.Err())
}

func main() {
	argNum := len(os.Args)
	//if no args are passed
	if argNum < 2 {
			fmt.Println("hfgrep: searchterm [file ...]")
			os.Exit(1)
	}

	term := os.Args[1]

	//if only a search term is passed
	if argNum == 2 {
		grep(os.Stdin, term)
		return
	}

	//if files(s) and a search term is passed
	for _, arg := range os.Args[2:] {
		file, err := os.Open(arg)
		errc(err)
		defer file.Close()
		grep(file, term)
	}
}
