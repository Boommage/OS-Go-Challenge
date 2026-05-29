package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	argc := len(os.Args)

	//check if no args are passed
	if argc < 2 {
		os.Stderr.Write([]byte("hfzip: file1 [file2 ...]\n"))
		os.Exit(1)
	}

	//buffer to hold final zipped result
	buf := new(bytes.Buffer)

	//zip through each file into one
	for _, arg := range os.Args[1:] {
		file, err := os.Open(arg)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		//use a file reader to capture first byte
		read := bufio.NewReader(file)
		first, err := read.ReadByte()
		if err == io.EOF {
			return
		}
		var current byte = first
		var count uint32 = 1

		//read through the rest of the bytes
		for {
			b, err := read.ReadByte()
			if err == io.EOF {
				break
			}
			if b == current {
				count++
			} else {
				//write out count and char in binary format
				binary.Write(buf, binary.LittleEndian, count)
				buf.WriteByte(current)

				//reset count and char
				current = b
				count = 1
			}
		}
		//write remaining count and char
		binary.Write(buf, binary.LittleEndian, count)
		buf.WriteByte(current)
	}
	//write bytes to std out
	fmt.Print(buf)
}