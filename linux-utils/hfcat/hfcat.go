package main

import (
	"os"
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

		buf := make([]byte, 4096)
		for n, err := file.Read(buf); err == nil; n, err = file.Read(buf) {
			os.Stdout.Write(buf[:n]) 
		}
	}
}
