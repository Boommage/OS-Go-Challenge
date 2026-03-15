package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
)

//error handler for redundancy
func errc(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	//read every passed argument and cat it... - remember for in range is faster than a for loop
	for _, arg := range os.Args[1:] {
		//open the file
		file, err := os.Open(arg)
		errc(err)
		defer file.Close()

		//read every line in the file and print to screen - not this currently has an error
		//file sizes too big throw an error for Scanner - look up why later
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		err = scanner.Err()
		errc(err)
	} 
}
