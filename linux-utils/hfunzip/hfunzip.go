package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func main() {
	argc := len(os.Args)
	//check if no args are passed
	if argc < 2 {
		os.Stderr.Write([]byte("hfunzip: file1 [file2 ...]\n"))
		os.Exit(1)
	}

	//create a buffer for count and char
	num := make([]byte,4) //first 4 bytes
	char := make([]byte, 1) //remaining 1 byte

	//unzip all files in order
	for _, arg := range os.Args[1:] {
		file, err := os.Open(arg)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		//reading each num/char combo in file
		for {
			_, err := file.Read(num)
			if err != nil {
				break
			}
			_, err = file.Read(char)
			if err != nil {
				panic(err)
			}
			//convert binary count to int
			count := binary.LittleEndian.Uint32(num)
			var i uint32
			//print char based on count
			for i = 0; i < count; i++ {
				fmt.Printf("%c",char[0])
			}
		}
	}
}
