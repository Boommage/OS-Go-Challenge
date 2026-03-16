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

		//read every line in the file and print to screen - not this currently has an error
		//file sizes too big throw an error for Scanner - look up why later
		//scanner := bufio.NewScanner(file)
		//for scanner.Scan() {
		//	fmt.Println(scanner.Text())
		//}
		//err = scanner.Err()
		//errc(err)

		//buf := make([]byte, 256)

		//for {
		//	n, err := file.Read(buf)
		//	errc(err)
		//	if n == 0 {
		//		break
		//	}
		//	fmt.Println(string(buf[:n]))
		//}

		buf := make([]byte, 4096)
		for n, err := file.Read(buf); err == nil; n, err = file.Read(buf) {
			os.Stdout.Write(buf[:n]) 
		}
	}
}
